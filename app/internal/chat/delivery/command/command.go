package command

import (
	"github.com/charmingruby/clowork/internal/chat/delivery/grpc/client"
	"github.com/spf13/cobra"
)

type Command struct {
	client  *client.Client
	session *session
	cmd     *cobra.Command
}

type session struct {
	Nickname string
	Hostname string
}

func New(cmd *cobra.Command, client *client.Client) *Command {
	return &Command{
		client:  client,
		session: &session{},
		cmd:     cmd,
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

	c.cmd.AddCommand(
		c.Auth(),
		listWrapper,
		createWrapper,
	)
}
