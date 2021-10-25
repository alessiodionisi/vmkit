package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type stopOptions struct {
	*globalOptions
	name string
}

func newStopCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.ExactArgs(1),
		Short: "Stop a running virtual machine",
		Use:   "stop [name]",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			opts := &stopOptions{
				name:          args[0],
				globalOptions: globalOptions,
			}

			if err := runStop(opts); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runStop(opts *stopOptions) error {
	eng, err := newEngine(opts.globalOptions)
	if err != nil {
		return err
	}

	vm := eng.FindVirtualMachine(opts.name)
	if vm == nil {
		return fmt.Errorf(`virtual machine "%s" not found`, opts.name)
	}

	return vm.Stop()
}
