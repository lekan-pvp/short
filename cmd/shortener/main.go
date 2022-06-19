package main

import (
	_ "embed"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/handlers"
	"github.com/lekan-pvp/short/internal/mware"
	"github.com/lekan-pvp/short/internal/pprofservice"
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
		pprofservice.PprofService()
	}

	serverAddress := config.Cfg.ServerAddress

	log.Println("\nServer address: ", serverAddress)

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	repo := storage.NewConnector(config.Cfg)

	router.With(mware.Ping).Get("/ping", handlers.PingDB(repo))
	router.With(mware.RequestHandle, mware.GzipHandle).Post("/", handlers.PostURL(repo))
	router.With(mware.GzipHandle).Get("/{short}", handlers.GetShort(repo))
	router.Route("/api/shorten", func(r chi.Router) {
		r.With(mware.RequestHandle, mware.GzipHandle).Post("/", handlers.APIShorten(repo))
		r.Post("/batch", handlers.PostBatch(repo))
	})
	router.Route("/api/user", func(r chi.Router) {
		r.Delete("/urls", handlers.SoftDelete(repo))
		r.Get("/urls", handlers.GetURLs(repo))
	})
	router.Get("/api/internal/stats", handlers.Stats(repo))

	server.Run(config.Cfg, router)
}
