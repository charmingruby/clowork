package command

import (
	"bufio"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func (c *Command) Auth() *cobra.Command {
	cmd := &cobra.Command{
		Use: "auth",
		RunE: func(cmd *cobra.Command, args []string) error {
			nickname, err := cmd.Flags().GetString("nickname")
			if err != nil {
				return err
			}

			if nickname == "" {
				Print("Enter your nickname: ", 1, false)

				reader := bufio.NewReader(os.Stdin)

				input, err := reader.ReadString('\n')
				if err != nil {
					return err
				}

				nickname = strings.TrimSpace(input)
			}

			hostname, err := os.Hostname()
			if err != nil {
				return err
			}

			c.session.Nickname = nickname
			c.session.Hostname = hostname

			return nil
		},
	}

	cmd.Flags().String("nickname", "", "Your nickname for the session")

	return cmd
}
