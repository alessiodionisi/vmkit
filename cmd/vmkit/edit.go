package main

import (
	"github.com/spf13/cobra"
)

type editCmdOptions struct {
	commonCmdOptions
}

func newEditCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit [name]",
		Short: "Edit a virtual machine",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := runEditCmd(editCmdOptions{}); err != nil {
				handleCmdError(err)
				return nil
			}

			return nil
		},
	}

	return cmd
}

func runEditCmd(_ editCmdOptions) error {
	clientConn, _, err := newClient()
	if err != nil {
		return err
	}
	defer clientConn.Close()

	// ctx := context.Background()

	// _, err = client.Edit(ctx, &proto.EditRequest{})
	// if err != nil {
	// 	return err
	// }

	return nil
}
