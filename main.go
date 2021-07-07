package main

import (
	"fmt"
	"log"
	"os"

	"sort"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/ashutoshpith/base"
	"github.com/ashutoshpith/download"
	"github.com/ashutoshpith/xlsx"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func main() {
	config := box.Config{Px: 12, Py: 2, Type: "", TitlePos: "Inside"}
    boxNew := box.Box{TopRight: "*", TopLeft: "*", BottomRight: "*", BottomLeft: "*", Horizontal: "-", Vertical: "|", Config: config}
    boxNew.Println("Pith", "A CLi")

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
 
		After: func(c *cli.Context) error {
			fmt.Fprintf(c.App.Writer, "\n\tPhew!\n")
			return nil
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