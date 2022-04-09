package memrepo

import (
	"context"
	"encoding/json"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/models"
	"log"
	"os"
)

// PostURL is a function for saving generated short url, original url, user id and delete flag
// in memory and in json file.
// Used in hendlers: PostURL and APIShorten.
func (r *MemoryRepo) PostURL(_ context.Context, rec models.Storage) (string, error) {
	for _, v := range r.db {
		log.Println("in posturl: ", v)
	}

	log.Println("in mem postURL")
	var err error
	filePath := config.Cfg.FileStoragePath

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer func() {
		if cerr := f.Close(); err != nil {
			log.Println("error defer func", cerr)
			err = cerr
		}
	}()
	if err != nil {
		log.Println("error open file", err)
		return "", err
	}

	r.db = append(r.db, rec)

	log.Println(rec.ShortURL)

	return rec.ShortURL, json.NewEncoder(f).Encode(rec)
}
