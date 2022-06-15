package main

import (
	"github.com/spf13/cobra"
)

type deleteCmdOptions struct {
	commonCmdOptions
}

func newDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [name]",
		Short: "Delete a virtual machine",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := runDeleteCmd(deleteCmdOptions{}); err != nil {
				handleCmdError(err)
				return nil
			}

			return nil
		},
	}

	return cmd
}

func runDeleteCmd(_ deleteCmdOptions) error {
	clientConn, _, err := newClient()
	if err != nil {
		return err
	}
	defer clientConn.Close()

	// ctx := context.Background()

	// _, err = client.Delete(ctx, &proto.DeleteRequest{})
	// if err != nil {
	// 	return err
	// }

	return nil
}
