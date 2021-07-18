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
	"github.com/adnsio/vmkit/pkg/driver"
	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Short: "Spin up Linux VMs with QEMU and Apple virtualization framework",
		Use:   "vmkit",
	}

	cmd.AddCommand(newCompletionCommand())
	cmd.AddCommand(newImagesCommand())
	cmd.AddCommand(newListCommand())
	cmd.AddCommand(newLogsCommand())
	cmd.AddCommand(newMacAddressCommand())
	cmd.AddCommand(newPullCommand())
	cmd.AddCommand(newRemoveCommand())
	cmd.AddCommand(newStartCommand())
	cmd.AddCommand(newStopCommand())

	cmd.PersistentFlags().String("driver", string(driver.DefaultDriver), "driver to use (qemu, avfvm)")
	cmd.PersistentFlags().String("avfvm", driver.AVFVMExecutableName, "avfvm executable name")
	cmd.PersistentFlags().String("qemu", driver.QEMUExecutableName, "qemu executable name")
	cmd.PersistentFlags().String("qemu-img", "qemu-img", "qemu-img executable name")

	return cmd
}
