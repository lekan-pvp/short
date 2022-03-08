package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/dbrepo"
	"github.com/lekan-pvp/short/internal/handlers/dbhandlers"
	"github.com/lekan-pvp/short/internal/handlers/memhandlers"
	"github.com/lekan-pvp/short/internal/memrepo"
	"github.com/lekan-pvp/short/internal/mware"
	"github.com/lekan-pvp/short/internal/pprofservoce"
	"log"
	"net/http"
)

func main() {

	config.New()

	if config.GetPprofStatus() {
		pprofservoce.PprofService()
	}

	serverAddress := config.GetServerAddress()

	dbDSN := config.GetDatabaseURI()

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	if dbDSN != "" {
		dbrepo.New()
		router.With(mware.Ping).Get("/ping", dbhandlers.Ping)
		router.Post("/", dbhandlers.PostURL)
		router.Get("/{short}", dbhandlers.GetShort)
		router.Route("/api/shorten", func(r chi.Router) {
			r.Post("/", dbhandlers.APIShorten)
			r.Post("/batch", dbhandlers.PostBatch)
		})
		router.Get("/api/user/urls", dbhandlers.GetURLS)
		router.Route("/api/user", func(r chi.Router) {
			r.Delete("/urls", dbhandlers.SoftDelete)
		})
	} else {
		memrepo.New()
		router.With(mware.RequestHandle, mware.GzipHandle).Post("/", memhandlers.PostURL)
		router.With(mware.GzipHandle).Get("/{short}", memhandlers.GetShort)
		router.With(mware.RequestHandle, mware.GzipHandle).Post("/api/shorten", memhandlers.APIShorten)
		router.Get("/api/user/urls", memhandlers.GetURLS)
		router.Post("/api/shorten/batch", memhandlers.PostBatch)
	}

	err := http.ListenAndServe(serverAddress, router)
	if err != nil {
		log.Println("server error", err)
		panic(err)
	}
}
