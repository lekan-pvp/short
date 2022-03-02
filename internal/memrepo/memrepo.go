package memrepo

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/lekan-pvp/short/internal/config"
	"io"
	"log"
	"os"
)

type Storage struct {
	UUID          string `json:"uuid"`
	ShortURL      string `json:"short_url"`
	OriginalURL   string `json:"original_url"`
	CorrelationID string `json:"correlation_id"`
	DeleteFlag    bool   `json:"delete_flag"`
}

type URL struct {
	URL string `json:"url"`
}

type ResultResponse struct {
	Result string `json:"result"`
}

var urls []Storage
var filePath string

func New() error {
	var err error
	filePath = config.GetFilePath()
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		log.Println("open file error", err)
		return err
	}

	if _, err := f.Seek(0, 0); err != nil {
		log.Fatalln("cant find file")
		return err
	}
	d := json.NewDecoder(f)
	for err == nil {
		var r Storage
		if err = d.Decode(&r); err == nil {
			urls = append(urls, r)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}

func PostURL(ctx context.Context, url Storage) error {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		return err
	}

	urls = append(urls, url)

	return json.NewEncoder(f).Encode(&url)
}

func GetOriginal(ctx context.Context, short string) (string, error) {
	var url string
	for _, v := range urls {
		if v.ShortURL == short {
			url = v.OriginalURL
			return url, nil
		}
	}
	return "", errors.New("URL not found")
}

type ListResponse struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func GetURLsList(ctx context.Context, uuid string) []ListResponse {
	base := config.GetBaseURL()
	var list []ListResponse
	for _, v := range urls {
		if v.UUID == uuid {
			list = append(list, ListResponse{
				ShortURL:    base + "/" + v.ShortURL,
				OriginalURL: v.OriginalURL,
			})
		}
	}
	return list
}
