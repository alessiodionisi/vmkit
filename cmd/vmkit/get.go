package main

import (
	"github.com/spf13/cobra"
)

type getCmdOptions struct {
	commonCmdOptions
	resource string
}

func newGetCmd() *cobra.Command {
	getCmd := &cobra.Command{
		Use:  "get",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := runGetCmd(getCmdOptions{
				resource: args[0],
			}); err != nil {
				handleCmdError(err)
				return nil
			}

			return nil
		},
	}

	return getCmd
}

func runGetCmd(opts getCmdOptions) error {
	// clientConn, client, err := newClient()
	// if err != nil {
	// 	return err
	// }
	// defer clientConn.Close()

	// ctx := context.Background()

	// _, err = client.Get(ctx, &proto.GetRequest{
	// 	Resource: opts.resource,
	// })
	// if err != nil {
	// 	return err
	// }

	return nil
}
