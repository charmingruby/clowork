package command

import (
	"fmt"
	"os"

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

			hostname, err := os.Hostname()
			if err != nil {
				return err
			}

			c.session.Nickname = nickname
			c.session.Hostname = hostname

			print(
				fmt.Sprintf("Authenticated successfully (%s:%s)", c.session.Nickname, c.session.Hostname),
				1,
				true,
				ResultSymbol,
			)

			return nil
		},
	}

	cmd.Flags().String("nickname", "", "Your nickname for the session")

	return cmd
}
