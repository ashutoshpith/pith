package base

import (
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

type InfoType struct {
	Name string
	Version string
	Usage string
	Help string
	Copyright string
	ArgsUsage string
}

func Info() *InfoType {
	info := InfoType{
		Name: "Pith",
		Version: "1.0.3",
		Usage: "A Cli App",
		Help: "Not Yet Needed",
		Copyright: "(c) 2021 MIT",
		ArgsUsage: "[args and such]",
	}
	return &info
}

func Author() *cli.Author{
	author := cli.Author{
		Name: "Ashutosh Singh",
		Email: "ashutoshpith@gmail.com",
	} 
	return &author
}

func AuthorInfo(flags []cli.Flag) *cli.Command {
	c := cli.Command {
		Name: "author",
		Usage: "Author Info",
		Aliases: []string{"a"},
		Flags: flags,
		Action: func(c *cli.Context) error {
			pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
				{"Name", "Email", "Who AM I", "Job Profile"},
				{Author().Name, Author().Email, "A Magician", "Full Stack Developer"},
			}).Render()
			return nil
		},
	}
	
	return &c
}