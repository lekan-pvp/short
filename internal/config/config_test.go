package config

import "testing"

func TestGetFilePath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "success test",
			want: "test.json",
		},
	}
	New()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := GetFilePath(); got != tt.want {
				t.Errorf("GetFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDatabaseURI(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "success test",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := GetDatabaseURI(); got != tt.want {
				t.Errorf("GetDatabaseURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBaseURL(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "success test",
			want: "http://localhost:8080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBaseURL(); got != tt.want {
				t.Errorf("GetBaseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetServerAddress(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "success test",
			want: "localhost:8080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetServerAddress(); got != tt.want {
				t.Errorf("GetServerAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPprofStatus(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "success test pprof false",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPprofStatus(); got != tt.want {
				t.Errorf("GetPprofStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
