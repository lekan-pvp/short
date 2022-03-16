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

func ExampleGetURLS() {
	router := chi.NewRouter()
	config.New()
	serverAddress := config.GetServerAddress()
	dbDSN := config.GetDatabaseURI()
	if dbDSN != "" {
		dbrepo.New()
		router.Get("/api/user/urls", GetURLS)
	}
	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func BenchmarkGetURLS(b *testing.B) {
	r, _ := http.NewRequest("GET", "/api/user/urls", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetURLS)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(w, r)
	}
}
