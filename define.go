package main

import (
	"encoding/json"
	"fmt"

	"github.com/alexflint/go-arg"
	"github.com/levigross/grequests"
)

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

func getDefinition(wordToFind string) {
	resp, err := grequests.Get("https://api.dictionaryapi.dev/api/v2/entries/en/"+wordToFind, nil)

	if err != nil {
		fmt.Println("Failed to fetch word from dictionary API: ", err)
	}

	var dictionaryObject []DictionaryObject
	err = json.Unmarshal(resp.Bytes(), &dictionaryObject)

	if err != nil {
		fmt.Println("Failed to find a definition for: " + wordToFind)
	} else {
		fmt.Println(dictionaryObject[0].Meanings[0].Definitions[0].Definition)
	}

}

// https://api.dictionaryapi.dev/api/v2/entries/<language_code>/<word>
// https://api.dictionaryapi.dev/api/v2/entries/en/hello
// https://api.dictionaryapi.dev/api/v2/entries/ja/緑
func main() {
	var args struct {
		Word string `arg:"positional, required"`
	}

	arg.MustParse(&args)

	if args.Word != "" {
		getDefinition(args.Word)
	}
}
