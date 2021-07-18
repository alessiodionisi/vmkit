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
	"errors"
	"fmt"
	"os"

	"github.com/adnsio/vmkit/pkg/driver"
	"github.com/adnsio/vmkit/pkg/engine"
	"github.com/spf13/cobra"
)

type startOptions struct {
	driver string
	name   string
	qemu   string
}

func newStartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.ExactArgs(1),
		Short: "Start a stopped virtual machine",
		Use:   "start [vm]",
		RunE: func(cmd *cobra.Command, args []string) error {
			driver, err := cmd.Flags().GetString("driver")
			if err != nil {
				return err
			}

			qemu, err := cmd.Flags().GetString("qemu")
			if err != nil {
				return err
			}

			opts := &startOptions{
				driver: driver,
				name:   args[0],
				qemu:   qemu,
			}

			if err := runStart(opts); err != nil {
				fmt.Printf("error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runStart(opts *startOptions) error {
	var drv driver.Driver
	var err error

	switch driver.DriverType(opts.driver) {
	case driver.DriverTypeQEMU:
		drv, err = driver.NewQEMU(opts.qemu)
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid driver")
	}

	eng, err := engine.New(&engine.NewOptions{
		Writer: os.Stderr,
		Driver: drv,
	})
	if err != nil {
		return err
	}

	vm := eng.FindVirtualMachine(opts.name)
	if vm == nil {
		return fmt.Errorf(`virtual machine "%s" not found`, opts.name)
	}

	return vm.Start()
}
