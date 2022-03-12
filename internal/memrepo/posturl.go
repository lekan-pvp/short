package memrepo

import (
	"context"
	"encoding/json"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/models"
	"os"
)

func (r *MemoryRepo) PostURL(_ context.Context, rec models.Storage) (string, error) {
	var err error
	filePath := config.Cfg.FileStoragePath

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer func() {
		if cerr := f.Close(); err != nil {
			err = cerr
		}
	}()
	if err != nil {
		return "", err
	}

	r.db = append(r.db, rec)

	return rec.ShortURL, json.NewEncoder(f).Encode(&rec)
}
