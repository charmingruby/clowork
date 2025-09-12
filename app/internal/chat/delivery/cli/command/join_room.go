package command

import (
	"context"
	"os"

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

			if err := c.client.JoinRoom(roomID, nickname, hostname); err != nil {
				return err
			}

			go c.client.ListenToServerEvents()

			for {
				msg := <-c.client.console
				cli.Print(msg, 1, true, cli.ResultSymbol)
			}
		},
	}

	cmd.Flags().String("nickname", "dummy", "Nickname")
	cmd.Flags().String("room", "", "Room ID")

	return cmd
}
