package main

import (
	_ "embed"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/dbrepo"
	"github.com/lekan-pvp/short/internal/handlers"
	"github.com/lekan-pvp/short/internal/memrepo"
	"github.com/lekan-pvp/short/internal/mware"
	"github.com/lekan-pvp/short/internal/pprofservoce"
	"log"
	"net/http"
)

//go:generate go run -ldflags "-X main.BuildVersion=v1.19.1 -X 'main.BuildDate=$(time +'%Y/%m/%d %H:%M:%S')'" main.go

var (
	BuildVersion = "N/A"
	BuildDate    = "N/A"
	BuildCommit  = "N/A"
)

func main() {
	fmt.Println("Build version: ", BuildVersion)
	fmt.Println("Build date: ", BuildDate)
	fmt.Println("Build commit: ", BuildCommit)

	config.New()

	if !config.Cfg.PprofEnabled {
		pprofservoce.PprofService()
	}

	serverAddress := config.Cfg.ServerAddress

	log.Println("\nServer address: ", serverAddress)

	router := chi.NewRouter()

	dbDSN := config.Cfg.DatabaseDSN

	router.Use(middleware.Logger)

	var memRepo memrepo.MemoryRepo
	var dbRepo dbrepo.DBRepo

	if dbDSN != "" {
		log.Println("dbrepo")
		dbRepo = dbrepo.New(config.Cfg)
		router.With(mware.Ping).Get("/ping", handlers.PingDB(&dbRepo))
		router.Post("/", handlers.PostURL(&dbRepo))
		router.Get("/{short}", handlers.GetShort(&dbRepo))
		router.Route("/api/shorten", func(r chi.Router) {
			r.Post("/", handlers.APIShorten(&dbRepo))
			r.Post("/batch", handlers.PostBatch(&dbRepo))
		})
		router.Route("/api/user", func(r chi.Router) {
			r.Delete("/urls", handlers.SoftDelete(&dbRepo))
			r.Get("/urls", handlers.GetURLs(&dbRepo))
		})
	} else {
		log.Println("memrepo")
		memRepo = memrepo.New(config.Cfg)
		router.With(mware.RequestHandle, mware.GzipHandle).Post("/", handlers.PostURL(&memRepo))
		router.With(mware.GzipHandle).Get("/{short}", handlers.GetShort(&memRepo))
		router.With(mware.RequestHandle, mware.GzipHandle).Post("/api/shorten", handlers.APIShorten(&memRepo))
		router.Get("/api/user/urls", handlers.GetURLs(&memRepo))
		router.Post("/api/shorten/batch", handlers.PostBatch(&memRepo))
	}

	err := http.ListenAndServe(serverAddress, router)
	if err != nil {
		log.Println("server error", err)
		panic(err)
	}
}
