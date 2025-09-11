package command

import (
	"fmt"

	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/spf13/cobra"
)

func (c *Command) CreateRoom() *cobra.Command {
	cmd := &cobra.Command{
		Use: "room",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}

			topic, err := cmd.Flags().GetString("topic")
			if err != nil {
				return err
			}

			id, err := c.client.CreateRoom(&pb.CreateRoomRequest{
				Name:  name,
				Topic: topic,
			})
			if err != nil {
				return err
			}

			print(
				fmt.Sprintf("Room created successfully; id: %s", id),
				1,
				true,
				ResultSymbol,
			)

			return nil
		},
	}

	cmd.Flags().String("name", "", "Room name")
	cmd.Flags().String("topic", "", "Room Topic")

	return cmd
}
