package download

import (
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/schollz/progressbar/v3"
)
type Payload struct {
	filepath string
	url string 
	wg *sync.WaitGroup
}
func UrlDownloadFile(payload Payload) error {
	if payload.wg != nil {
		defer payload.wg.Done()
	}
	resp, err := http.Get(payload.url)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	out , err := os.Create(payload.filepath)
	if err != nil {
		return err
	}

	defer out.Close()
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	

	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	return err
}
