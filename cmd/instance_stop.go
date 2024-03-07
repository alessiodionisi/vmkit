package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type virtualMachineStopOptions struct {
	*globalOptions

	name string
}

func newInstanceStopCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop a virtual machine",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			options, err := newVirtualMachineStopOptions(globalOptions, cmd)
			if err != nil {
				return err
			}

			if err := virtualMachineStop(options); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().String("name", "", "Name of the virtual machine")

	cmd.MarkFlagRequired("name")

	return cmd
}

func newVirtualMachineStopOptions(globalOptions *globalOptions, cmd *cobra.Command) (*virtualMachineStopOptions, error) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return nil, err
	}

	return &virtualMachineStopOptions{
		globalOptions: globalOptions,
		name:          name,
	}, nil
}

func virtualMachineStop(opts *virtualMachineStopOptions) error {
	return nil
}
