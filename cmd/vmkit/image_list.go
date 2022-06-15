package main

import (
	"context"
	"strings"

	"github.com/adnsio/vmkit/proto"
	"github.com/spf13/cobra"
)

type listImagesCmdOptions struct {
	commonCmdOptions
}

func newListImagesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all images",
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := runListImagesCmd(listImagesCmdOptions{}); err != nil {
				handleCmdError(err)
				return nil
			}

			return nil
		},
	}

	return cmd
}

func runListImagesCmd(_ listImagesCmdOptions) error {
	clientConn, client, err := newClient()
	if err != nil {
		return err
	}
	defer clientConn.Close()

	ctx := context.Background()

	res, err := client.ListImages(ctx, &proto.ListImagesRequest{})
	if err != nil {
		return err
	}

	rows := make([][]string, len(res.Images))
	for i, img := range res.Images {
		rows[i] = []string{
			img.Name,
			img.Description,
			strings.Join(img.Archs, ", "),
			"-",
		}
	}

	writeTable(&writeTableOptions{
		Header: []string{
			"Name",
			"Description",
			"Archs",
			"Status",
		},
		Rows: rows,
	})

	return nil
}
