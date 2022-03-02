package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/handlers"
	"github.com/lekan-pvp/short/internal/memrepo"
	"github.com/lekan-pvp/short/internal/mware"
	"log"
	"net/http"
)

func main() {
	config.New()

	err := memrepo.New()
	if err != nil {
		log.Println("Error initialization memrepo")
		panic(err)
	}

	serverAddress := config.GetServerAddress()

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.With(mware.RequestHandle, mware.GzipHandle).Post("/", handlers.PostURL)
	router.With(mware.GzipHandle).Get("/{short}", handlers.GetShort)
	router.With(mware.RequestHandle, mware.GzipHandle).Post("/api/shorten", handlers.APIShorten)
	router.Get("/api/user/urls", handlers.GetURLS)

	err = http.ListenAndServe(serverAddress, router)
	if err != nil {
		log.Println("server error", err)
		panic(err)
	}
}
