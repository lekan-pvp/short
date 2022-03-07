package dbhandlers

import (
	"context"
	"github.com/lekan-pvp/short/internal/dbrepo"
	"net/http"
)

// Ping handler checks database connection
func Ping(w http.ResponseWriter, r *http.Request) {
	ctx, stop := context.WithCancel(r.Context())
	defer stop()

	err := dbrepo.Ping(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
