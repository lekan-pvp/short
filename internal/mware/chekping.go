package mware

import (
	"log"
	"net/http"
)

func Checkping(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		if r.RequestURI == "/ping" {
			next.ServeHTTP(w, r)
		}

	})
}
