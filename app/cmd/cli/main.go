package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/charmingruby/clowork/internal/chat/delivery/command"
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

	client := client.New(clientConn)
	err = client.Stream(context.Background())
	if err != nil {
		os.Exit(1)
	}

	rootCmd := &cobra.Command{}

	cmdHandler := command.New(rootCmd, client)
	cmdHandler.Register()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		command.Print("", 0, false)
		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		args := strings.Fields(line)

		rootCmd.SetArgs(args)

		if err := rootCmd.Execute(); err != nil {
			command.Print(
				fmt.Sprintf("⚠️ Error: %s", err.Error()),
				0,
				true,
			)
		}
	}
}
