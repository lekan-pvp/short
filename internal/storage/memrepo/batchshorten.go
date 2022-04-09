package memrepo

import (
	"context"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/makeshort"
	"github.com/lekan-pvp/short/internal/models"
)

// BatchShorten is a function for save an array of short urls in memory and in json file.
// Used in PostBatch handler.
func (r MemoryRepo) BatchShorten(_ context.Context, uuid string, in []models.BatchRequest) ([]models.BatchResponse, error) {
	base := config.Cfg.BaseURL
	var res []models.BatchResponse
	for _, v := range in {
		short := makeshort.GenerateShortLink(v.OriginalURL, v.CorrelationID)
		res = append(res, models.BatchResponse{CorrelationID: v.CorrelationID, ShortURL: base + "/" + short})
		r.db = append(r.db, models.Storage{UUID: uuid, ShortURL: short, OriginalURL: v.OriginalURL, CorrelationID: v.CorrelationID, DeleteFlag: false})
	}
	return res, nil
}
