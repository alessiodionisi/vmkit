package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/alessiodionisi/vmkit/proto"
	"github.com/spf13/cobra"
)

type virtualMachineCreateOptions struct {
	*globalOptions

	name string
}

func newVirtualMachineCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new virtual machine",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			options, err := newVirtualMachineCreateOptions(globalOptions, cmd)
			if err != nil {
				return err
			}

			if err := runVirtualMachineCreate(options); err != nil {
				slog.Error(fmt.Sprintf("error running virtual machine create command: %v", err))
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().String("name", "", "Name of the virtual machine")

	cmd.MarkFlagRequired("name")

	return cmd
}

func newVirtualMachineCreateOptions(globalOptions *globalOptions, cmd *cobra.Command) (*virtualMachineCreateOptions, error) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return nil, err
	}

	return &virtualMachineCreateOptions{
		globalOptions: globalOptions,
		name:          name,
	}, nil
}

func runVirtualMachineCreate(opts *virtualMachineCreateOptions) error {
	client, conn, err := newClient(opts.globalOptions)
	if err != nil {
		return err
	}
	defer conn.Close()

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	_, err = client.VirtualMachineCreate(
		context.Background(),
		&proto.VirtualMachineCreateRequest{
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
