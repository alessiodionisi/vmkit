package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/alessiodionisi/vmkit/proto"
	"github.com/spf13/cobra"
)

func newVirtualMachineListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List virtual machines",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			if err := runVirtualMachineList(globalOptions); err != nil {
				slog.Error(fmt.Sprintf("error running virtual machine list command: %v", err))
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runVirtualMachineList(opts *globalOptions) error {
	client, conn, err := newClient(opts)
	if err != nil {
		return err
	}
	defer conn.Close()

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	_, err = client.VirtualMachineList(
		context.Background(),
		&proto.VirtualMachineListRequest{},
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
