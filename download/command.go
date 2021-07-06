package download

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)


func Url(flags []cli.Flag) *cli.Command {
	c := cli.Command { 
		Name: "url",
		Aliases: []string{"u"},
		Flags: flags,
		Action: func (c *cli.Context) error  {
		 fmt.Println("Url Downloader started")
		 text := c.String("text")
		 filepath, _ := os.Getwd()
		 fmt.Println("cw ", filepath)
		 setPath := filepath + "/dir/" + "logo.svg"
		 fmt.Println("set ", setPath)
		 err := UrlDownloadFile("lo.svg", text)
		 if err != nil {
			 panic(err)
		 }
		 fmt.Println("Downloaded: ", text)
		 return nil
		 },
	}
	return &c
}
