package memrepo

import (
	"context"
	"errors"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/models"
)

func (r *MemoryRepo) GetURLsList(_ context.Context, uuid string) ([]models.ListResponse, error) {
	base := config.Cfg.BaseURL
	var list []models.ListResponse

	if len(r.db) == 0 {
		return nil, errors.New("not found")
	}

	for _, v := range r.db {
		if v.UUID == uuid {
			list = append(list, models.ListResponse{
				ShortURL:    base + "/" + v.ShortURL,
				OriginalURL: v.OriginalURL,
			})
		}
	}
	return list, nil
}
