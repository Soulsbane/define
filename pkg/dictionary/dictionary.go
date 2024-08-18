package dictionary

import (
	"fmt"
	"github.com/imroc/req/v3"
)

// https://api.dictionaryapi.dev/api/v2/entries/<language_code>/<word>
// https://api.dictionaryapi.dev/api/v2/entries/en/hello
// https://api.dictionaryapi.dev/api/v2/entries/ja/緑

const API_URL = "https://api.dictionaryapi.dev/api/v2/entries/en/"
const errorNoDefinition = "Failed to find a definition for %s"

type DefinitionsObject struct {
	Definition string   `json:"definition"`
	Example    string   `json:"example"`
	Synonyms   []string `json:"synonyms"`
}

// DefinitionResult Structure containing the values fetched from dictionaryapi.dev
type DefinitionResult struct {
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

func GetDefinition(wordToFind string) (*[]DefinitionsObject, error) {
	var dictionaryObject []DefinitionResult
	client := req.C()

	_, err := client.R().SetSuccessResult(&dictionaryObject).Get(API_URL + wordToFind)

	if err != nil {
		return nil, fmt.Errorf(errorNoDefinition, wordToFind)
	} else {
		if len(dictionaryObject) == 0 {
			return nil, fmt.Errorf(errorNoDefinition, wordToFind)
		} else {
			return &dictionaryObject[0].Meanings[0].Definitions, nil
		}
	}
}
