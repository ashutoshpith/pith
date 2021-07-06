package download

import (
	"io"
	"net/http"
	"os"
	"sync"
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

	_, err = io.Copy(out, resp.Body)
	return err
}
