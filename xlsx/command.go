package xlsx

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Xlsx(flags [] cli.Flag) *cli.Command {
	c := cli.Command{
		Name: "xlsx",
			Aliases: []string{"xl"},
			Usage: "Xlsx ",
			Flags: flags,
			Subcommands: []*cli.Command{
				{
					Name: "count",
					Usage: "Count The Sheet",
					Flags: flags,
					Action: func(c *cli.Context) error {
						text := c.String("text")
						fmt.Println("Xlsx Dir ", text)
						xl := XlsxStr{Text: text}
						xl.Count()
						return nil
					},
				},
			},
	}
	return &c
}