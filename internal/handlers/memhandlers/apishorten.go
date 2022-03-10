package memhandlers

import (
	"encoding/json"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/cookies"
	"github.com/lekan-pvp/short/internal/makeshort"
	"github.com/lekan-pvp/short/internal/memrepo"
	"log"
	"net/http"
	"strings"
)

// APIShorten is a handler to make short URL and save them in memory and file.
//
// Endpoint:
// /api/shorten [post]
//
// Content-Type: application/json
//
// Request body example:
//
//  {
//    "url": "http://google.com"
//  }
//
// "url" is an original URL for making a short URL for one
//
// Possible response statuses:
//
// 201 Created Success status
// 500 Internal server error
// 401 Unauthorized
// 400 Bed Request
// 409 Status Conflict
func APIShorten(w http.ResponseWriter, r *http.Request) {
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

	record := memrepo.Storage{
		UUID:          uuid,
		ShortURL:      short,
		OriginalURL:   long.URL,
		CorrelationID: "123",
		DeleteFlag:    false,
	}

	err = memrepo.PostURL(record)
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
