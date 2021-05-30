package main

import "github.com/spf13/cobra"

func newExecCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exec [vm] [command]",
		Short: "Run a command in a virtual machine",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
