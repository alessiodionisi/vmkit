package main

import (
	"github.com/spf13/cobra"
)

type startCmdOptions struct {
	commonCmdOptions
}

func newStartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start [name]",
		Short: "Start a virtual machine",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := runStartCmd(startCmdOptions{}); err != nil {
				handleCmdError(err)
				return nil
			}

			return nil
		},
	}

	return cmd
}

func runStartCmd(_ startCmdOptions) error {
	clientConn, _, err := newClient()
	if err != nil {
		return err
	}
	defer clientConn.Close()

	// ctx := context.Background()

	// _, err = client.Start(ctx, &proto.StartRequest{})
	// if err != nil {
	// 	return err
	// }

	return nil
}
