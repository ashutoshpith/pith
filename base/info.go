package base

import "github.com/urfave/cli/v2"

type InfoType struct {
	Name string
	Version string
	Usage string
	Help string
}

func Info() *InfoType {
	info := InfoType{
		Name: "Pith",
		Version: "1.0.2",
		Usage: "A Cli App",
		Help: "Not Yet Needed",
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