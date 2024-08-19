package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Soulsbane/define/pkg/dictionary"
	"github.com/alexflint/go-arg"
	"github.com/jedib0t/go-pretty/v6/table"
)

func getOutputTable() table.Writer {

	outputTable := table.NewWriter()

	outputTable.SetOutputMirror(os.Stdout)
	outputTable.AppendHeader(table.Row{"Definition"})
	outputTable.SetStyle(table.StyleRounded)
	outputTable.Style().Options.SeparateRows = true

	return outputTable
}

func ListDefinitions(definitions *[]dictionary.DefinitionsObject, listAll bool) {
	outputTable := getOutputTable()

	if listAll {
		for _, definitionObject := range *definitions {
			outputTable.AppendRow(table.Row{definitionObject.Definition})
		}

	} else {
		outputTable.AppendRow(table.Row{(*definitions)[0].Definition})
	}

	outputTable.Render()
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
