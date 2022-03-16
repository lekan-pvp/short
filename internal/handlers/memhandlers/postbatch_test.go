package memhandlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

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

	handler := http.HandlerFunc(PostBatch)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(w, r)
	}
}
