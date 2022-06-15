package main

import (
	"github.com/spf13/cobra"
)

type listCmdOptions struct {
	commonCmdOptions
}

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all virtual machines",
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := runListCmd(listCmdOptions{}); err != nil {
				handleCmdError(err)
				return nil
			}

			return nil
		},
	}

	return cmd
}

func runListCmd(_ listCmdOptions) error {
	clientConn, _, err := newClient()
	if err != nil {
		return err
	}
	defer clientConn.Close()

	// ctx := context.Background()

	// _, err = client.ListImage(ctx, &proto.ListImageRequest{})
	// if err != nil {
	// 	return err
	// }

	return nil
}
