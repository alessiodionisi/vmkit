package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type xtermOptions struct {
	*globalOptions
	name string
}

func newXtermCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.ExactArgs(1),
		Short: "Connect via SSH to a running virtual machine",
		Use:   "xterm [name]",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			opts := &xtermOptions{
				name:          args[0],
				globalOptions: globalOptions,
			}

			if err := runXterm(opts); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runXterm(opts *xtermOptions) error {
	eng, err := newEngine(opts.globalOptions)
	if err != nil {
		return err
	}

	vm := eng.FindVirtualMachine(opts.name)
	if vm == nil {
		return fmt.Errorf(`virtual machine "%s" not found`, opts.name)
	}

	return vm.SSHSessionWithXterm()
}
