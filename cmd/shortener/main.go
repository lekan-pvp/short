package main

import (
	_ "embed"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/handlers"
	"github.com/lekan-pvp/short/internal/mware"
	"github.com/lekan-pvp/short/internal/pprofservoce"
	"github.com/lekan-pvp/short/internal/server"
	"github.com/lekan-pvp/short/internal/storage"
	"log"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	fmt.Println("Build version: ", buildVersion)
	fmt.Println("Build date: ", buildDate)
	fmt.Println("Build commit: ", buildCommit)

	config.New()

	if config.Cfg.PprofEnabled {
		pprofservoce.PprofService()
	}

	serverAddress := config.Cfg.ServerAddress

	log.Println("\nServer address: ", serverAddress)

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	repo := storage.NewConnector(config.Cfg)

	router.With(mware.Ping).Get("/ping", handlers.PingDB(repo))
	router.Post("/", handlers.PostURL(repo))
	router.Get("/{short}", handlers.GetShort(repo))
	router.Route("/api/shorten", func(r chi.Router) {
		r.Post("/", handlers.APIShorten(repo))
		r.Post("/batch", handlers.PostBatch(repo))
	})
	router.Route("/api/user", func(r chi.Router) {
		r.Delete("/urls", handlers.SoftDelete(repo))
		r.Get("/urls", handlers.GetURLs(repo))
	})

	server.Run(config.Cfg, router)
}
