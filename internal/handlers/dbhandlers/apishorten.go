package dbhandlers

import (
	"context"
	"encoding/json"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/cookies"
	"github.com/lekan-pvp/short/internal/dbrepo"
	"github.com/lekan-pvp/short/internal/makeshort"
	"github.com/lekan-pvp/short/internal/memrepo"
	"log"
	"net/http"
	"strings"
)

func APIShorten(w http.ResponseWriter, r *http.Request) {
	ctx, stop := context.WithCancel(r.Context())
	defer stop()

	cookie, err := r.Cookie("token")
	if err != nil || !cookies.CheckCookie(cookie) {
		cookie = cookies.CreateCookie()
	}

	http.SetCookie(w, cookie)

	values := strings.Split(cookie.Value, ":")
	if len(values) != 2 {
		log.Println("Unauthorized")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	uuid := values[0]

	long := &memrepo.URL{}

	if err := json.NewDecoder(r.Body).Decode(long); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(long)

	short := makeshort.GenerateShortLink(long.URL, uuid)

	record := dbrepo.Storage{
		UUID:          uuid,
		ShortURL:      short,
		OriginalURL:   long.URL,
		CorrelationID: "123",
		DeleteFlag:    false,
	}

	err = dbrepo.PostURL(ctx, record)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	base := config.GetBaseURL()

	result := memrepo.ResultResponse{
		Result: base + "/" + short,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	if err := json.NewEncoder(w).Encode(&result); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
