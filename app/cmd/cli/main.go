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
		os.Exit(1)
	}

	grpcConn, err := grpc.NewClient(
		cfg.GRPCServerAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		os.Exit(1)
	}
	defer grpcConn.Close()

	client := client.New(grpcConn)
	err = client.Stream(context.TODO())
	if err != nil {
		os.Exit(1)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
