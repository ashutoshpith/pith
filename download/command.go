package download

import (
	"fmt"
	"sync"

	"github.com/ashutoshpith/box"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)


func Url(flags []cli.Flag) *cli.Command {
	c := cli.Command { 
		Name: "url",
		Aliases: []string{"u"},
		Flags: flags,
		Action: func (c *cli.Context) error  {
		 text1 := c.String("text1")
		 text2 := c.String("text2")
		 arg1 := c.String("arg1")
		 arg2 := c.String("arg2")
		 
		UrlDownloadFile(Payload{
			filepath: arg1,
			url: text1,
		 })
		 if text2 != "" {
		 var wg sync.WaitGroup
		 for i := 0; i < 2; i++ {
			 wg.Add(1)
		     go UrlDownloadFile(Payload{
			   filepath: arg2,
			   url: text2,
			   wg: &wg,
			 })
		 }
		 wg.Wait()
		}
		box := box.Box()
		fmt.Println()
		color.Success.Block("Downloaded Url: ")
		box.Println(text1, text1)
		 return nil
		 },
	}
	return &c
}
