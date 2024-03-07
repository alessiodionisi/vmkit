package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newImageListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List images",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			if err := imageList(globalOptions); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func imageList(opts *globalOptions) error {
	return nil
}
