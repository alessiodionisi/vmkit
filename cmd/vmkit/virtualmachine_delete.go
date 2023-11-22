package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/alessiodionisi/vmkit/proto"
	"github.com/spf13/cobra"
)

type virtualMachineDeleteOptions struct {
	*globalOptions

	name string
}

func newVirtualMachineDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a virtual machine",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			options, err := newVirtualMachineDeleteOptions(globalOptions, cmd)
			if err != nil {
				return err
			}

			if err := runVirtualMachineDelete(options); err != nil {
				slog.Error(fmt.Sprintf("error running virtual machine delete command: %v", err))
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().String("name", "", "Name of the virtual machine")

	cmd.MarkFlagRequired("name")

	return cmd
}

func newVirtualMachineDeleteOptions(globalOptions *globalOptions, cmd *cobra.Command) (*virtualMachineDeleteOptions, error) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return nil, err
	}

	return &virtualMachineDeleteOptions{
		globalOptions: globalOptions,
		name:          name,
	}, nil
}

func runVirtualMachineDelete(opts *virtualMachineDeleteOptions) error {
	client, conn, err := newClient(opts.globalOptions)
	if err != nil {
		return err
	}
	defer conn.Close()

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	_, err = client.VirtualMachineDelete(
		context.Background(),
		&proto.VirtualMachineDeleteRequest{
			Name: opts.name,
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
