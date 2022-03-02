package memhandlers

import (
	"encoding/json"
	"github.com/lekan-pvp/short/internal/cookies"
	"github.com/lekan-pvp/short/internal/memrepo"
	"net/http"
	"strings"
)

func PostBatch(w http.ResponseWriter, r *http.Request) {
	var uuid string

	var in []memrepo.BatchRequest

	cookie, err := r.Cookie("token")
	if err != nil || !cookies.CheckCookie(cookie) {
		cookie = cookies.CreateCookie()
	}

	http.SetCookie(w, cookie)

	values := strings.Split(cookie.Value, ":")
	uuid = values[0]

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := memrepo.BatchShorten(uuid, in)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	if err := json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
