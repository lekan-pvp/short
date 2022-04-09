package models

// ListResponse is a struct to  form response body in
// dbhandlers.GetURLS handler.
type ListResponse struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}
