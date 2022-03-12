package memrepo

import (
	"github.com/lekan-pvp/short/internal/config"
	"reflect"
	"testing"
)

func TestMemoryRepo_GetURLsList(t *testing.T) {
	type fields struct {
		db []Storage
	}
	type args struct {
		uuid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []ListResponse
		wantErr bool
	}{
		{
			name: "Success test",
			fields: fields{
				db: []Storage{
					{
						UUID:          "1234",
						ShortURL:      "4rSPg8ap",
						OriginalURL:   "http://yandex.ru",
						CorrelationID: "4",
						DeleteFlag:    false,
					},
				},
			},
			args: args{
				uuid: "1234",
			},
			want: []ListResponse{
				{
					ShortURL:    "4rSPg8ap",
					OriginalURL: "http://yandex.ru",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.New()
			r := &MemoryRepo{
				db: tt.fields.db,
			}
			got, err := r.GetURLsList(tt.args.uuid)
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
