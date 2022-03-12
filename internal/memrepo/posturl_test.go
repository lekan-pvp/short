package memrepo

import "testing"

func TestMemoryRepo_PostURL(t *testing.T) {
	type fields struct {
		db []Storage
	}
	type args struct {
		url Storage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
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
						CorrelationID: "1",
						DeleteFlag:    false,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MemoryRepo{
				db: tt.fields.db,
			}
			if err := r.PostURL(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("PostURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
