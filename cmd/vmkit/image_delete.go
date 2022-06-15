package main

import (
	"github.com/spf13/cobra"
)

type deleteImageCmdOptions struct {
	commonCmdOptions
}

func newDeleteImageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [name]",
		Short: "Delete an image",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := runDeleteImageCmd(deleteImageCmdOptions{}); err != nil {
				handleCmdError(err)
				return nil
			}

			return nil
		},
	}

	return cmd
}

func runDeleteImageCmd(_ deleteImageCmdOptions) error {
	clientConn, _, err := newClient()
	if err != nil {
		return err
	}
	defer clientConn.Close()

	// ctx := context.Background()

	// _, err = client.DeleteImage(ctx, &proto.DeleteImageRequest{})
	// if err != nil {
	// 	return err
	// }

	return nil
}
