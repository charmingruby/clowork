package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/charmingruby/clowork/internal/chat/delivery/command"
	"github.com/charmingruby/clowork/internal/chat/delivery/grpc/client"
	"github.com/chzyer/readline"
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

	rootCmd := &cobra.Command{
		Use: "Clowork",
	}

	cmdHandler := command.New(rootCmd, client)
	cmdHandler.Register()

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          fmt.Sprintf("%s ", command.InputSymbol),
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		os.Exit(1)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		args := strings.Fields(line)

		rootCmd.SetArgs(args)
		rootCmd.SilenceErrors = true
		rootCmd.SilenceUsage = true

		if err := rootCmd.Execute(); err != nil {
			command.ReportFailure(err)
		}
	}
}
