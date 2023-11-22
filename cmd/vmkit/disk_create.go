package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/alessiodionisi/vmkit/proto"
	"github.com/alessiodionisi/vmkit/size"
	"github.com/spf13/cobra"
)

type diskCreateOptions struct {
	*globalOptions

	name string
	size string
}

func newDiskCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new disk",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			options, err := newDiskCreateOptions(globalOptions, cmd)
			if err != nil {
				return err
			}

			if err := runDiskCreate(options); err != nil {
				slog.Error(fmt.Sprintf("error running disk create command: %v", err))
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().String("name", "", "Name of the disk")
	cmd.Flags().String("size", "", "Size of the disk")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("size")

	return cmd
}

func newDiskCreateOptions(globalOptions *globalOptions, cmd *cobra.Command) (*diskCreateOptions, error) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return nil, err
	}

	size, err := cmd.Flags().GetString("size")
	if err != nil {
		return nil, err
	}

	return &diskCreateOptions{
		globalOptions: globalOptions,
		name:          name,
		size:          size,
	}, nil
}

func runDiskCreate(opts *diskCreateOptions) error {
	client, conn, err := newClient(opts.globalOptions)
	if err != nil {
		return err
	}
	defer conn.Close()

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	size, err := size.ToBytes(opts.size)
	if err != nil {
		return err
	}

	_, err = client.DiskCreate(
		context.Background(),
		&proto.DiskCreateRequest{
			Name: opts.name,
			Size: size,
		},
	)
	if err != nil {
		return err
	}

	// for {
	// 	response, err := stream.Recv()
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}

	// 		return err
	// 	}

	// 	fmt.Println(*response.Message)
	// }

	return nil
}
