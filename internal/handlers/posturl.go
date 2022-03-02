package handlers

import (
	"context"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/cookies"
	"github.com/lekan-pvp/short/internal/makeshort"
	"github.com/lekan-pvp/short/internal/memrepo"
	"io"
	"log"
	"net/http"
	"strings"
)

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

	record := memrepo.Storage{
		UUID:        uuid,
		ShortURL:    short,
		OriginalURL: url,
	}

	err = memrepo.PostURL(ctx, record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	baseURL := config.GetBaseURI()

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(baseURL + "/" + short))
}
