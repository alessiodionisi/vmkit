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

	"github.com/adnsio/vmkit/pkg/engine"
	"github.com/spf13/cobra"
)

type createOptions struct {
	*globalOptions
	cpu    int
	image  string
	memory int
	name   string
}

func newCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Example: `  Create an Ubuntu Hirsute virtual machine:
    vmkit create vm1 --image ubuntu:hirsute

  Create a virtual machine with 4 cpus and 4096 megabytes of ram:
    vmkit create vm1 --image ubuntu:hirsute --cpu 4 --memory 4096`,
		Args:  cobra.ExactArgs(1),
		Short: "Create and start a new virtual machine",
		Use:   "create [name]",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			image, err := cmd.Flags().GetString("image")
			if err != nil {
				return err
			}

			cpu, err := cmd.Flags().GetInt("cpu")
			if err != nil {
				return err
			}

			memory, err := cmd.Flags().GetInt("memory")
			if err != nil {
				return err
			}

			opts := &createOptions{
				cpu:           cpu,
				globalOptions: globalOptions,
				image:         image,
				memory:        memory,
				name:          args[0],
			}

			if err := runCreate(opts); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().StringP("image", "i", "", "image to use")
	cmd.Flags().Int("cpu", 1, "number of cpu(s)")
	cmd.Flags().Int("memory", 512, "ram in megabytes")

	cmd.MarkFlagRequired("image")

	return cmd
}

func runCreate(opts *createOptions) error {
	eng, err := newEngine(opts.globalOptions)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	vm, err := eng.CreateVirtualMachine(&engine.CreateVirtualMachineOptions{
		CPU:    opts.cpu,
		Image:  opts.image,
		Memory: opts.memory,
		Name:   opts.name,
	})
	if err != nil {
		return err
	}

	return vm.Start()
}
