package main

import (
	"github.com/spf13/cobra"
)

type pullImageCmdOptions struct {
	commonCmdOptions
}

func newPullImageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pull [name]",
		Short: "Pull an image",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := runPullImageCmd(pullImageCmdOptions{}); err != nil {
				handleCmdError(err)
				return nil
			}

			return nil
		},
	}

	return cmd
}

func runPullImageCmd(_ pullImageCmdOptions) error {
	clientConn, _, err := newClient()
	if err != nil {
		return err
	}
	defer clientConn.Close()

	// ctx := context.Background()

	// _, err = client.PullImage(ctx, &proto.PullImageRequest{})
	// if err != nil {
	// 	return err
	// }

	return nil
}
