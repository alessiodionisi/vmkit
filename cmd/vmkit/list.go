package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List virtual machines",
		Aliases: []string{
			"ls",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			userHomeDir, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			vmsDir := fmt.Sprintf("%s/.vmkit/vms", userHomeDir)

			_, err = os.Stat(vmsDir)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					return nil
				}

				return err
			}

			return nil
		},
	}

	return cmd
}
