package memhandlers

import (
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/short/internal/memrepo"
	"log"
	"net/http"
)

// GetShort is a handler that receives original URL by short URL.
//
// Endpoint:
// GET /{short}
//
// where {short} is a short URL
//
// Content-Type: text/plain
//
// Possible response statuses:
// 307 Temporary Redirect - Success
// 400 Bad Request
// 404 Not Found if record not found in memory or file
func GetShort(w http.ResponseWriter, r *http.Request) {
	short := chi.URLParam(r, "short")
	if short == "" {
		http.Error(w, "url is empty", http.StatusNotFound)
		return
	}

	url, err := memrepo.GetOriginal(short)
	if err != nil {
		log.Println("GetOriginal error in MEM")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(url)
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
