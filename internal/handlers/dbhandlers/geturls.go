package dbhandlers

import (
	"context"
	"encoding/json"
	"github.com/lekan-pvp/short/internal/cookies"
	"github.com/lekan-pvp/short/internal/dbrepo"
	"net/http"
	"strings"
)

func GetURLS(w http.ResponseWriter, r *http.Request) {
	ctx, stop := context.WithCancel(r.Context())
	defer stop()

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

	var list []dbrepo.ListResponse

	list, err = dbrepo.GetURLsList(ctx, uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
