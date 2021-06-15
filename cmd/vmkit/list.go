// Virtual Machine manager that supports QEMU and Apple virtualization framework on macOS
// Copyright (C) 2021 VMKit Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
