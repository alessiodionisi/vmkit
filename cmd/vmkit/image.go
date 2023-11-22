package main

import (
	"github.com/spf13/cobra"
)

func newImageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "image",
		Short: "Manage images",
	}

	cmd.AddCommand(newImageListCommand())

	return cmd
}
