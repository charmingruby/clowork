package command

import (
	"fmt"

	"github.com/charmingruby/clowork/api/proto/pb"
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

			list(func() {
				for idx, m := range members {
					print(
						fmt.Sprintf("%d. %s [%s]", idx+1, m.GetNickname(), m.GetHostname()),
						1,
						true,
						ResultSymbol,
					)

					maybeSeparate(idx, len(members))
				}
			})

			return nil
		},
	}

	cmd.Flags().String("room", "", "Room ID")
	cmd.Flags().Int("page", 0, "Page")

	return cmd
}
