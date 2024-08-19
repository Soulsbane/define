package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Soulsbane/define/pkg/dictionary"
	"github.com/alexflint/go-arg"
	"github.com/jedib0t/go-pretty/v6/table"
)

func ListDefinitions(definitions *[]dictionary.DefinitionsObject, listAll bool) {
	if listAll {
		outputTable := table.NewWriter()

		outputTable.SetOutputMirror(os.Stdout)
		outputTable.AppendHeader(table.Row{"Definition"})

		for _, definitionObject := range *definitions {
			outputTable.AppendRow(table.Row{definitionObject.Definition})
		}

		outputTable.SetStyle(table.StyleRounded)
		outputTable.Style().Options.SeparateRows = true
		outputTable.Render()
	} else {
		fmt.Println((*definitions)[0].Definition)
	}
}

func main() {
	var args ProgramArgs

	arg.MustParse(&args)

	if args.Word != "" {
		definitions, err := dictionary.GetDefinition(args.Word)

		if errors.As(err, &dictionary.ErrorNoDefinition) {
			fmt.Println("No definition found")
		} else {
			ListDefinitions(definitions, args.ListAll)
		}
	}
}
