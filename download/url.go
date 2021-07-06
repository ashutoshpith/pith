package download

import (
	"io"
	"net/http"
	"os"
)

func UrlDownloadFile(filepath string , url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	out , err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
