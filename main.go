package main

import (
	"fmt"
	"log"
	"os"

	"sort"

	"github.com/ashutoshpith/base"
	"github.com/ashutoshpith/dictionary"
	"github.com/ashutoshpith/download"
	"github.com/ashutoshpith/sysInfo"
	"github.com/ashutoshpith/xlsx"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func main() {

	app := &cli.App{
		Name: base.Info().Name,
		Usage: base.Info().Usage,
		Version: base.Info().Version,
		EnableBashCompletion: true,
		Authors: []*cli.Author{
			base.Author(),
		},
		Copyright: base.Info().Copyright,
		ArgsUsage: base.Info().ArgsUsage,
		CommandNotFound: func(c *cli.Context, s string) {
			fmt.Println("There is No Command Like This")
		},		
		  
	}
    text1Flag := base.Text1Flag()
    text2Flag := base.Text2Flag()
	hostFlag := base.HostFlag()
	arg1Flag := base.Arg1Flag()
	arg2Flag := base.Arg2Flag()
	loadFloag := base.LoadFlag()

	app.Flags = []cli.Flag {
		altsrc.NewIntFlag(&cli.IntFlag{
			Name: "test",
			Aliases: []string{"ts"},
		}),
		
		text1Flag, text2Flag, hostFlag, arg1Flag, arg2Flag, loadFloag,
		
	}
	app.Before = altsrc.InitInputSourceWithContext(app.Flags, altsrc.NewYamlSourceFromFlagFunc("yaml"))

	nameCommand := base.Name(app.Flags)
    commandCommand := base.Command(app.Flags)
	xlsxCommand := xlsx.Xlsx(app.Flags)
	urlDownloadCommand := download.Url(app.Flags)
	dictCommand := dictionary.Dict(app.Flags)
	authorCommand := base.AuthorInfo(app.Flags)
	sysCommand := sysInfo.Sys(app.Flags)

	app.Commands = []*cli.Command {
		nameCommand,
		commandCommand,
		xlsxCommand,
		urlDownloadCommand,
		dictCommand,
		authorCommand,
		sysCommand,

	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))


	err := app.Run(os.Args)
	if err != nil {
	  log.Fatal(err)
	}
  }