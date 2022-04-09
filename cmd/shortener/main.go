package main

import (
	"context"
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
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	var err error

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

	isHTTPS := config.Cfg.EnableHTTPS
	certFile := config.Cfg.CertFile
	keyFile := config.Cfg.KeyFile

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	srv := &http.Server{
		Addr:    serverAddress,
		Handler: router,
	}

	if isHTTPS {
		go func() {
			if err = srv.ListenAndServeTLS(certFile, keyFile); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()
		log.Println("server started")
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("server started")

	<-done
	log.Println("server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %+v", err)
	}
	log.Print("server exited properly")
}
