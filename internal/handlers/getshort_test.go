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

func ExampleGetShort() {
	router := chi.NewRouter()
	config.New()
	serverAddress := config.Cfg.ServerAddress
	repo := memrepo.New(config.Cfg.FileStoragePath)
	router.Get("/{short}", GetShort(repo))
	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func BenchmarkGetShort(b *testing.B) {
	r, _ := http.NewRequest("GET", "/UZKV5qBG", nil)
	w := httptest.NewRecorder()
	config.New()
	repo := memrepo.New(config.Cfg.FileStoragePath)
	handler := GetShort(repo)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(w, r)
	}
}
