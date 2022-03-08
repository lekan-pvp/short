package dbhandlers

import (
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/dbrepo"
	"log"
	"net/http"
)

// Example using Chi router
func ExampleAPIShorten() {
	router := chi.NewRouter()
	config.New()
	serverAddress := config.GetServerAddress()
	dbDSN := config.GetDatabaseURI()
	if dbDSN == "" {
		dbrepo.New()
		router.Post("/api/shorten", APIShorten)
	}
	log.Fatal(http.ListenAndServe(serverAddress, router))
}
