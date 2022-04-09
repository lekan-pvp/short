package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/storage"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func ExamplePostBatch() {
	router := chi.NewRouter()
	config.New()
	serverAddress := config.Cfg.ServerAddress
	repo := storage.NewConnector(config.Cfg)
	router.Post("/api/shorten/batch", PostBatch(repo))

	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func BenchmarkPostBatch(b *testing.B) {
	var datas []url.Values
	for i := 0; i < 5; i++ {
		data := url.Values{}
		data.Set("correlation_id", strconv.Itoa(i))
		data.Set("original_url", "http://yandex.ru")
		datas = append(datas, data)
	}

	body, _ := json.Marshal(datas)
	r, _ := http.NewRequest("POST", "/api/shorten/batch", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	config.New()
	repo := storage.NewConnector(config.Cfg)
	handler := PostBatch(repo)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(w, r)
	}
}
