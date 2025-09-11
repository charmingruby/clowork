package command

import (
	"github.com/charmingruby/clowork/internal/chat/delivery/grpc/client"
	"github.com/spf13/cobra"
)

type Command struct {
	client *client.Client
	cmd    *cobra.Command
}

func New(cmd *cobra.Command, client *client.Client) *Command {
	return &Command{
		client: client,
		cmd:    cmd,
	}
}

func (c *Command) Register() {
	listWrapper := &cobra.Command{
		Use: "list",
	}
	listWrapper.AddCommand(
		c.ListRooms(),
		c.ListMessages(),
		c.ListRoomMembers(),
	)

	createWrapper := &cobra.Command{
		Use: "create",
	}
	createWrapper.AddCommand(
		c.CreateRoom(),
	)

	chatWrapper := &cobra.Command{
		Use: "chat",
	}
	chatWrapper.AddCommand(
		c.JoinRoom(),
	)

	c.cmd.AddCommand(
		listWrapper,
		createWrapper,
		chatWrapper,
	)
}
