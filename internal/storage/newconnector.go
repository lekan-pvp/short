package storage

import (
	"fmt"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/handlers"
	"github.com/lekan-pvp/short/internal/storage/dbrepo"
	"github.com/lekan-pvp/short/internal/storage/memrepo"
	"strings"
)

func NewConnector(cfg config.Config) handlers.Repo {
	switch {
	case strings.Contains(cfg.FileStoragePath, "storage.json"):
		return memrepo.New(cfg.FileStoragePath)
	case strings.Contains(cfg.DatabaseDSN, "postgresql"):
		return dbrepo.New(cfg.DatabaseDSN)
	default:
		fmt.Printf("unknown repo %s or %s", cfg.DatabaseDSN, cfg.FileStoragePath)
		return nil
	}
}
