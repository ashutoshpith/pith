package sysInfo

import (
	"github.com/urfave/cli/v2"
)


func Sys(flags []cli.Flag) *cli.Command {
	c := cli.Command { 
		Name: "sys",
		Usage: "System Info",
		Aliases: []string{"s"},
		Flags: flags,
		Action: func (c *cli.Context) error  {
		//  SysInfo()
		 return nil
		 },
	}
	return &c
}
