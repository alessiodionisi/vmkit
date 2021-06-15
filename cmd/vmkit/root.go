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
	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vmkit",
		Short: "VMKit CLI controls the running VMKit Daemon",
	}

	cmd.AddCommand(newApplyCommand())
	cmd.AddCommand(newCompletionCommand())
	cmd.AddCommand(newListCommand())
	cmd.AddCommand(newLogsCommand())
	cmd.AddCommand(newMacAddressCommand())
	cmd.AddCommand(newRemoveCommand())
	cmd.AddCommand(newStartCommand())
	cmd.AddCommand(newStopCommand())

	cmd.PersistentFlags().StringP("address", "a", "unix:///var/run/vmkitd.sock", "(unix, tcp)")

	return cmd
}
