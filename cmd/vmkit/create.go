package main

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/adnsio/vmkit/proto"
	"github.com/spf13/cobra"
)

type createCmdOptions struct {
	commonCmdOptions

	name  string
	image string
}

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [name]",
		Short: "Create a new virtual machine",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			image, err := cmd.Flags().GetString("image")
			if err != nil {
				return err
			}

			if err := runCreateCmd(createCmdOptions{
				name:  name,
				image: image,
			}); err != nil {
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

func runCreateCmd(opts createCmdOptions) error {
	clientConn, client, err := newClient()
	if err != nil {
		return err
	}
	defer clientConn.Close()

	ctx := context.Background()

	str, err := client.Create(ctx, &proto.CreateRequest{
		Name:  opts.name,
		Image: opts.image,
	})
	if err != nil {
		return err
	}

	for {
		res, err := str.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return err
		}

		switch res.Step {
		case proto.CreateResponse_DOWNLOADING_IMAGE:
			fmt.Println("Downloading image")
		case proto.CreateResponse_RESIZING_DISK:
			fmt.Println("Resizing disk")
		case proto.CreateResponse_STARTING:
			fmt.Println("Starting")
		}

	}

	return nil
}
