package models

// BatchResponse is a struct for encode response body.
// Used in dbhandlers.GetURLS handler.
type BatchResponse struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

// BatchRequest is a struct for decode request body and inserting data to database.
// Used in dbhandlers.GetURLS.
type BatchRequest struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}
