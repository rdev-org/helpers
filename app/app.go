package app

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

type ServerConfig struct {
	Name string
	Host string
	Port string
}

func RunGRPC(cfg ServerConfig, grpcServer *grpc.Server, cleanup func()) {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", cfg.Port, err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("%s gRPC server running on port %s...", cfg.Name, cfg.Port)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	sig := <-sigChan
	log.Printf("Received signal %v, shutting down...", sig)

	grpcServer.GracefulStop()
	if cleanup != nil {
		cleanup()
	}
	log.Println("Server stopped gracefully.")
}
