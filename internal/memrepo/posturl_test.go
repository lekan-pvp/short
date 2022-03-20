package memrepo

import (
	"context"
	"fmt"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoryRepo_PostURL(t *testing.T) {
	type fields struct {
		db []models.Storage
	}
	type args struct {
		in0 context.Context
		rec models.Storage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success test",
			fields: fields{
				db: []models.Storage{
					{},
				},
			},
			args: args{
				in0: context.Background(),
				rec: models.Storage{
					UUID:          "b37284d5-78f9-4f3d-a1a8-bbe1a5e0c602",
					ShortURL:      "ehUadd7f",
					OriginalURL:   "http://google.com",
					CorrelationID: "123",
					DeleteFlag:    false,
				},
			},
			want:    "ehUadd7f",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.New()
			r := &MemoryRepo{
				db: tt.fields.db,
			}
			got, err := r.PostURL(tt.args.in0, tt.args.rec)
			if !tt.wantErr(t, err, fmt.Sprintf("PostURL(%v, %v)", tt.args.in0, tt.args.rec)) {
				return
			}
			assert.Equalf(t, tt.want, got, "PostURL(%v, %v)", tt.args.in0, tt.args.rec)
		})
	}
}
