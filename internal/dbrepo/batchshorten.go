package dbrepo

import (
	"context"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/makeshort"
	"github.com/lekan-pvp/short/internal/models"
	"log"
)

// BatchShorten is a method accepting in the request body a set of URLs to shorten and returning a list of original URLs.
// Used in dbhandlers.
func (r *DBRepo) BatchShorten(ctx context.Context, uuid string, in []models.BatchRequest) ([]models.BatchResponse, error) {
	var res []models.BatchResponse
	base := config.Cfg.BaseURL

	log.Printf("in is %v", in)

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO users(user_id, short_url, orig_url, correlation_id) 
												VALUES($1, $2, $3, $4)`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	for _, v := range in {
		log.Printf("%s %s", v.OriginalURL, v.CorrelationID)
		short := makeshort.GenerateShortLink(v.OriginalURL, v.CorrelationID)
		log.Printf("short is %s", short)
		if _, err = stmt.ExecContext(ctx, uuid, short, v.OriginalURL, v.CorrelationID); err != nil {
			log.Println("insert exec error")
			return nil, err
		}
		res = append(res, models.BatchResponse{CorrelationID: v.CorrelationID, ShortURL: base + "/" + short})
	}
	if err := tx.Commit(); err != nil {
		log.Println("commit error")
		return nil, err
	}
	return res, nil
}
