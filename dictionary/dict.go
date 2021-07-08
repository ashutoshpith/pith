package dictionary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ashutoshpith/api"
	"github.com/ashutoshpith/box"
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
	box := box.Box()
		for _, j := range data {
		fmt.Println("Search:", j.Word)
		fmt.Println()
		for _, md := range j.Meanings {
			fmt.Println("Part of Speech: ", md.PartOfSpeech)
			for _, def := range md.Definitions {
				defOut := "Definition: "+ def.Definition
				examOut := "Example: " +def.Example
			    box.Println(defOut, examOut)
			}
		}
	}
}