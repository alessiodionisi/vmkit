package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type writeTableOptions struct {
	Header []string
	Rows   [][]string
}

func writeTable(opts *writeTableOptions) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoFormatHeaders(true)
	table.SetAutoWrapText(false)
	table.SetBorder(false)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetNoWhiteSpace(true)
	table.SetRowSeparator("")
	table.SetTablePadding("\t")

	table.AppendBulk(opts.Rows)
	table.SetHeader(opts.Header)

	table.Render()
}
