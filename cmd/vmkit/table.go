package main

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

type writeTableOptions struct {
	Writer io.Writer
	Header []string
	Rows   [][]string
}

func writeTable(opts *writeTableOptions) {
	table := tablewriter.NewWriter(opts.Writer)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

	table.SetHeader(opts.Header)
	table.AppendBulk(opts.Rows)

	table.Render()
}
