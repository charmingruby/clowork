package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/charmingruby/clowork/config"
	"github.com/charmingruby/clowork/internal/chat/delivery/grpc/client"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.NewCLI()
	if err != nil {
		failAndExit(nil)
	}

	grpcConn, err := grpc.NewClient(
		cfg.GRPCServerAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		failAndExit(nil)
	}

	client := client.New(grpcConn)
	err = client.Stream(context.TODO())
	if err != nil {
		failAndExit(grpcConn)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	signal := gracefulShutdown(grpcConn)

	os.Exit(signal)
}

func failAndExit(grpcConn *grpc.ClientConn) {
	gracefulShutdown(grpcConn)
	os.Exit(1)
}

func gracefulShutdown(grpcConn *grpc.ClientConn) int {
	var hasError bool

	if grpcConn != nil {
		if err := grpcConn.Close(); err != nil {
			hasError = true
		}
	}

	if hasError {
		return 1
	}

	return 0
}
