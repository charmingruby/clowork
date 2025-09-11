package command

import (
	"fmt"

	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/spf13/cobra"
)

func (c *Command) CreateRoom() *cobra.Command {
	var roomName, roomTopic string

	cmd := &cobra.Command{
		Use: "room",
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := c.client.CreateRoom(&pb.CreateRoomRequest{
				Name:  roomName,
				Topic: roomTopic,
			})
			if err != nil {
				return err
			}

			print(
				fmt.Sprintf("Room created successfully; id:%s", id),
				1,
				true,
				ResultSymbol,
			)

			return nil
		},
	}

	cmd.Flags().StringVarP(&roomName, "name", "n", "", "Room name")
	cmd.Flags().StringVarP(&roomTopic, "topic", "t", "", "Room topic")

	return cmd
}
