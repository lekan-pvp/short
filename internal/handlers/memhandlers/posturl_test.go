package memhandlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostURL(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PostURL(tt.args.w, tt.args.r)
		})
	}
}

func BenchmarkPostURL(b *testing.B) {
	data := "http://yandex.ru"
	r, _ := http.NewRequest("POST", "/", strings.NewReader(data))
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(PostURL)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(w, r)
	}
}
