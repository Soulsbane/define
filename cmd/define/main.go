package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Soulsbane/define/pkg/dictionary"
	"github.com/alexflint/go-arg"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/tiagomelo/go-clipboard/clipboard"
)

func getOutputTable(maxWidth int) table.Writer {

	outputTable := table.NewWriter()

	outputTable.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignLeft, WidthMax: maxWidth},
	})

	outputTable.SetOutputMirror(os.Stdout)
	outputTable.AppendHeader(table.Row{"Definition"})
	outputTable.SetStyle(table.StyleRounded)
	outputTable.Style().Options.SeparateRows = true

	return outputTable
}

func handleCopyToClipboard(definition string) {
	c := clipboard.New()

	if err := c.CopyText(definition); err != nil {
		fmt.Println(err)
	}
}

func listDefinitions(definitions *[]dictionary.DefinitionsObject, listAll bool, copyToClipboard bool, maxWidth int) {
	outputTable := getOutputTable(maxWidth)

	if listAll {
		for _, definitionObject := range *definitions {
			outputTable.AppendRow(table.Row{definitionObject.Definition})
		}

	} else {
		outputTable.AppendRow(table.Row{(*definitions)[0].Definition})

		if copyToClipboard {
			handleCopyToClipboard((*definitions)[0].Definition)
		}
	}

	outputTable.Render()
}

func main() {
	var args ProgramArgs

	arg.MustParse(&args)

	if args.Word != "" {
		definitions, err := dictionary.GetDefinition(args.Word)

		if errors.Is(err, dictionary.ErrNoDefinition) {
			fmt.Println("No definition found for", args.Word)
		} else {
			listDefinitions(definitions, args.ListAll, args.Copy, args.MaxWidth)
		}
	}
}
