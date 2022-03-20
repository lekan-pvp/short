package memrepo

import (
	"context"
	"fmt"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoryRepo_GetURLsList(t *testing.T) {
	type fields struct {
		db []models.Storage
	}
	type args struct {
		in0  context.Context
		uuid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.ListResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success test",
			fields: fields{
				db: []models.Storage{
					{
						UUID:          "123",
						ShortURL:      "4rSPg8ap",
						OriginalURL:   "http://yandex.ru",
						CorrelationID: "",
						DeleteFlag:    false,
					},
					{
						UUID:          "123",
						ShortURL:      "e34rfdfsfd",
						OriginalURL:   "http://google.com",
						CorrelationID: "",
						DeleteFlag:    false,
					},
					{
						UUID:          "123",
						ShortURL:      "rtyf74sq",
						OriginalURL:   "http://educative.io",
						CorrelationID: "",
						DeleteFlag:    false,
					},
					{
						UUID:          "e1055b28-3f92-48be-a4db-e0d2ee4d1145",
						ShortURL:      "XBjSWzTD",
						OriginalURL:   "http://yanwwsdex.ru",
						CorrelationID: "",
						DeleteFlag:    false,
					},
				},
			},
			args: args{
				in0:  context.Background(),
				uuid: "123",
			},
			want: []models.ListResponse{
				{
					ShortURL:    "http://localhost:8080/4rSPg8ap",
					OriginalURL: "http://yandex.ru",
				},
				{
					ShortURL:    "http://localhost:8080/e34rfdfsfd",
					OriginalURL: "http://google.com",
				},
				{
					ShortURL:    "http://localhost:8080/rtyf74sq",
					OriginalURL: "http://educative.io",
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.Cfg.FileStoragePath = "test.json"
			config.Cfg.BaseURL = "http://localhost:8080"
			r := &MemoryRepo{
				db: tt.fields.db,
			}
			got, err := r.GetURLsList(tt.args.in0, tt.args.uuid)
			if !tt.wantErr(t, err, fmt.Sprintf("GetURLsList(%v, %v)", tt.args.in0, tt.args.uuid)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetURLsList(%v, %v)", tt.args.in0, tt.args.uuid)
		})
	}
}
