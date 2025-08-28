package main

import (
	"context"
	"gcjade/services/catalogue-service/internal/infrastructure/grpc"
	"gcjade/services/catalogue-service/internal/infrastructure/repository"
	"gcjade/services/catalogue-service/internal/service"
	grpcsvr "google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var GrpcAddr = ":9092"

func main() {
	repo := repository.NewInmemRepository()
	categoryService := service.NewCategoryService(repo)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		cancel()
	}()

	lis, err := net.Listen("tcp", GrpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpcsvr.NewServer()
	grpc.NewHandler(grpcServer, categoryService)

	log.Printf("Starting Catalogue gRPC Service on port %s", lis.Addr().String())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to serve: %v", err)
			cancel()
		}
	}()

	<-ctx.Done()
	grpcServer.GracefulStop()
}
