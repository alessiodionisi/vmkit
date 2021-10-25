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

	"github.com/spf13/cobra"
)

type pullOptions struct {
	name          string
	globalOptions *globalOptions
}

func newPullCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.ExactArgs(1),
		Short: "Pull an image",
		Use:   "pull [name]",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			opts := &pullOptions{
				name:          args[0],
				globalOptions: globalOptions,
			}

			if err := runPull(opts); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runPull(opts *pullOptions) error {
	eng, err := newEngine(opts.globalOptions)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	img := eng.FindImage(opts.name)
	if img == nil {
		return fmt.Errorf(`image "%s" not found`, opts.name)
	}

	return img.Pull()
}
