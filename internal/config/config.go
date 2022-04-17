package config

import (
	"encoding/json"
	"flag"
	"github.com/caarlos0/env"
	"log"
	"os"
	"sync"
)

// Config основной тип для конфигурации приложения
type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:"localhost:8080" json:"server_address"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080" json:"base_url"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"storage.json" json:"file_storage_path"`
	DatabaseDSN     string `env:"DATABASE_DSN" envDefault:"" json:"database_dsn"`
	PprofEnabled    bool   `json:"pprof_enabled"`
	EnableHTTPS     bool   `env:"ENABLE_HTTPS" envDefault:"false" json:"enable_https"`
	CertFile        string `json:"cert_file"`
	KeyFile         string `json:"key_file"`
	TrustedSubnet   string `env:"TRUSTED_SUBNET" envDefault:"" json:"trusted_subnet"`
	Config          string `env:"CONFIG" envDefault:""`
}

var Cfg Config

// New инициализирует все необходимые переменные для работы приложения.
// Используется синглтон.
func New() {
	var serverAddress string
	var databaseDSN string
	var fileStoragePath string
	var baseURL string
	var pprofEnabled bool
	var enableHTTPS bool
	var cfg string
	var trustedSubnet string

	var once sync.Once
	once.Do(func() {
		if err := env.Parse(&Cfg); err != nil {
			log.Fatal("can't parse Config", err)
		}

		flag.StringVar(&serverAddress, "a", Cfg.ServerAddress, "адрес и порт запуска сервиса")
		flag.StringVar(&databaseDSN, "d", Cfg.DatabaseDSN, "URI подключения к БД")
		flag.StringVar(&fileStoragePath, "f", Cfg.FileStoragePath, "file storage")
		flag.StringVar(&baseURL, "b", Cfg.BaseURL, "Base URL")
		flag.BoolVar(&pprofEnabled, "p", Cfg.PprofEnabled, "Pprof is enabled")
		flag.BoolVar(&enableHTTPS, "s", Cfg.EnableHTTPS, "To enable HTTPS server")
		flag.StringVar(&cfg, "c", Cfg.Config, "Read current config from file linker.json")
		flag.StringVar(&trustedSubnet, "t", Cfg.TrustedSubnet, "строковое представление бесклассовой адресации (CIDR)")

		flag.Parse()

		if cfg != "" {
			if err := getConfig(cfg); err != nil {
				log.Println("error: can't read config from file.")
				log.Fatal(err)
			}
		} else {
			Cfg.ServerAddress = serverAddress
			Cfg.DatabaseDSN = databaseDSN
			Cfg.FileStoragePath = fileStoragePath
			Cfg.BaseURL = baseURL
			Cfg.PprofEnabled = pprofEnabled
			Cfg.EnableHTTPS = enableHTTPS
			Cfg.CertFile = "certificate.crt"
			Cfg.KeyFile = "key.key"
			Cfg.TrustedSubnet = trustedSubnet
		}
	})
}

func getConfig(name string) error {
	f, err := os.OpenFile(name, os.O_RDONLY, 0644)
	if err != nil {
		log.Println("error: can't open file")
		return err
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(&Cfg); err != nil {
		log.Println("error: can't decode json file")
		return err
	}
	return nil
}
