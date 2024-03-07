package cmd

import (
	"github.com/spf13/cobra"
)

func newInstanceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "instance",
		Short: "Manage instances",
	}

	cmd.AddCommand(newInstanceCreateCommand())
	cmd.AddCommand(newInstanceDeleteCommand())
	cmd.AddCommand(newInstanceListCommand())
	cmd.AddCommand(newInstanceStartCommand())
	cmd.AddCommand(newInstanceStopCommand())

	return cmd
}
