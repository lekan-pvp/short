package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/makeshort"
	"github.com/lekan-pvp/short/internal/storage/dbrepo"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func ExampleSoftDelete() {
	router := chi.NewRouter()
	config.New()
	serverAddress := config.Cfg.ServerAddress
	dbDSN := config.Cfg.DatabaseDSN
	if dbDSN != "" {
		dbRepo := dbrepo.New(config.Cfg)
		router.Delete("/urls", PostURL(&dbRepo))
	}
	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func BenchmarkSoftDelete(b *testing.B) {
	url := "http://yandex.ru"
	var datas []string
	for i := 0; i < b.N; i++ {
		data := makeshort.GenerateShortLink(url, strconv.Itoa(i))
		datas = append(datas, data)
	}

	body, _ := json.Marshal(datas)
	r, _ := http.NewRequest("DELETE", "/urls", strings.NewReader(string(body)))
	w := httptest.NewRecorder()

	config.New()

	dbRepo := dbrepo.New(config.Cfg)

	handler := SoftDelete(&dbRepo)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(w, r)
	}
}
