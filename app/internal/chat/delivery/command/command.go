package command

import (
	"fmt"

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
	c.cmd.AddCommand(
		c.Auth(),
	)
}

func Print(msg string, ident int, breakline bool) {
	var identation string

	if ident == 0 {
		identation = ">"
	} else {
		for i := range ident {
			identation += "  "

			if i == ident-1 {
				identation += ">"
				continue
			}
		}
	}

	if breakline {
		fmt.Printf("%s %s \n", identation, msg)
		return
	}

	fmt.Printf("%s %s", identation, msg)
}
