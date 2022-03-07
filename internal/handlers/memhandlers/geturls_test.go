package memhandlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
