package main

import (
	"os"

	"github.com/charmingruby/clowork/internal/chat/delivery/cli"
	"github.com/charmingruby/clowork/internal/chat/delivery/cli/command"
	"github.com/charmingruby/clowork/internal/chat/delivery/grpc/client"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddr, serverAddrExists := os.LookupEnv("CLOWORK_SERVER")
	if !serverAddrExists {
		os.Exit(1)
	}

	clientConn, err := grpc.NewClient(
		serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		os.Exit(1)
	}
	defer clientConn.Close()

	msgCh := make(chan string, 10)

	unaryCl, streamCl := client.New(clientConn, msgCh)

	rootCmd := &cobra.Command{
		Use: "Clowork",
	}

	cmdHandler := command.New(rootCmd, msgCh, unaryCl, streamCl)
	cmdHandler.Register()

	if err := rootCmd.Execute(); err != nil {
		cli.ReportFailure(err)
	}
}
