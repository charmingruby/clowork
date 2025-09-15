package command

import (
	"fmt"

	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/internal/chat/delivery/cli"
	"github.com/spf13/cobra"
)

func (c *Command) ListRooms() *cobra.Command {
	cmd := &cobra.Command{
		Use: "rooms",
		RunE: func(cmd *cobra.Command, args []string) error {
			page, err := cmd.Flags().GetInt("page")
			if err != nil {
				return err
			}

			rooms, err := c.client.ListRooms(&pb.ListRoomsRequest{
				Page: int64(page),
			})
			if err != nil {
				return err
			}

			for _, r := range rooms {
				cli.Print(
					fmt.Sprintf("[%s] %s (%s)", r.GetId(), r.GetTopic(), r.GetName()),
					0,
					true,
					cli.ResultSymbol,
				)
			}

			return nil
		},
	}

	cmd.Flags().Int("page", 0, "Page")

	return cmd
}
