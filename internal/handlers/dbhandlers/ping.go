package dbhandlers

import (
	"context"
	"github.com/lekan-pvp/short/internal/dbrepo"
	"net/http"
)

// Ping handler checks database connection.
//
// Endpoint: GET /ping
//
// Content-Type: text/plain
//
// Possible response statuses:
// 200 OK
// 500 Internal Server Error
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
