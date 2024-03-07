package cmd

import (
	"fmt"
	"os"

	"github.com/alessiodionisi/vmkit/bytesize"
	"github.com/spf13/cobra"
)

func newInstanceListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List virtual machines",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			if err := instanceList(globalOptions); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func instanceList(opts *globalOptions) error {
	e, err := newEngine(opts)
	if err != nil {
		return err
	}

	instances := e.ListInstances()
	rows := make([][]string, 0, len(instances))

	for _, i := range instances {
		rows = append(rows, []string{
			i.Name,
			i.State().String(),
			fmt.Sprintf("%d", i.VCPU),
			bytesize.FormatBinary(i.Memory),
		})
	}

	writeTable(&writeTableOptions{
		header: []string{
			"Name",
			"State",
			"vCPU",
			"Memory",
		},
		rows: rows,
	})

	return nil
}
