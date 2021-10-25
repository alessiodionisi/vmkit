package cmd

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

type writeTableOptions struct {
	header []string
	rows   [][]string
	writer io.Writer
}

func writeTable(opts *writeTableOptions) {
	table := tablewriter.NewWriter(opts.writer)
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

	table.AppendBulk(opts.rows)
	table.SetHeader(opts.header)

	table.Render()
}
