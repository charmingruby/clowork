package main

import (
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
		panic(err)
	}

	conn, err := grpc.NewClient(
		cfg.GRPCServerAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	cl := client.New(conn)
	id, err := cl.CreateRoom()
	if err != nil {
		panic(err)
	}

	println(id)
}
