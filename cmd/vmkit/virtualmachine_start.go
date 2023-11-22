package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/alessiodionisi/vmkit/proto"
	"github.com/spf13/cobra"
)

type virtualMachineStartOptions struct {
	*globalOptions

	name string
}

func newVirtualMachineStartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start a virtual machine",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			options, err := newVirtualMachineStartOptions(globalOptions, cmd)
			if err != nil {
				return err
			}

			if err := runVirtualMachineStart(options); err != nil {
				slog.Error(fmt.Sprintf("error running virtual machine start command: %v", err))
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().String("name", "", "Name of the virtual machine")

	cmd.MarkFlagRequired("name")

	return cmd
}

func newVirtualMachineStartOptions(globalOptions *globalOptions, cmd *cobra.Command) (*virtualMachineStartOptions, error) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return nil, err
	}

	return &virtualMachineStartOptions{
		globalOptions: globalOptions,
		name:          name,
	}, nil
}

func runVirtualMachineStart(opts *virtualMachineStartOptions) error {
	client, conn, err := newClient(opts.globalOptions)
	if err != nil {
		return err
	}
	defer conn.Close()

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	_, err = client.VirtualMachineStart(
		context.Background(),
		&proto.VirtualMachineStartRequest{
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
