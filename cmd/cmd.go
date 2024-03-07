package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/alessiodionisi/vmkit/engine"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type globalOptions struct {
	dataDir string
	debug   bool
}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vmkit",
		Short: "VMKit manages virtual machines instances, volumes and images",
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	cmd.PersistentFlags().String("data-dir", path.Join(homeDir, ".vmkit"), "Directory to store data")
	cmd.PersistentFlags().Bool("debug", false, "Enable debug mode")

	cmd.AddCommand(newInstanceCommand())
	cmd.AddCommand(newVolumeCommand())
	cmd.AddCommand(newImageCommand())

	return cmd
}

func newGlobalOptions(cmd *cobra.Command) (*globalOptions, error) {
	dataDir, err := cmd.Flags().GetString("data-dir")
	if err != nil {
		return nil, err
	}

	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		return nil, err
	}

	return &globalOptions{
		dataDir: dataDir,
		debug:   debug,
	}, nil
}

func newEngine(opts *globalOptions) (*engine.Engine, error) {
	eng, err := engine.New(engine.NewOptions{
		Debug:     opts.debug,
		DataDir:   opts.dataDir,
		LogWriter: os.Stdout,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating engine: %w", err)
	}

	return eng, nil
}

type writeTableOptions struct {
	header []string
	rows   [][]string
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

	table.AppendBulk(opts.rows)
	table.SetHeader(opts.header)

	table.Render()
}
