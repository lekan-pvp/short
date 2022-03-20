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

func ExamplePingDB() {
	router := chi.NewRouter()
	config.New()
	serverAddress := config.Cfg.ServerAddress
	dbDSN := config.Cfg.DatabaseDSN

	if dbDSN != "" {
		dbRepo := dbrepo.New(config.Cfg)
		router.Get("/ping", PingDB(&dbRepo))
	}
	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func BenchmarkPing(b *testing.B) {
	r, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	config.New()
	dbRepo := dbrepo.New(config.Cfg)
	handler := GetShort(&dbRepo)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(w, r)
	}
}
