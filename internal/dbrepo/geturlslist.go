package dbrepo

import (
	"context"
	"github.com/lekan-pvp/short/internal/models"
)

// GetURLSList is a method for receive and form response in dbhandlers.GetURLS
// handler.
func (r DBRepo) GetURLsList(ctx context.Context, uuid string) ([]models.ListResponse, error) {
	var list []models.ListResponse

	rows, err := r.db.QueryContext(ctx, `SELECT short_url, orig_url FROM users WHERE user_id=$1`, uuid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var v models.ListResponse
		err = rows.Scan(&v.ShortURL, &v.OriginalURL)
		if err != nil {
			return nil, err
		}
		list = append(list, v)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return list, nil
}
