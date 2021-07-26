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

	"github.com/spf13/cobra"
)

type restartOptions struct {
	*globalOptions
	name string
}

func newRestartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.ExactArgs(1),
		Short: "Restart a running virtual machine",
		Use:   "restart [name]",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			opts := &restartOptions{
				name:          args[0],
				globalOptions: globalOptions,
			}

			if err := runRestart(opts); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runRestart(opts *restartOptions) error {
	eng, err := newEngine(opts.globalOptions)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	vm := eng.FindVirtualMachine(opts.name)
	if vm == nil {
		return fmt.Errorf(`virtual machine "%s" not found`, opts.name)
	}

	if err := vm.Stop(); err != nil {
		return err
	}

	return vm.Start()
}
