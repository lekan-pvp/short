package dbhandlers

import (
	"encoding/json"
	"github.com/lekan-pvp/short/internal/cookies"
	"github.com/lekan-pvp/short/internal/dbrepo"
	"io"
	"log"
	"net/http"
	"strings"
)

func SoftDelete(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil || !cookies.CheckCookie(cookie) {
		cookie = cookies.CreateCookie()
	}

	http.SetCookie(w, cookie)

	values := strings.Split(cookie.Value, ":")
	if len(values) != 2 {
		log.Println("cookie format error...")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uuid := values[0]

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("reading body error...")
		http.Error(w, err.Error(), 500)
		return
	}

	log.Println(string(body))

	var in []string

	if err = json.Unmarshal(body, &in); err != nil {
		log.Println("decoding json error...")
		http.Error(w, err.Error(), 500)
		return
	}

	log.Println(in)

	if err = dbrepo.SoftDelete(r.Context(), in, uuid); err != nil {
		log.Println("update db error")
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(202)
}
