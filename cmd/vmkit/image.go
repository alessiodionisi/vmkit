package main

import (
	"github.com/spf13/cobra"
)

func newImageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "image",
		Short:   "Commands to manage images",
		Aliases: []string{"img", "images"},
	}

	cmd.AddCommand(newDeleteImageCmd())
	cmd.AddCommand(newListImagesCmd())
	cmd.AddCommand(newPullImageCmd())

	return cmd
}
