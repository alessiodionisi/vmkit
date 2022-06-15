package main

import (
	"github.com/spf13/cobra"
)

type restartCmdOptions struct {
	commonCmdOptions
}

func newRestartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restart [name]",
		Short: "Restart a virtual machine",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := runRestartCmd(restartCmdOptions{}); err != nil {
				handleCmdError(err)
				return nil
			}

			return nil
		},
	}

	return cmd
}

func runRestartCmd(_ restartCmdOptions) error {
	clientConn, _, err := newClient()
	if err != nil {
		return err
	}
	defer clientConn.Close()

	// ctx := context.Background()

	// _, err = client.Restart(ctx, &proto.RestartRequest{})
	// if err != nil {
	// 	return err
	// }

	return nil
}
