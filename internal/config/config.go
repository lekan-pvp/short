package config

import (
	"flag"
	"github.com/caarlos0/env"
	"log"
	"sync"
)

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"test.json"`
	DatabaseDSN     string `env:"DATABASE_DSN" envDefault:"user=postgres password='postgres' dbname=pqgotest sslmode=disable"`
}

var (
	config Config
)

func New() {
	var serverAddress string
	var databaseDSN string
	var fileStoragePath string
	var baseURL string

	var once sync.Once
	once.Do(func() {
		if err := env.Parse(&config); err != nil {
			log.Fatalln("can't parse Config")
		}

		flag.StringVar(&serverAddress, "a", config.ServerAddress, "адрес и порт запуска сервиса")
		flag.StringVar(&databaseDSN, "d", config.DatabaseDSN, "URI подключения к БД")
		flag.StringVar(&fileStoragePath, "f", config.FileStoragePath, "file storage")
		flag.StringVar(&baseURL, "b", config.BaseURL, "Base URL")

		flag.Parse()

		config.ServerAddress = serverAddress
		config.DatabaseDSN = databaseDSN
		config.FileStoragePath = fileStoragePath
		config.BaseURL = baseURL
	})
}

func GetFilePath() string {
	return config.FileStoragePath
}

func GetDatabaseURI() string {
	return config.DatabaseDSN
}

func GetBaseURL() string {
	return config.BaseURL
}

func GetServerAddress() string {
	return config.ServerAddress
}
