package main

import "github.com/spf13/cobra"

func newStartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start [vm]",
		Short: "Start a stopped virtual machine",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
