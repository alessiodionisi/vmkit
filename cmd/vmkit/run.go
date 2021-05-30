package main

import "github.com/spf13/cobra"

func newRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Boot and run a virtual machine",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.Flags().Uint64P("cpu", "c", 1, "sets the number of CPUs")
	cmd.Flags().StringP("memory", "m", "1Gi", "sets the amount of physical memory")

	cmd.MarkFlagRequired("cpus")
	cmd.MarkFlagRequired("memory")

	return cmd
}
