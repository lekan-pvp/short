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

func ExampleGetShort() {
	router := chi.NewRouter()
	config.New()
	serverAddress := config.GetServerAddress()
	dbDSN := config.GetDatabaseURI()
	if dbDSN != "" {
		dbrepo.New()
		router.Get("/{short}", GetShort)
	}
	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func BenchmarkGetShort(b *testing.B) {
	r, _ := http.NewRequest("GET", "/UZKV5qBG", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetShort)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(w, r)
	}
}
