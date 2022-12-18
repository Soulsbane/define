package main

import (
	"fmt"

	"github.com/Soulsbane/define/pkg/dictionary"
	"github.com/alexflint/go-arg"
)

func main() {
	var args ProgramArgs

	arg.MustParse(&args)

	if args.Word != "" {
		definition, err := dictionary.GetDefinition(args.Word)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(definition.Definition)
		}
	}
}
