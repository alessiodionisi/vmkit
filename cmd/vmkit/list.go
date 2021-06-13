package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type listOptions struct {
	address string
}

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List virtual machines",
		Aliases: []string{
			"ls",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			address, err := cmd.Flags().GetString("address")
			if err != nil {
				return err
			}

			opts := &listOptions{
				address: address,
			}

			if err := runList(opts); err != nil {
				fmt.Printf("error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runList(opts *listOptions) error {
	client, err := NewRPCClient(opts.address)
	if err != nil {
		return err
	}

	virtualMachines, err := client.ListVirtualMachines()
	if err != nil {
		return err
	}

	tableRows := make([][]string, len(virtualMachines))
	for i, vm := range virtualMachines {
		tableRows[i] = []string{
			vm.Config.Metadata.Name,
			strings.Title(string(vm.Status)),
		}
	}

	writeTable(&writeTableOptions{
		Writer: os.Stdout,
		Header: []string{"Name", "Status"},
		Rows:   tableRows,
	})

	return nil
}
