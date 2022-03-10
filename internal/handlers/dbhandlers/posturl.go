package dbhandlers

import (
	"context"
	"github.com/jackc/pgerrcode"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/cookies"
	"github.com/lekan-pvp/short/internal/dbrepo"
	"github.com/lekan-pvp/short/internal/makeshort"
	"github.com/lib/pq"
	"io"
	"log"
	"net/http"
	"strings"
)

// PostURL is a handler that makes a short url and save it in database.
//
// Endpoint POST /
//
// Content-Type: text/plain
//
// Request body example:
// http://yandex.ru
//
// Possible response statuses:
// 201 Status Created
// 401 Status Unauthorized
// 400 Status Bad Request
// 500 Status Internal Server Error
func PostURL(w http.ResponseWriter, r *http.Request) {
	var uuid string
	var cookie *http.Cookie
	var err error

	cookie, err = r.Cookie("token")
	if err != nil || !cookies.CheckCookie(cookie) {
		cookie = cookies.CreateCookie()
	}

	http.SetCookie(w, cookie)

	values := strings.Split(cookie.Value, ":")
	if len(values) != 2 {
		log.Println("Unauthorized")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	uuid = values[0]

	body, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url := string(body)
	short := makeshort.GenerateShortLink(url, uuid)

	ctx, stop := context.WithCancel(r.Context())
	defer stop()

	record := dbrepo.Storage{
		UUID:        uuid,
		ShortURL:    short,
		OriginalURL: url,
	}

	status := http.StatusCreated

	short, err = dbrepo.PostURL(ctx, record)

	log.Println(short)

	if err != nil {
		if err.(*pq.Error).Code == pgerrcode.UniqueViolation {
			status = http.StatusConflict
		} else {
			log.Println("error insert in DB:", err)
			http.Error(w, err.Error(), 500)
			return
		}
	}

	baseURL := config.GetBaseURL()

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	w.Write([]byte(baseURL + "/" + short))

}
