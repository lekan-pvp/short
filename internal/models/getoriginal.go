package models

type URL struct {
	URL string `json:"url"`
}

// OriginURL is a struct for returning result from GetOriginal method and for making json response body in
// dbhandlers.GetShort handler.
type OriginURL struct {
	URL     string
	Deleted bool
}

// IsDeleted is a method for check delete flag in database.
func (u OriginURL) IsDeleted() bool {
	return u.Deleted
}
