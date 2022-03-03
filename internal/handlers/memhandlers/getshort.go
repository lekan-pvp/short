package memhandlers

import (
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/short/internal/memrepo"
	"log"
	"net/http"
)

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

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
