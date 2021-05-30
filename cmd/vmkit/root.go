package main

import "github.com/spf13/cobra"

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "vmkit",
	}

	cmd.AddCommand(newRunCommand())
	cmd.AddCommand(newExecCommand())
	cmd.AddCommand(newStartCommand())
	cmd.AddCommand(newStopCommand())
	cmd.AddCommand(newListCommand())
	cmd.AddCommand(newRemoveCommand())

	return cmd
}
