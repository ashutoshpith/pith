package download

import (
	"fmt"
	"os"
	"sync"

	"github.com/urfave/cli/v2"
)


func Url(flags []cli.Flag) *cli.Command {
	c := cli.Command { 
		Name: "url",
		Aliases: []string{"u"},
		Flags: flags,
		Before: func(c *cli.Context) error {
			fmt.Println("Process Started")
			return nil
		},
		After: func(c *cli.Context) error {
			fmt.Println("Here We done with Your Work")
			return nil
		},
		Action: func (c *cli.Context) error  {
		 fmt.Println("Url Downloader started")
		 text1 := c.String("text1")
		 text2 := c.String("text2")
		 arg1 := c.String("arg1")
		 arg2 := c.String("arg2")
		 filepath, _ := os.Getwd()
		 fmt.Println("cw ", filepath)
		//  setPath := filepath + "/dir/" + "logo.svg"
		//  fmt.Println("set ", setPath)
		 var wg sync.WaitGroup
		 for i := 0; i < 2; i++ {
			 wg.Add(1)
		     go UrlDownloadFile(Payload{
			   filepath: arg1,
			   url: text1,
			   wg: &wg,
			 })
		 }
		 UrlDownloadFile(Payload{
			filepath: arg2,
			url: text2,
		 })
		 wg.Wait()

		 fmt.Println("Downloaded Url: ", text1)
		 fmt.Println("Downloaded Url: ", text2)
		 return nil
		 },
	}
	return &c
}
