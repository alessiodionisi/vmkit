package main

import (
	"github.com/spf13/cobra"
)

type logsCmdOptions struct {
	commonCmdOptions
}

func newLogsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logs [name]",
		Short: "Fetch the logs of a virtual machine",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := runLogsCmd(logsCmdOptions{}); err != nil {
				handleCmdError(err)
				return nil
			}

			return nil
		},
	}

	return cmd
}

func runLogsCmd(_ logsCmdOptions) error {
	clientConn, _, err := newClient()
	if err != nil {
		return err
	}
	defer clientConn.Close()

	// ctx := context.Background()

	// _, err = client.Logs(ctx, &proto.LogsRequest{})
	// if err != nil {
	// 	return err
	// }

	return nil
}
