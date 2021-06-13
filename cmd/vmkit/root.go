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
