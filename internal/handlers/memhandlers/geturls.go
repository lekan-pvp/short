package memhandlers

import (
	"encoding/json"
	"github.com/lekan-pvp/short/internal/cookies"
	"github.com/lekan-pvp/short/internal/memrepo"
	"net/http"
	"strings"
)

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
