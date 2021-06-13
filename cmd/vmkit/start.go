package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type startOptions struct {
	address string
	name    string
}

func newStartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start [vm]",
		Short: "Start a stopped virtual machine",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			address, err := cmd.Flags().GetString("address")
			if err != nil {
				return err
			}

			opts := &startOptions{
				address: address,
				name:    args[0],
			}

			if err := runStart(opts); err != nil {
				fmt.Printf("error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runStart(opts *startOptions) error {
	client, err := NewRPCClient(opts.address)
	if err != nil {
		return err
	}

	if err := client.StartVirtualMachine(opts.name); err != nil {
		return err
	}

	return nil
}