package cookies

import (
	"net/http"
	"reflect"
	"testing"
)

func TestCreateCookie(t *testing.T) {
	tests := []struct {
		name string
		want *http.Cookie
	}{
		{
			name: "success test",
			want: CreateCookie(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.want
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCookie() = %v, want %v", got, tt.want)
			}
			if got.Path != tt.want.Path {
				t.Errorf("CreateCookie() = %v, want %v", got.Path, tt.want.Path)
			}
		})
	}
}

func TestCheckCookie(t *testing.T) {
	type args struct {
		cookie *http.Cookie
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success test",
			args: args{CreateCookie()},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckCookie(tt.args.cookie); got != tt.want {
				t.Errorf("CheckCookie() = %v, want %v", got, tt.want)
			}
		})
	}
}
