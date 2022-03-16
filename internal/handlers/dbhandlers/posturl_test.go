package dbhandlers

import (
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/dbrepo"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func ExamplePostURL() {
	router := chi.NewRouter()
	config.New()
	serverAddress := config.GetServerAddress()
	dbDSN := config.GetDatabaseURI()
	if dbDSN != "" {
		dbrepo.New()
		router.Post("/", PostURL)
	}
	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func BenchmarkPostURL(b *testing.B) {
	var baseURLs []string
	for i := 0; i < b.N; i++ {
		url := "http://yandex" + strconv.Itoa(i) + ".ru"
		baseURLs = append(baseURLs, url)
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r, _ := http.NewRequest("POST", "/", strings.NewReader(baseURLs[i]))
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(PostURL)

		handler.ServeHTTP(w, r)
	}
}
