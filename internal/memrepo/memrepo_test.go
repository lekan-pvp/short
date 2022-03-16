package memrepo

import (
	"github.com/lekan-pvp/short/internal/config"
	"reflect"
	"testing"
)

func TestPostURL(t *testing.T) {
	type args struct {
		url Storage
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success test",
			args: args{
				url: Storage{
					UUID:        "123",
					ShortURL:    "4rSPg8ap",
					OriginalURL: "http://yandex.ru",
				},
			},
		},
	}
	config.New()
	New()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PostURL(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("PostURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOriginal(t *testing.T) {
	type args struct {
		short string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success test",
			args: args{
				short: "4rSPg8ap",
			},
			want: "http://yandex.ru",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetOriginal(tt.args.short)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOriginal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetOriginal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetURLsList(t *testing.T) {
	type args struct {
		uuid string
	}
	tests := []struct {
		name    string
		args    args
		want    []ListResponse
		wantErr bool
	}{
		{
			name: "success test",
			args: args{
				uuid: "123",
			},
			want: []ListResponse{
				{
					ShortURL:    "http://localhost:8080/4rSPg8ap",
					OriginalURL: "http://yandex.ru",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetURLsList(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetURLsList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetURLsList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
