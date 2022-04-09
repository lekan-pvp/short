package handlers

import (
	"encoding/json"
	"github.com/lekan-pvp/short/internal/cookies"
	"io"
	"log"
	"net/http"
	"strings"
)

// SoftDelete is a asynchronous handler which accepts a list of short URL identifiers to remove in the format:
//
// Endpoint: DELETE /urls
//
//  [ "a", "b", "c", "d", ...]
//
// The user who created the URL can successfully delete the URL.
// When requesting a remote URL using the GET /{id} handler, you need to return the status 410 Gone.
//
// Possible response statuses:
// 202 Accepted it's OK
// 500 Internal Server Error
//
// TODO
// Handle 401 Unauthorized status
func SoftDelete(repo Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil || !cookies.CheckCookie(cookie) {
			cookie = cookies.New()
		}

		http.SetCookie(w, cookie)

		values := strings.Split(cookie.Value, ":")
		if len(values) != 2 {
			log.Println("cookie format error...")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		uuid := values[0]

		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("reading body error...")
			http.Error(w, err.Error(), 500)
			return
		}

		var in []string

		if err = json.Unmarshal(data, &in); err != nil {
			log.Println("decoding json error...")
			http.Error(w, err.Error(), 500)
			return
		}

		log.Println(in)

		if err = repo.SoftDelete(r.Context(), in, uuid); err != nil {
			log.Println("update db error")
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(http.StatusAccepted)
	}
}
