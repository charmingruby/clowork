package command

import (
	"fmt"

	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/internal/chat/delivery/cli"
	"github.com/spf13/cobra"
)

func (c *Command) ListMessages() *cobra.Command {
	cmd := &cobra.Command{
		Use: "messages",
		RunE: func(cmd *cobra.Command, args []string) error {
			room, err := cmd.Flags().GetString("room")
			if err != nil {
				return err
			}

			page, err := cmd.Flags().GetInt("page")
			if err != nil {
				return err
			}

			messages, err := c.client.ListRoomMessages(&pb.ListRoomMessagesRequest{
				Page:   int64(page),
				RoomId: room,
			})
			if err != nil {
				return err
			}

			cli.List(func() {
				for _, m := range messages {
					cli.Print(
						fmt.Sprintf("%s: %s", m.GetSenderId(), m.GetContent()),
						1,
						true,
						cli.ResultSymbol,
					)
				}
			})

			return nil
		},
	}

	cmd.Flags().String("room", "", "Room")
	cmd.Flags().Int("page", 0, "Page")

	return cmd
}
