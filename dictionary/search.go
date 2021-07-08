package dictionary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Delta456/box-cli-maker/v2"
)
type Dicti struct {
	url string
	Query string
	// data Data
}

type Data struct{
	Word string
	Meanings []Meanings

}

type Meanings struct {
   PartOfSpeech string
   Definitions []Definitions
}

type Definitions struct {
   Definition string
   Example string
}

func (api *Dicti) Search() (*[]Data, error)  {
	req := "https://api.dictionaryapi.dev/api/v2/entries/en_US/" 
	api.url = req
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	res := api.url + api.Query
	        url, _ := client.Get(res)
			body, err := ioutil.ReadAll(url.Body)
			if err != nil {
				return nil, err
			}
			defer url.Body.Close()
            var data []Data
			err = json.Unmarshal(body, &data)
			if err != nil {
				return nil, err
			}
			return &data, nil
}

func (api *Dicti) Iterate(data []Data) {
	config := box.Config{Px: 12, Py: 2, Type: "", TitlePos: "Inside"}
    boxNew := box.Box{TopRight: "*", TopLeft: "*", BottomRight: "*", BottomLeft: "*", Horizontal: "-", Vertical: "|", Config: config}
	for _, j := range data {
		fmt.Println("Search:", j.Word)
		fmt.Println()
		for _, md := range j.Meanings {
			fmt.Println("Part of Speech: ", md.PartOfSpeech)
			for _, def := range md.Definitions {
				defOut := "Definition: "+ def.Definition
				examOut := "Example: " +def.Example
			    boxNew.Println(defOut, examOut)
			}
		}
	}
}