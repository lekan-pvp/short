package dbhandlers

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/short/internal/dbrepo"
	"log"
	"net/http"
)

// GetShort is a handler that receives original URL by short URL.
//
// Endpoint:
// GET /{short}
//
// where {short} is a short URL
//
// Content-Type: text/plain
//
// Possible response statuses:
// 307 Temporary Redirect - Success
// 400 Bad Request
// 410 Gone if record in database is deleted
// 404 Not Found if record not found in database
func GetShort(w http.ResponseWriter, r *http.Request) {
	ctx, stop := context.WithCancel(r.Context())
	defer stop()

	short := chi.URLParam(r, "short")
	if short == "" {
		log.Println("SHORT ERROR")
		http.Error(w, "url is empty", http.StatusNotFound)
		return
	}

	origin, err := dbrepo.GetOriginal(ctx, short)
	if err != nil {
		log.Printf("error GetOriginal %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if origin == nil {
		log.Println("Not found")
		http.NotFound(w, r)
		return
	}

	if !origin.IsDeleted() {
		log.Println("NOT DELETED")
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Location", origin.URL)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	if origin.IsDeleted() {
		log.Println("DELETED")
		w.WriteHeader(http.StatusGone)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		return
	}
}
