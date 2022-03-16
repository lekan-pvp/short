package mware

import "net/http"

func Ping(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/ping" {
			next.ServeHTTP(w, r)
		}
	})
}
