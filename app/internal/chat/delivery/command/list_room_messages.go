package command

import (
	"fmt"

	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/spf13/cobra"
)

func (c *Command) ListMessages() *cobra.Command {
	var roomID string
	var page int

	cmd := &cobra.Command{
		Use: "messages",
		RunE: func(cmd *cobra.Command, args []string) error {
			messages, err := c.client.ListRoomMessages(&pb.ListRoomMessagesRequest{
				Page:   int64(page),
				RoomId: roomID,
			})
			if err != nil {
				return err
			}

			for _, m := range messages {
				Print(
					fmt.Sprintf("%s: %s", m.GetSenderId(), m.GetContent()),
					1,
					true,
					DefaultCommandType,
				)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&roomID, "room", "r", "", "Room id")
	cmd.Flags().IntVarP(&page, "page", "p", 0, "Page of members")

	return cmd
}
