package handlers

import (
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/storage/memrepo"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ExampleGetURLs() {
	router := chi.NewRouter()
	config.New()
	serverAddress := config.Cfg.ServerAddress
	config.New()
	repo := memrepo.New(config.Cfg.FileStoragePath)
	router.Get("/api/user/urls", GetURLs(repo))

	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func BenchmarkGetURLs(b *testing.B) {
	r, _ := http.NewRequest("GET", "/api/user/urls", nil)
	w := httptest.NewRecorder()
	config.New()
	repo := memrepo.New(config.Cfg.ServerAddress)
	handler := GetURLs(repo)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(w, r)
	}
}
