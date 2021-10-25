package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Short: "List virtual machines",
		Use:   "list",
		Aliases: []string{
			"ls",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			if err := runList(globalOptions); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runList(opts *globalOptions) error {
	eng, err := newEngine(opts)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	vms := eng.ListVirtualMachines()

	tableRows := make([][]string, 0, len(vms))
	for _, vm := range vms {
		status, err := vm.Status()
		if err != nil {
			return err
		}

		tableRows = append(tableRows, []string{
			vm.Name,
			strings.Title(string(status)),
		})
	}

	writeTable(&writeTableOptions{
		writer: os.Stdout,
		header: []string{
			"Name",
			"Status",
		},
		rows: tableRows,
	})

	return nil
}
