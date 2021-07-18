// Spin up Linux VMs with QEMU and Apple virtualization framework
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

	"github.com/adnsio/vmkit/pkg/engine"
	"github.com/spf13/cobra"
)

type pullOptions struct {
	name string
}

func newPullCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.ExactArgs(1),
		Short: "Pull an image",
		Use:   "pull [image]",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := &pullOptions{
				name: args[0],
			}

			if err := runPull(opts); err != nil {
				fmt.Printf("error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runPull(opts *pullOptions) error {
	eng, err := engine.New(&engine.NewOptions{
		Writer: os.Stderr,
	})
	if err != nil {
		return err
	}

	img := eng.FindImage(opts.name)
	if img == nil {
		return fmt.Errorf(`image "%s" not found`, opts.name)
	}

	return img.Pull()
}
