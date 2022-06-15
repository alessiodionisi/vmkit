package main

import (
	"context"

	"github.com/adnsio/vmkit/proto"
	"github.com/spf13/cobra"
)

type createCmdOptions struct {
	commonCmdOptions
}

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [name]",
		Short: "Create a new virtual machine",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := runCreateCmd(createCmdOptions{}); err != nil {
				handleCmdError(err)
				return nil
			}

			return nil
		},
	}

	cmd.Flags().StringP("image", "i", "", "image to use")

	cmd.MarkFlagRequired("image")

	return cmd
}

func runCreateCmd(_ createCmdOptions) error {
	clientConn, client, err := newClient()
	if err != nil {
		return err
	}
	defer clientConn.Close()

	ctx := context.Background()

	_, err = client.Create(ctx, &proto.CreateRequest{})
	if err != nil {
		return err
	}

	return nil
}
