package main

import (
	"fmt"

	"github.com/Soulsbane/define/pkg/dictionary"
	"github.com/alexflint/go-arg"
)

func ListDefinitions(word string, listAll bool) {
	definitions, err := dictionary.GetDefinition(word, false)

	if err != nil {
		fmt.Println(err)
	} else {
		if listAll {
			for _, definitionObject := range *definitions {
				fmt.Println(definitionObject.Definition)
			}
		} else {
			fmt.Println((*definitions)[0].Definition)
		}
	}
}

func main() {
	var args ProgramArgs

	arg.MustParse(&args)

	if args.Word != "" {
		ListDefinitions("cat", args.ListAll)
	}
}
