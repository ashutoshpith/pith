package dictionary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ashutoshpith/api"

	"github.com/Delta456/box-cli-maker/v2"
)

type Dictionary struct{
	api api.Api
}


func (dict *Dictionary) DictSearch() (*[]DictData, error)  {
	req := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/%s/%s", dict.api.Lang, dict.api.Query) 

	dict.api.Url = req
	client := http.Client{
		Timeout: dict.api.Timeout,
	}
	        url, _ := client.Get(dict.api.Url)
			body, err := ioutil.ReadAll(url.Body)
			if err != nil {
				return nil, err
			}
			defer url.Body.Close()
            var data []DictData
			err = json.Unmarshal(body, &data)
			if err != nil {
				return nil, err
			}
			return &data, nil
}

func (dict *Dictionary) DictIterate(data []DictData) {
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