package command

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmingruby/clowork/internal/chat/delivery/cli"
	"github.com/spf13/cobra"
)

var ErrQuit = errors.New("quit signal")

func (c *Command) Chat() *cobra.Command {
	cmd := &cobra.Command{
		Use: "join",
		RunE: func(cmd *cobra.Command, args []string) error {
			roomID, err := cmd.Flags().GetString("room")
			if err != nil {
				return err
			}

			nickname, err := cmd.Flags().GetString("nickname")
			if err != nil {
				return err
			}

			hostname, err := os.Hostname()
			if err != nil {
				return err
			}

			ctx := context.Background()
			if err := c.client.ConnectStream(ctx); err != nil {
				return err
			}

			go func() {
				if err := c.client.Stream(); err != nil {
					cli.ReportCommandFailure(err)
					os.Exit(1)
				}
			}()

			if err := c.client.JoinRoom(roomID, nickname, hostname); err != nil {
				return err
			}

			cmdCh := make(chan string)

			go readInput(cmdCh)

			for {
				select {
				case msg := <-c.client.msgCh:
					cli.PreserveTyping()

					cli.Print(msg, 1, true, cli.ResultSymbol)

					cli.Cursor()

				case input := <-cmdCh:
					if err := c.handleInput(input); err != nil {
						return err
					}

				case <-time.After(100 * time.Millisecond):
					continue
				}
			}
		},
	}

	cmd.Flags().String("nickname", "dummy", "Nickname")
	cmd.Flags().String("room", "", "Room")

	return cmd
}

func readInput(cmdCh chan<- string) {
	scanner := bufio.NewScanner(os.Stdin)
	cli.Cursor()

	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input != "" {
			cmdCh <- input
		}
	}
}

func (c *Command) handleInput(input string) error {
	switch input {
	case "quit", "q":
		if err := c.client.LeaveRoom(); err != nil {
			return fmt.Errorf("%w: %w", ErrQuit, err)
		}

		return ErrQuit

	case "clear", "c":
		cli.Clear()
		cli.Cursor()

	default:
		if err := c.client.SendMessage(input); err != nil {
			cli.Print(err.Error(), 1, true, cli.FailureSymbol)
		}

		cli.Cursor()
	}

	return nil
}
