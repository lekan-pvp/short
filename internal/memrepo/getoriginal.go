package memrepo

import (
	"context"
	"errors"
	"github.com/lekan-pvp/short/internal/models"
	"log"
)

func (r *MemoryRepo) GetOriginal(_ context.Context, short string) (*models.OriginURL, error) {
	log.Println("Get original IN MEM")
	var url *models.OriginURL
	for _, v := range r.db {
		if v.ShortURL == short {
			url.URL = v.OriginalURL
			url.Deleted = v.DeleteFlag
			return url, nil
		}
	}
	return nil, errors.New("URL not found")
}
