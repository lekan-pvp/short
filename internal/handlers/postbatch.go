package handlers

import (
	"context"
	"encoding/json"
	"github.com/lekan-pvp/short/internal/cookies"
	"github.com/lekan-pvp/short/internal/models"
	"net/http"
	"strings"
)

// PostBatch is a handler that accepts in the request body a set of URLs to shorten in the format:
//
// Endpoint: POST /api/shorten/batch
//
//  [
//    {
//        "correlation_id": "<string id>",
//        "original_url": "<URL for shorten>"
//    },
//    ...
//  ]
//
// As a response, the handler should return data in the format:
//
//  [
//    {
//        "correlation_id": "string id from request object",
//        "short_url": "<result short URL>"
//    },
//    ...
//  ]
//
// Possible response statuses:
// 201 Created it's OK
// 500 Internal Server Error
func PostBatch(repo Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, stop := context.WithCancel(r.Context())
		defer stop()

		var uuid string

		var in []models.BatchRequest

		cookie, err := r.Cookie("token")
		if err != nil || !cookies.CheckCookie(cookie) {
			cookie = cookies.New()
		}

		http.SetCookie(w, cookie)

		values := strings.Split(cookie.Value, ":")
		uuid = values[0]

		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res, err := repo.BatchShorten(ctx, uuid, in)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)

		if err := json.NewEncoder(w).Encode(&res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
