package main

import (
	"github.com/spf13/cobra"
)

type applyCmdOptions struct {
	commonCmdOptions
	filename string
}

func newApplyCmd() *cobra.Command {
	applyCmd := &cobra.Command{
		Use: "apply",
		RunE: func(cmd *cobra.Command, args []string) error {
			filenameFlag, err := cmd.Flags().GetString("filename")
			if err != nil {
				return err
			}

			if err := runApplyCmd(applyCmdOptions{
				filename: filenameFlag,
			}); err != nil {
				handleCmdError(err)
				return nil
			}

			return nil
		},
	}

	applyCmd.Flags().StringP("filename", "f", "", "")

	applyCmd.MarkFlagRequired("filename")

	return applyCmd
}

func runApplyCmd(opts applyCmdOptions) error {
	// clientConn, client, err := newClient()
	// if err != nil {
	// 	return err
	// }
	// defer clientConn.Close()

	// data, err := os.ReadFile(opts.filename)
	// if err != nil {
	// 	return err
	// }

	// ctx := context.Background()

	// applyClient, err := client.Apply(ctx, &proto.ApplyRequest{
	// 	Data: data,
	// })
	// if err != nil {
	// 	return err
	// }

	// for {
	// 	reply, err := applyClient.Recv()
	// 	if err != nil {
	// 		if errors.Is(err, io.EOF) {
	// 			break
	// 		}

	// 		return err
	// 	}

	// 	fmt.Println(reply.Message)
	// }

	return nil
}
