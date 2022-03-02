package handlers

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/short/internal/memrepo"
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

	url, err := memrepo.GetOriginal(ctx, short)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
