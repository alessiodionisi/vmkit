package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type logsOptions struct {
	address string
	name    string
}

func newLogsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "logs [vm]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			address, err := cmd.Flags().GetString("address")
			if err != nil {
				return err
			}

			opts := &logsOptions{
				address: address,
				name:    args[0],
			}

			if err := runLogs(opts); err != nil {
				fmt.Printf("error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runLogs(opts *logsOptions) error {
	client, err := NewRPCClient(opts.address)
	if err != nil {
		return err
	}

	logs, err := client.VirtualMachineLogs(opts.name)
	if err != nil {
		return err
	}

	if logs != "" {
		fmt.Println(logs)
	}

	return nil
}
