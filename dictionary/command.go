package dictionary

import (
	"github.com/urfave/cli/v2"
)


func Dict(flags []cli.Flag) *cli.Command {

	c := cli.Command {
		Name: "dict",
		Usage: "Dictionary",
		Aliases: []string{"di"},
		Flags: flags,
		Action: func(c *cli.Context) error {
			value := c.String("text1")
			dicti := Dicti{
				Query: value,
			}
			data, _ := dicti.Search()
			dicti.Iterate(*data)
	
			return nil
		},
	}
	
	return &c
}