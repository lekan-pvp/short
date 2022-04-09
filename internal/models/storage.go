package models

// Storage is a main data type for store datas in database.
type Storage struct {
	UUID          string `json:"uuid" db:"user_id"`                  // UUID string ID of user from cookie
	ShortURL      string `json:"short_url" db:"short_url"`           // ShortURL string is a generated short url from original url
	OriginalURL   string `json:"original_url" db:"orig_url"`         // OriginalURL string is an original url for which generate short url
	CorrelationID string `json:"correlation_id" db:"correlation_id"` // CorrelationID string identifier
	DeleteFlag    bool   `json:"delete_flag" db:"is_deleted"`        // DeleteFlag is a flag for soft deleting in database
}
