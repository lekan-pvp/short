package handlers

import (
	"context"
	"encoding/json"
	"github.com/lekan-pvp/short/internal/checkip"
	"net/http"
)

func Stats(repo Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if ok, err := checkip.CheckIP(r); err != nil || !ok {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		ctx, stop := context.WithCancel(r.Context())
		defer stop()
		stats, err := repo.GetStats(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(&stats); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
