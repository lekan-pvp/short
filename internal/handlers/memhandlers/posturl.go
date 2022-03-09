package memhandlers

import (
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/cookies"
	"github.com/lekan-pvp/short/internal/makeshort"
	"github.com/lekan-pvp/short/internal/memrepo"
	"io"
	"log"
	"net/http"
	"strings"
)

// PostURL is a handler that makes a short url and save it in database.
//
// Endpoint / [post]
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

	record := memrepo.Storage{
		UUID:        uuid,
		ShortURL:    short,
		OriginalURL: url,
	}

	err = memrepo.PostURL(record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	baseURL := config.GetBaseURL()

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(baseURL + "/" + short))
}
