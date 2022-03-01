package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/handlers"
	"github.com/lekan-pvp/short/internal/memrepo"
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

	router.Post("/", handlers.PostURL)
	router.Get("/{short}", handlers.GetShort)
	router.Post("/api/shorten", handlers.APIShorten)

	err = http.ListenAndServe(serverAddress, router)
	if err != nil {
		log.Println("server error", err)
		panic(err)
	}
}
