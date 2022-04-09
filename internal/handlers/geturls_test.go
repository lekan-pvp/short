package handlers

import (
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/storage/dbrepo"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ExampleGetURLs() {
	router := chi.NewRouter()
	config.New()
	serverAddress := config.Cfg.ServerAddress
	dbDSN := config.Cfg.DatabaseDSN
	config.New()
	if dbDSN != "" {
		dbRepo := dbrepo.New(config.Cfg)
		router.Get("/api/user/urls", GetURLs(&dbRepo))
	}
	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func BenchmarkGetURLs(b *testing.B) {
	r, _ := http.NewRequest("GET", "/api/user/urls", nil)
	w := httptest.NewRecorder()
	config.New()
	dbRepo := dbrepo.New(config.Cfg)
	handler := GetURLs(&dbRepo)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(w, r)
	}
}
