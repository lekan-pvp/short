package memrepo

import (
	"encoding/json"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/models"
	"io"
	"log"
	"os"
)

type MemoryRepo struct {
	db []models.Storage
}

func New(cfg config.Config) *MemoryRepo {
	var err error
	var r *MemoryRepo
	filePath := cfg.FileStoragePath
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	defer func() {
		if cerr := f.Close(); cerr != nil {
			err = cerr
		}
	}()

	if err != nil {
		log.Println("open file error", err)
		panic(err)
	}

	if _, err := f.Seek(0, 0); err != nil {
		log.Fatalln("cant find file")
		panic(err)
	}
	d := json.NewDecoder(f)
	for err == nil {
		var row models.Storage
		if err = d.Decode(&row); err == nil {
			r.db = append(r.db, row)
		}
	}
	if err == io.EOF {
		return nil
	}
	return r
}
