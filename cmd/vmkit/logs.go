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

	"github.com/spf13/cobra"
)

type logsOptions struct {
	address string
	name    string
}

func newLogsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "logs [vm]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			address, err := cmd.Flags().GetString("address")
			if err != nil {
				return err
			}

			opts := &logsOptions{
				address: address,
				name:    args[0],
			}

			if err := runLogs(opts); err != nil {
				fmt.Printf("error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runLogs(opts *logsOptions) error {
	client, err := NewRPCClient(opts.address)
	if err != nil {
		return err
	}

	logs, err := client.VirtualMachineLogs(opts.name)
	if err != nil {
		return err
	}

	if logs != "" {
		fmt.Println(logs)
	}

	return nil
}
