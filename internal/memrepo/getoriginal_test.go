package memrepo

import (
	"context"
	"github.com/lekan-pvp/short/internal/models"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMemoryRepo_GetOriginal(t *testing.T) {
	type fields struct {
		db []models.Storage
	}
	type args struct {
		in0   context.Context
		short string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.OriginURL
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
				},
			},
			args: args{
				in0:   context.Background(),
				short: "4rSPg8ap",
			},
			want: models.OriginURL{
				URL:     "http://yandex.ru",
				Deleted: false,
			},
			wantErr: assert.NoError,
		},
		{
			name: "short url not found",
			fields: fields{
				db: []models.Storage{
					{
						UUID:          "123",
						ShortURL:      "nothing",
						OriginalURL:   "http://bigdick",
						CorrelationID: "",
						DeleteFlag:    false,
					},
				},
			},
			args: args{
				in0:   context.Background(),
				short: "4rSPg8ap",
			},
			want:    models.OriginURL{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MemoryRepo{
				db: tt.fields.db,
			}
			got, err := r.GetOriginal(tt.args.in0, tt.args.short)
			tt.wantErr(t, err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOriginal() got = %v, want %v", got, tt.want)
			}
		})
	}
}
