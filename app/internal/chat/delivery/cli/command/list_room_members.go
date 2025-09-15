package command

import (
	"fmt"

	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/internal/chat/delivery/cli"
	"github.com/spf13/cobra"
)

func (c *Command) ListRoomMembers() *cobra.Command {
	cmd := &cobra.Command{
		Use: "members",
		RunE: func(cmd *cobra.Command, args []string) error {
			room, err := cmd.Flags().GetString("room")
			if err != nil {
				return err
			}

			page, err := cmd.Flags().GetInt("page")
			if err != nil {
				return err
			}

			members, err := c.client.ListRoomMembers(&pb.ListRoomMembersRequest{
				Page:   int64(page),
				RoomId: room,
			})
			if err != nil {
				return err
			}

			for _, m := range members {
				cli.Print(
					fmt.Sprintf("[%s] %s (%s)", m.GetId(), m.GetNickname(), m.GetHostname()),
					0,
					true,
					cli.ResultSymbol,
				)
			}

			return nil
		},
	}

	cmd.Flags().String("room", "", "Room")
	cmd.Flags().Int("page", 0, "Page")

	return cmd
}
