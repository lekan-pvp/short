package memhandlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func BenchmarkAPIShorten(b *testing.B) {
	b.Run("endpoint: POST /api/shorten", func(b *testing.B) {
		data := url.Values{}
		data.Set("url", "http://yandex.ru")
		r, _ := http.NewRequest("POST", "/api/shorten", strings.NewReader(data.Encode()))
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(APIShorten)

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			handler.ServeHTTP(w, r)
		}
	})
}
