package server

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/short/internal/config"
	pb "github.com/lekan-pvp/short/internal/shortengrpc"
	"github.com/lekan-pvp/short/internal/shortengrpc/grpcserver"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg config.Config, router chi.Router) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	pb.RegisterShortenGrpcServer(s, &grpcserver.UsersServer{})

	srv := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: router,
	}

	ishttps := cfg.EnableHTTPS
	switch ishttps {
	case true:
		go func() {
			if err := srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile); err != nil {
				log.Fatalf("listen: %s\n", err)
			}
		}()
		go func() {
			if err := s.Serve(listen); err != nil {
				log.Fatal(err)
			}
		}()
	case false:
		go func() {
			if err := srv.ListenAndServe(); err != nil {
				log.Fatalf("listen: %s\n", err)
			}
		}()
		go func() {
			if err := s.Serve(listen); err != nil {
				log.Fatal(err)
			}
		}()
	}
	log.Println("server started")
	log.Println("grpc server started")

	<-done
	log.Println("server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %+v", err)
	}

	log.Print("server exited properly")

}
