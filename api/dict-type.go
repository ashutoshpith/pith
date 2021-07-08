package api

type DictData struct{
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