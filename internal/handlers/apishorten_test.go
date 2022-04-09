package handlers

import (
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/storage/memrepo"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// Example using Chi router
func ExampleAPIShorten() {
	router := chi.NewRouter()
	config.New()
	serverAddress := config.Cfg.ServerAddress
	repo := memrepo.New(config.Cfg.FileStoragePath)
	router.Post("/api/shorten", APIShorten(repo))
	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func BenchmarkAPIShorten(b *testing.B) {
	b.Run("endpoint: POST /api/shorten", func(b *testing.B) {
		config.New()
		data := url.Values{}
		data.Set("url", "http://yandex.ru")
		r, _ := http.NewRequest("POST", "/api/shorten", strings.NewReader(data.Encode()))
		w := httptest.NewRecorder()
		repo := memrepo.New(config.Cfg.FileStoragePath)
		handler := APIShorten(repo)

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			handler.ServeHTTP(w, r)
		}
	})
}
