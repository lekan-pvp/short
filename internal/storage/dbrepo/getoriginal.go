package dbrepo

import (
	"context"
	"github.com/lekan-pvp/short/internal/models"
	"log"
)

// GetOriginal is a method to get original url from database.
// Used in dbhandlers.GetShort handler.
func (r DBRepo) GetOriginal(ctx context.Context, short string) (models.OriginURL, error) {
	log.Println("GetOriginal IN DB")
	result := models.OriginURL{}

	err := r.db.QueryRowContext(ctx, `SELECT orig_url, is_deleted FROM users WHERE short_url=$1;`, short).Scan(&result.URL, &result.Deleted)
	if err != nil {
		return models.OriginURL{}, err
	}

	return result, nil
}
