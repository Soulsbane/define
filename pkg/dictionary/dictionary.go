package dictionary

import (
	"fmt"
	"github.com/imroc/req/v3"
)

// https://api.dictionaryapi.dev/api/v2/entries/<language_code>/<word>
// https://api.dictionaryapi.dev/api/v2/entries/en/hello
// https://api.dictionaryapi.dev/api/v2/entries/ja/緑

const API_URL = "https://api.dictionaryapi.dev/api/v2/entries/en/"

type DefinitionsObject struct {
	Definition string   `json:"definition"`
	Example    string   `json:"example"`
	Synonyms   []string `json:"synonyms"`
}

// DictionaryObject Structure containing the values fetched from dictionaryapi.dev
type DictionaryObject struct {
	Meanings []struct {
		Definitions  []DefinitionsObject `json:"definitions"`
		PartOfSpeech string              `json:"partOfSpeech"`
	} `json:"meanings"`
	Phonetics []struct {
		Audio string `json:"audio"`
		Text  string `json:"text"`
	} `json:"phonetics"`
	Word string `json:"word"`
}

func GetDefinition(wordToFind string, all bool) (*[]DefinitionsObject, error) {
	var dictionaryObject []DictionaryObject
	client := req.C()

	_, err := client.R().SetSuccessResult(&dictionaryObject).Get(API_URL + wordToFind)

	if err != nil {
		return nil, fmt.Errorf("failed to find a definition for: %s", wordToFind)
	} else {
		return &dictionaryObject[0].Meanings[0].Definitions, nil
	}
}
