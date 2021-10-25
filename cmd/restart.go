package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type restartOptions struct {
	*globalOptions
	name string
}

func newRestartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.ExactArgs(1),
		Short: "Restart a running virtual machine",
		Use:   "restart [name]",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			opts := &restartOptions{
				name:          args[0],
				globalOptions: globalOptions,
			}

			if err := runRestart(opts); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runRestart(opts *restartOptions) error {
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

	if err := vm.Stop(); err != nil {
		return err
	}

	return vm.Start()
}
