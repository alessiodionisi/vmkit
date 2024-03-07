package cmd

import (
	"fmt"
	"os"

	"github.com/alessiodionisi/vmkit/bytesize"
	"github.com/spf13/cobra"
)

func newVolumeListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List volumes",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			if err := volumeList(globalOptions); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func volumeList(opts *globalOptions) error {
	e, err := newEngine(opts)
	if err != nil {
		return err
	}

	volumes := e.ListVolumes()
	rows := make([][]string, 0, len(volumes))

	for _, v := range volumes {
		rows = append(rows, []string{
			v.Name,
			bytesize.FormatBinary(v.Size),
			v.Instance,
		})
	}

	writeTable(&writeTableOptions{
		header: []string{
			"Name",
			"Size",
			"Instance",
		},
		rows: rows,
	})

	return nil
}
