package dictionary

import (
	"encoding/json"
	"fmt"

	"github.com/levigross/grequests"
)

// https://api.dictionaryapi.dev/api/v2/entries/<language_code>/<word>
// https://api.dictionaryapi.dev/api/v2/entries/en/hello
// https://api.dictionaryapi.dev/api/v2/entries/ja/緑

// DictionaryObject Structure containing the values fetched from dictionaryapi.dev
type DictionaryObject struct {
	Meanings []struct {
		Definitions []struct {
			Definition string   `json:"definition"`
			Example    string   `json:"example"`
			Synonyms   []string `json:"synonyms"`
		} `json:"definitions"`
		PartOfSpeech string `json:"partOfSpeech"`
	} `json:"meanings"`
	Phonetics []struct {
		Audio string `json:"audio"`
		Text  string `json:"text"`
	} `json:"phonetics"`
	Word string `json:"word"`
}

func GetDefinition(wordToFind string) (string, error) {
	var dictionaryObject []DictionaryObject
	resp, err := grequests.Get("https://api.dictionaryapi.dev/api/v2/entries/en/"+wordToFind, nil)

	if err != nil {
		return "<Unknown>", fmt.Errorf("failed to fetch word from dictionary API: %s", err)
	}

	err = json.Unmarshal(resp.Bytes(), &dictionaryObject)

	if err != nil {
		return "<Unknown>", fmt.Errorf("failed to find a definition for: %s", wordToFind)
	} else {
		return dictionaryObject[0].Meanings[0].Definitions[0].Definition, nil
	}
}
