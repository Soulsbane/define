package dictionary

import (
	"encoding/json"
	"errors"
	"net/http"
)

// https://api.dictionaryapi.dev/api/v2/entries/<language_code>/<word>
// https://api.dictionaryapi.dev/api/v2/entries/en/hello
// https://api.dictionaryapi.dev/api/v2/entries/ja/緑

const dictionaryURL = "https://api.dictionaryapi.dev/api/v2/entries/en/"

var ErrNoDefinition = errors.New("failed to find a definition for that word")
var ErrDownloadFailed = errors.New("failed to download definitions file")

// DefinitionsObject Structure containing the definition, example and synonyms

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

func GetDefinition(wordToFind string) ([]DefinitionsObject, error) {
	var result []DefinitionResult

	resp, err := http.Get(dictionaryURL + wordToFind)

	if err != nil {
		return nil, ErrDownloadFailed
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return nil, ErrNoDefinition
	}

	if len(result) == 0 {
		return nil, ErrNoDefinition
	} else {
		return result[0].Meanings[0].Definitions, nil
	}
}
