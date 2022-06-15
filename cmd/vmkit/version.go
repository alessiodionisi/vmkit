package main

import (
	"github.com/spf13/cobra"
)

type versionCmdOptions struct {
	commonCmdOptions
}

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the client and server version information",
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := runVersionCmd(versionCmdOptions{}); err != nil {
				handleCmdError(err)
				return nil
			}

			return nil
		},
	}

	return cmd
}

func runVersionCmd(_ versionCmdOptions) error {
	clientConn, _, err := newClient()
	if err != nil {
		return err
	}
	defer clientConn.Close()

	// ctx := context.Background()

	// _, err = client.Version(ctx, &proto.VersionRequest{})
	// if err != nil {
	// 	return err
	// }

	return nil
}
