package dictionary

import (
	"time"

	"github.com/ashutoshpith/api"
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
			dict := Dictionary{
					api.Api{
					 Query: value,
					 Timeout: 10 * time.Second,
					 Lang: "en_US",
					},
			}
			
			data, _ := dict.DictSearch()
			dict.DictIterate(*data)
	
			return nil
		},
	}
	
	return &c
}