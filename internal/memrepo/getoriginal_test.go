package memrepo

import "testing"

func TestMemoryRepo_GetOriginal(t *testing.T) {
	type fields struct {
		db []Storage
	}
	type args struct {
		short string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Success test",
			fields: fields{
				db: []Storage{
					{
						UUID:          "123",
						ShortURL:      "4rSPg8ap",
						OriginalURL:   "http://yandex.ru",
						CorrelationID: "2",
						DeleteFlag:    false,
					},
				},
			},
			args: args{
				short: "4rSPg8ap",
			},
			want:    "http://yandex.ru",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MemoryRepo{
				db: tt.fields.db,
			}
			got, err := r.GetOriginal(tt.args.short)
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
