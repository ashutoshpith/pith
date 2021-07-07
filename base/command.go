package base

import (
	"fmt"

	"github.com/urfave/cli/v2"
)


func Test(flags []cli.Flag) *cli.Command {
	c := cli.Command {
		Name: "A",
		Usage: "B",
		Aliases: []string{"a"},
		Flags: flags,
		Action: func(c *cli.Context) error {
			value := c.String("text")
			fmt.Println("here there ", value)
			return nil
		},
	}
	
	return &c
}

func Name(flags []cli.Flag) *cli.Command {
	c := cli.Command {
		Name: "name",
		Usage: "Name Printitng",
		Flags: flags,
		Action: func(c *cli.Context) error {
			flag := c.String("text")
			fmt.Print("Here ", flag)
			return nil
		},
	}
	return &c
}

func Command(flags []cli.Flag) *cli.Command {
	tasks := []string{"command", "code"}

	c := cli.Command { 
		Name: "command",
		Aliases: []string{"c"},
		Flags: flags,
		Action: func (c *cli.Context) error  {
		 fmt.Println("Command")
		 return nil
		 },
		 BashComplete: func(c *cli.Context) {
			 if c.NArg() > 0 {
				 return 
			 }
			 for _, t := range tasks {
				 fmt.Println(t)
			 }
		 },
	}
	return &c
}
