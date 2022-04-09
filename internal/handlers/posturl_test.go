package handlers

import (
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/storage/dbrepo"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func ExamplePostURL() {
	router := chi.NewRouter()
	config.New()
	serverAddress := config.Cfg.ServerAddress
	dbDSN := config.Cfg.DatabaseDSN
	if dbDSN != "" {
		dbRepo := dbrepo.New(config.Cfg)
		router.Post("/", PostURL(&dbRepo))
	}
	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func BenchmarkPostURL(b *testing.B) {
	data := "http://yandex.ru"
	r, _ := http.NewRequest("POST", "/", strings.NewReader(data))
	w := httptest.NewRecorder()
	dbRepo := dbrepo.New(config.Cfg)
	handler := PostURL(&dbRepo)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(w, r)
	}
}
