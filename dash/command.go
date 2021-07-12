package dash

import "github.com/urfave/cli/v2"

func Dashboard(flags []cli.Flag) *cli.Command {

	c := cli.Command {
		Name: "dash",
		Usage: "Dashboard",
		Aliases: []string{"d"},
		Flags: flags,
		Action: func(c *cli.Context) error {
			LoadUi()
	
			return nil
		},
	}
	
	return &c
}