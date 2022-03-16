package memhandlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetShort(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetShort(tt.args.w, tt.args.r)
		})
	}
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
