package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type applyOptions struct {
	file string
}

func newApplyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "Apply a configuration from file",
		Example: `  # Apply the configuration in ubuntu.yaml
  vmkit apply -f ./ubuntu.yaml`,
		RunE: func(cmd *cobra.Command, args []string) error {
			file, err := cmd.Flags().GetString("file")
			if err != nil {
				return err
			}

			opts := &applyOptions{
				file: file,
			}

			if err := runApply(opts); err != nil {
				fmt.Printf("error: %s", err)
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().StringP("file", "f", "", "configuration file path")

	return cmd
}

func runApply(opts *applyOptions) error {
	return nil
}
