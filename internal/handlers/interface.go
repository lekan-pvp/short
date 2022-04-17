package handlers

import (
	"context"
	"github.com/lekan-pvp/short/internal/models"
)

// Repo interface which implements two types: MemoryRepo and DBRepo.
type Repo interface {
	PingDB(ctx context.Context) error
	PostURL(ctx context.Context, rec models.Storage) (string, error)
	GetOriginal(ctx context.Context, short string) (models.OriginURL, error)
	GetURLsList(ctx context.Context, uuid string) ([]models.ListResponse, error)
	BatchShorten(ctx context.Context, uuid string, in []models.BatchRequest) ([]models.BatchResponse, error)
	SoftDelete(ctx context.Context, in []string, uuid string) error
	GetStats(ctx context.Context) (models.Stat, error)
}
