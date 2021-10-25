package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type removeOptions struct {
	*globalOptions
	name string
}

func newRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.ExactArgs(1),
		Short: "Remove a virtual machine",
		Use:   "remove [name]",
		Aliases: []string{
			"rm",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			opts := &removeOptions{
				name:          args[0],
				globalOptions: globalOptions,
			}

			if err := runRemove(opts); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runRemove(opts *removeOptions) error {
	eng, err := newEngine(opts.globalOptions)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	vm := eng.FindVirtualMachine(opts.name)
	if vm == nil {
		return fmt.Errorf(`virtual machine "%s" not found`, opts.name)
	}

	return vm.Remove()
}
