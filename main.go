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

    textFlag := base.TextFlag()
	hostFlag := base.HostFlag()
	argFlag := base.ArgFlag()

	app.Flags = []cli.Flag {
		textFlag, hostFlag, argFlag,
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