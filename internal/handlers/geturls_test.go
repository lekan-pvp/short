package handlers

import (
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/dbrepo"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ExampleGetURLS() {
	router := chi.NewRouter()
	config.New()
	serverAddress := config.Cfg.ServerAddress
	dbDSN := config.Cfg.DatabaseDSN
	config.New()
	if dbDSN != "" {
		dbRepo := dbrepo.New(config.Cfg)
		router.Get("/api/user/urls", GetURLs(dbRepo))
	}
	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func BenchmarkGetURLS(b *testing.B) {
	r, _ := http.NewRequest("GET", "/api/user/urls", nil)
	w := httptest.NewRecorder()
	config.New()
	dbRepo := dbrepo.New(config.Cfg)
	handler := GetURLs(dbRepo)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(w, r)
	}
}
