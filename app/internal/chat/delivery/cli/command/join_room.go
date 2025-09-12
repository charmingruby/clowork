package command

import (
	"bufio"
	"context"
	"os"
	"strings"
	"time"

	"github.com/charmingruby/clowork/internal/chat/delivery/cli"
	"github.com/spf13/cobra"
)

func (c *Command) JoinRoom() *cobra.Command {
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

			go c.client.ListenToServerEvents()

			if err := c.client.JoinRoom(roomID, nickname, hostname); err != nil {
				return err
			}

			history := []string{}
			cmdCh := make(chan string)

			go func() {
				scanner := bufio.NewScanner(os.Stdin)

				cli.Cursor()

				for scanner.Scan() {
					input := strings.TrimSpace(scanner.Text())

					if input != "" {
						cmdCh <- input
					}
				}
			}()

			for {
				select {
				case msg := <-c.client.msgCh:
					cli.PreserveTyping()

					cli.Print(msg, 1, true, cli.ResultSymbol)

					history = append(history, msg)
					if len(history) > 50 {
						history = history[1:]
					}

					cli.Cursor()

				case input := <-cmdCh:
					switch input {
					case "quit", "q":
						return nil
					case "history", "h":
						cli.List(func() {
							start := max(len(history)-10, 0)

							for i := start; i < len(history); i++ {
								if i == len(history)-1 {
									cli.Print(
										history[i],
										1,
										false,
										cli.ResultSymbol,
									)
								}

								cli.Print(
									history[i],
									1,
									true,
									cli.ResultSymbol,
								)
							}
						})

						cli.Cursor()
					case "clear", "c":
						cli.Clear()
						cli.Cursor()
					default:
						// send msg
						cli.Cursor()
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
