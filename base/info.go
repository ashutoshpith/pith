package base

import "github.com/urfave/cli/v2"

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
		Version: "1.0.2",
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