package main

import (
	"github.com/spf13/cobra"
)

type stopCmdOptions struct {
	commonCmdOptions
}

func newStopCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop [name]",
		Short: "Stop a virtual machine",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := runStopCmd(stopCmdOptions{}); err != nil {
				handleCmdError(err)
				return nil
			}

			return nil
		},
	}

	return cmd
}

func runStopCmd(_ stopCmdOptions) error {
	clientConn, _, err := newClient()
	if err != nil {
		return err
	}
	defer clientConn.Close()

	// ctx := context.Background()

	// _, err = client.Stop(ctx, &proto.StopRequest{})
	// if err != nil {
	// 	return err
	// }

	return nil
}
