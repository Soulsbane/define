package main

import (
	"fmt"
	"os"

	"github.com/Soulsbane/define/pkg/dictionary"
	"github.com/alexflint/go-arg"
	"github.com/jedib0t/go-pretty/v6/table"
)

func ListDefinitions(word string, listAll bool) {
	definitions, err := dictionary.GetDefinition(word)

	if err != nil {
		fmt.Println(err)
	} else {
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
}

func main() {
	var args ProgramArgs

	arg.MustParse(&args)

	if args.Word != "" {
		ListDefinitions(args.Word, args.ListAll)
	}
}
