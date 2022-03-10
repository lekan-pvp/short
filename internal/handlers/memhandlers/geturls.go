package memhandlers

import (
	"encoding/json"
	"github.com/lekan-pvp/short/internal/cookies"
	"github.com/lekan-pvp/short/internal/memrepo"
	"net/http"
	"strings"
)

// GetURLS is a handler that receives a list of Short and Original urls by user ID
//
// Endpoint: GET /api/user/urls
//
// Content-Type: application/json
//
// Response Example:
//
//  [
//    {
//        "short_url": "http://...",
//        "original_url": "http://..."
//    },
//    ...
//  ]
//
// Authorization by symmetric cookie.
//
// Possible response statuses:
// 200 OK
// 500 Internal Server Error
// 204 No Content
// 401 Unauthorized
func GetURLS(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil || !cookies.CheckCookie(cookie) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	http.SetCookie(w, cookie)

	values := strings.Split(cookie.Value, ":")
	if len(values) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	uuid := values[0]

	var list []memrepo.ListResponse

	list, err = memrepo.GetURLsList(uuid)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if list == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
