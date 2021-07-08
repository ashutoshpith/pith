package table

import (
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

func Table(flags []cli.Flag) *cli.Command {
	c := cli.Command {
		Name: "table",
		Usage: "table show",
		Aliases: []string{"t"},
		Flags: flags,
		Action: func(c *cli.Context) error {
			// value := c.String("text")
			pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
				{"Firstname", "Lastname", "Email"},
				{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
				{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
				{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
			}).Render()
			return nil
		},
	}
	
	return &c
}