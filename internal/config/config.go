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
	DatabaseDSN     string `env:"DATABASE_DSN" envDefault:""`
	PprofEnabled    bool
}

var (
	Cfg *Config
)

func New() {
	var serverAddress string
	var databaseDSN string
	var fileStoragePath string
	var baseURL string
	var pprofEnabled bool

	var once sync.Once
	once.Do(func() {
		if err := env.Parse(&Cfg); err != nil {
			log.Fatalln("can't parse Config")
		}

		flag.StringVar(&serverAddress, "a", Cfg.ServerAddress, "адрес и порт запуска сервиса")
		flag.StringVar(&databaseDSN, "d", Cfg.DatabaseDSN, "URI подключения к БД")
		flag.StringVar(&fileStoragePath, "f", Cfg.FileStoragePath, "file storage")
		flag.StringVar(&baseURL, "b", Cfg.BaseURL, "Base URL")
		flag.BoolVar(&pprofEnabled, "p", Cfg.PprofEnabled, "Pprof is enabled")

		flag.Parse()

		Cfg.ServerAddress = serverAddress
		Cfg.DatabaseDSN = databaseDSN
		Cfg.FileStoragePath = fileStoragePath
		Cfg.BaseURL = baseURL
		Cfg.PprofEnabled = pprofEnabled

	})
}
