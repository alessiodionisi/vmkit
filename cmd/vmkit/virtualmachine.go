package main

import (
	"github.com/spf13/cobra"
)

func newVirtualMachineCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "virtualmachine",
		Short:   "Manage virtual machines",
		Aliases: []string{"vm"},
	}

	cmd.AddCommand(newVirtualMachineCreateCommand())
	cmd.AddCommand(newVirtualMachineDeleteCommand())
	cmd.AddCommand(newVirtualMachineListCommand())
	cmd.AddCommand(newVirtualMachineStartCommand())
	cmd.AddCommand(newVirtualMachineStopCommand())

	return cmd
}
