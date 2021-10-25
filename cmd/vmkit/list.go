// Spin up Linux VMs with QEMU
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
