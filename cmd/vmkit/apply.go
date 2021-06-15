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

type applyOptions struct {
	file string
}

func newApplyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "Apply a configuration from file",
		Example: `  # Apply the configuration in ubuntu.yaml
  vmkit apply -f ./ubuntu.yaml`,
		RunE: func(cmd *cobra.Command, args []string) error {
			file, err := cmd.Flags().GetString("file")
			if err != nil {
				return err
			}

			opts := &applyOptions{
				file: file,
			}

			if err := runApply(opts); err != nil {
				fmt.Printf("error: %s", err)
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().StringP("file", "f", "", "configuration file path")

	return cmd
}

func runApply(opts *applyOptions) error {
	return nil
}
