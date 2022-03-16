package dbhandlers

import (
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/dbrepo"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ExamplePing() {
	router := chi.NewRouter()
	config.New()
	serverAddress := config.GetServerAddress()
	dbDSN := config.GetDatabaseURI()
	if dbDSN != "" {
		dbrepo.New()
		router.Get("/ping", Ping)
	}
	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func BenchmarkPing(b *testing.B) {
	config.New()
	dbrepo.New()
	r, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(Ping)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(w, r)
	}
}
