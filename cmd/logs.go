package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type logsOptions struct {
	name          string
	globalOptions *globalOptions
}

func newLogsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args: cobra.ExactArgs(1),
		Use:  "logs [name]",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			opts := &logsOptions{
				name:          args[0],
				globalOptions: globalOptions,
			}

			if err := runLogs(opts); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runLogs(opts *logsOptions) error {
	return nil
}
