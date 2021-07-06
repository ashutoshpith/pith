package main

import (
	"log"
	"os"

	"sort"

	"github.com/ashutoshpith/base"
	"github.com/ashutoshpith/download"
	"github.com/ashutoshpith/xlsx"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Pith"
	app.Usage = "Magic To Begin"
	app.Version = "v1.0.0"

    text1Flag := base.Text1Flag()
    text2Flag := base.Text2Flag()
	hostFlag := base.HostFlag()
	arg1Flag := base.Arg1Flag()
	arg2Flag := base.Arg2Flag()

	app.Flags = []cli.Flag {
		text1Flag, text2Flag, hostFlag, arg1Flag, arg2Flag,
	}

	testCommand := base.Test(app.Flags)
	nameCommand := base.Name(app.Flags)
    commandCommand := base.Command(app.Flags)
	xlsxCommand := xlsx.Xlsx(app.Flags)
	urlDownloadCommand := download.Url(app.Flags)

	app.Commands = []*cli.Command {
		testCommand,
		nameCommand,
		commandCommand,
		xlsxCommand,
		urlDownloadCommand,

	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
  
	err := app.Run(os.Args)
	if err != nil {
	  log.Fatal(err)
	}
  }