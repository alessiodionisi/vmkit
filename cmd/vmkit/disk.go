package main

import (
	"github.com/spf13/cobra"
)

func newDiskCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disk",
		Short: "Manage disks",
	}

	cmd.AddCommand(newDiskCreateCommand())
	cmd.AddCommand(newDiskListCommand())

	return cmd
}
