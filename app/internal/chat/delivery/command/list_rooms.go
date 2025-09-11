package command

import (
	"fmt"

	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/spf13/cobra"
)

func (c *Command) ListRooms() *cobra.Command {
	var page int

	cmd := &cobra.Command{
		Use: "rooms",
		RunE: func(cmd *cobra.Command, args []string) error {
			rooms, err := c.client.ListRooms(&pb.ListRoomsRequest{
				Page: int64(page),
			})
			if err != nil {
				return err
			}

			list(func() {
				for idx, r := range rooms {
					print(
						fmt.Sprintf("%d. [%s] %s", idx+1, r.GetTopic(), r.GetName()),
						1,
						true,
						ResultSymbol,
					)

					print(
						fmt.Sprintf("ID: %s", r.GetId()),
						2,
						true,
						ResultSymbol,
					)

					maybeSeparate(idx, len(rooms))
				}
			})

			return nil
		},
	}

	cmd.Flags().IntVarP(&page, "page", "p", 0, "Page of rooms")

	return cmd
}
