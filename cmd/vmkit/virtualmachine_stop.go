package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/alessiodionisi/vmkit/proto"
	"github.com/spf13/cobra"
)

type virtualMachineStopOptions struct {
	*globalOptions

	name string
}

func newVirtualMachineStopCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop a virtual machine",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			options, err := newVirtualMachineStopOptions(globalOptions, cmd)
			if err != nil {
				return err
			}

			if err := runVirtualMachineStop(options); err != nil {
				slog.Error(fmt.Sprintf("error running virtual machine stop command: %v", err))
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().String("name", "", "Name of the virtual machine")

	cmd.MarkFlagRequired("name")

	return cmd
}

func newVirtualMachineStopOptions(globalOptions *globalOptions, cmd *cobra.Command) (*virtualMachineStopOptions, error) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return nil, err
	}

	return &virtualMachineStopOptions{
		globalOptions: globalOptions,
		name:          name,
	}, nil
}

func runVirtualMachineStop(opts *virtualMachineStopOptions) error {
	client, conn, err := newClient(opts.globalOptions)
	if err != nil {
		return err
	}
	defer conn.Close()

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	_, err = client.VirtualMachineStop(
		context.Background(),
		&proto.VirtualMachineStopRequest{
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
