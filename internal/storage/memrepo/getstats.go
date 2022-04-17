package memrepo

import (
	"context"
	"github.com/lekan-pvp/short/internal/models"
)

func (r *MemoryRepo) GetStats(_ context.Context) (models.Stat, error) {
	var urls int
	var users int
	var stats models.Stat
	for _, v := range r.db {
		if v.OriginalURL != "" {
			urls += 1
		}
		if v.UUID != "" {
			users += 1
		}
	}
	stats.URLs = urls
	stats.Users = users
	return stats, nil
}
