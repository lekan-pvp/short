package handlers

import (
	"context"
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
func PingDB(repo Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, stop := context.WithCancel(r.Context())
		defer stop()

		err := repo.PingDB(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
	}
}
