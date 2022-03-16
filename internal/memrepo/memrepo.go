package memrepo

import (
	"encoding/json"
	"errors"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/makeshort"
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

var urls = make([]Storage, 0)
var filePath string

func New() {
	var err error
	filePath = config.GetFilePath()
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
		var r Storage
		if err = d.Decode(&r); err == nil {
			urls = append(urls, r)
		}
	}
	if err == io.EOF {
		return
	}
}

func PostURL(url Storage) error {
	var err error
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer func() {
		if cerr := f.Close(); err != nil {
			err = cerr
		}
	}()
	if err != nil {
		return err
	}

	urls = append(urls, url)

	return json.NewEncoder(f).Encode(&url)
}

func GetOriginal(short string) (string, error) {
	log.Println("Get original IN MEM")
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

func GetURLsList(uuid string) ([]ListResponse, error) {
	base := config.GetBaseURL()
	var list = make([]ListResponse, 0)

	if len(urls) == 0 {
		return nil, errors.New("not found")
	}

	for _, v := range urls {
		if v.UUID == uuid {
			list = append(list, ListResponse{
				ShortURL:    base + "/" + v.ShortURL,
				OriginalURL: v.OriginalURL,
			})
		}
	}
	return list, nil
}

type BatchResponse struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

type BatchRequest struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

func BatchShorten(uuid string, in []BatchRequest) ([]BatchResponse, error) {
	base := config.GetBaseURL()
	var res = make([]BatchResponse, 0)
	for _, v := range in {
		short := makeshort.GenerateShortLink(v.OriginalURL, v.CorrelationID)
		res = append(res, BatchResponse{CorrelationID: v.CorrelationID, ShortURL: base + "/" + short})
		urls = append(urls, Storage{UUID: uuid, ShortURL: short, OriginalURL: v.OriginalURL, CorrelationID: v.CorrelationID})
	}
	return res, nil
}
