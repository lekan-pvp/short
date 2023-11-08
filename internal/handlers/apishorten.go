package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/jackc/pgerrcode"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/cookies"
	"github.com/lekan-pvp/short/internal/makeshort"
	"github.com/lekan-pvp/short/internal/models"
	"github.com/lib/pq"
)

// APIShorten is a handler to make short URL and save them into database.
//
// Endpoint:
// /api/shorten [post]
//
// Content-Type: application/json
//
// Request body example:
//
//	{
//	  "url": "http://google.com"
//	}
//
// "url" is an original URL for making a short URL for one
//
// Possible response statuses:
//
// 201 Created Success status
// 500 Internal grpcserver error
// 401 Unauthorized
// 400 Bed Request
// 409 Status Conflict
func APIShorten(repo Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, stop := context.WithCancel(r.Context())
		defer stop()

		// Authorization is provided by the creation cookie
		cookie, err := r.Cookie("token")
		if err != nil || !cookies.CheckCookie(cookie) {
			cookie = cookies.New()
		}

		http.SetCookie(w, cookie)

		values := strings.Split(cookie.Value, ":")
		if len(values) != 2 {
			log.Println("Unauthorized")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		uuid := values[0]

		long := &models.URL{}

		if err := json.NewDecoder(r.Body).Decode(long); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Println(long)

		short := makeshort.GenerateShortLink(long.URL, uuid)

		record := models.Storage{
			UUID:          uuid,
			ShortURL:      short,
			OriginalURL:   long.URL,
			CorrelationID: "123",
			DeleteFlag:    false,
		}

		status := http.StatusCreated

		short, err = repo.PostURL(ctx, record)
		if err != nil {
			if err.(*pq.Error).Code == pgerrcode.UniqueViolation {
				status = http.StatusConflict
			} else {
				log.Println("error insert in DB:", err)
				http.Error(w, err.Error(), 500)
				return
			}
		}

		base := config.Cfg.BaseURL

		result := models.ResultResponse{
			Result: base + "/" + short,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)

		if err := json.NewEncoder(w).Encode(&result); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

	}
}
