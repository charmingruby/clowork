package command

import (
	"fmt"

	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/spf13/cobra"
)

func (c *Command) ListRoomMembers() *cobra.Command {
	var roomID string
	var page int

	cmd := &cobra.Command{
		Use: "members",
		RunE: func(cmd *cobra.Command, args []string) error {
			members, err := c.client.ListRoomMembers(&pb.ListRoomMembersRequest{
				Page:   int64(page),
				RoomId: roomID,
			})
			if err != nil {
				return err
			}

			for idx, m := range members {
				Print(
					fmt.Sprintf("%d. %s [%s]", idx+1, m.GetNickname(), m.GetHostname()),
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
