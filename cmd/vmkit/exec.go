package main

import (
	"github.com/spf13/cobra"
)

type execCmdOptions struct {
	commonCmdOptions
}

func newExecCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exec [name]",
		Short: "Run a command in a running virtual machine",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := runExecCmd(execCmdOptions{}); err != nil {
				handleCmdError(err)
				return nil
			}

			return nil
		},
	}

	return cmd
}

func runExecCmd(_ execCmdOptions) error {
	clientConn, _, err := newClient()
	if err != nil {
		return err
	}
	defer clientConn.Close()

	// ctx := context.Background()

	// _, err = client.Exec(ctx, &proto.ExecRequest{})
	// if err != nil {
	// 	return err
	// }

	return nil
}
