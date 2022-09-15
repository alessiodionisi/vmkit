package main

import (
	"context"

	"github.com/adnsio/vmkit/proto"
	"github.com/spf13/cobra"
)

type listCmdOptions struct {
	commonCmdOptions
}

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all virtual machines",
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := runListCmd(listCmdOptions{}); err != nil {
				handleCmdError(err)
				return nil
			}

			return nil
		},
	}

	return cmd
}

func runListCmd(_ listCmdOptions) error {
	clientConn, client, err := newClient()
	if err != nil {
		return err
	}
	defer clientConn.Close()

	ctx := context.Background()

	res, err := client.List(ctx, &proto.ListRequest{})
	if err != nil {
		return err
	}

	rows := make([][]string, len(res.VirtualMachines))
	for i, vm := range res.VirtualMachines {
		rows[i] = []string{
			vm.Name,
			vm.Status,
		}
	}

	writeTable(&writeTableOptions{
		Header: []string{
			"Name",
			"Status",
		},
		Rows: rows,
	})

	return nil
}
