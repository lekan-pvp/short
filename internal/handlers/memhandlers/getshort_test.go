package memhandlers

import (
	"net/http"
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
