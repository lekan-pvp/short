package dbhandlers

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/short/internal/dbrepo"
	"log"
	"net/http"
)

func GetShort(w http.ResponseWriter, r *http.Request) {
	ctx, stop := context.WithCancel(r.Context())
	defer stop()

	short := chi.URLParam(r, "short")
	if short == "" {
		http.Error(w, "url is empty", http.StatusNotFound)
		return
	}

	origin, err := dbrepo.GetOriginal(ctx, short)
	if err != nil {
		log.Printf("error GetOriginal %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(origin.URL)

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Location", origin.URL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
