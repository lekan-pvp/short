package dbrepo

import (
	"context"
	"github.com/jackc/pgerrcode"
	"github.com/lekan-pvp/short/internal/models"
	"github.com/lib/pq"
	"log"
)

// PostURL method for inserting short url and original url in database by user id.
// Used in dbhandlers.PostURL and dbhandlers.APIShorten handlers.
func (r DBRepo) PostURL(ctx context.Context, rec models.Storage) (string, error) {
	log.Println("IN InsertUserRepo short url =", rec.ShortURL)
	_, err := r.db.ExecContext(ctx, `INSERT INTO users(user_id, short_url, orig_url) VALUES ($1, $2, $3);`, rec.UUID, rec.ShortURL, rec.OriginalURL)

	var result string

	if err != nil {
		if err.(*pq.Error).Code == pgerrcode.UniqueViolation {
			notOk := r.db.QueryRowContext(ctx, `SELECT short_url FROM users WHERE orig_url=$1;`, rec.OriginalURL).Scan(&result)
			if notOk != nil {
				return "", notOk
			}
			return result, err
		}
	}

	return rec.ShortURL, nil
}
