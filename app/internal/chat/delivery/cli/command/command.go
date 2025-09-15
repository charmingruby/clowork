package command

import (
	"github.com/charmingruby/clowork/internal/chat/delivery/grpc/client/stream"
	"github.com/charmingruby/clowork/internal/chat/delivery/grpc/client/unary"
	"github.com/spf13/cobra"
)

type unaryClient = unary.Client
type streamClient = stream.Client

type Client struct {
	*unaryClient
	*streamClient

	msgCh chan string
}

type Command struct {
	client *Client
	cmd    *cobra.Command
}

func New(
	cmd *cobra.Command,
	msgCh chan string,
	unaryCl *unary.Client,
	streamCl *stream.Client,
) *Command {
	return &Command{
		client: &Client{
			msgCh:        msgCh,
			unaryClient:  unaryCl,
			streamClient: streamCl,
		},
		cmd: cmd,
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
		listWrapper,
		createWrapper,
		c.Chat(),
	)
}
