package memrepo

import (
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
		// TODO: Add test cases.
	}
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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

func TestBatchShorten(t *testing.T) {
	type args struct {
		uuid string
		in   []BatchRequest
	}
	tests := []struct {
		name    string
		args    args
		want    []BatchResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BatchShorten(tt.args.uuid, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("BatchShorten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchShorten() got = %v, want %v", got, tt.want)
			}
		})
	}
}
