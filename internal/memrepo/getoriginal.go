package memrepo

import (
	"context"
	"errors"
	"github.com/lekan-pvp/short/internal/models"
	"log"
)

// GetOriginal is a function for getting original url fom memory repository or from json file
// by short url and user id. User must be authorized.
// Used in GetShort handler.
func (r *MemoryRepo) GetOriginal(_ context.Context, short string) (models.OriginURL, error) {
	for _, v := range r.db {
		log.Println(v)
	}
	log.Println("Get original IN MEM")
	var url models.OriginURL
	for _, v := range r.db {
		if short == v.ShortURL {
			url.URL = v.OriginalURL
			url.Deleted = v.DeleteFlag
			return url, nil
		}
	}
	return models.OriginURL{}, errors.New("URL not found")
}
