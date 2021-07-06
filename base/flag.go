package base

import "github.com/urfave/cli/v2"

func TextFlag() *cli.StringFlag {
	f := cli.StringFlag {     
	  Name: "text",
	  Aliases: []string{"t"},
	  Usage: "Show Name",
	}
	return &f
	
}
 
func HostFlag() *cli.StringFlag {
	 f := cli.StringFlag {     
	   Name: "host",
	   Aliases: []string{"ho"},
	   Usage: "Host Flag",
	 }
	 return &f
	 
}