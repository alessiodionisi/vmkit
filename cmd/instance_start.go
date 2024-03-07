package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type virtualMachineStartOptions struct {
	*globalOptions

	name string
}

func newInstanceStartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start a virtual machine",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			options, err := newVirtualMachineStartOptions(globalOptions, cmd)
			if err != nil {
				return err
			}

			if err := virtualMachineStart(options); err != nil {
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

func newVirtualMachineStartOptions(globalOptions *globalOptions, cmd *cobra.Command) (*virtualMachineStartOptions, error) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return nil, err
	}

	return &virtualMachineStartOptions{
		globalOptions: globalOptions,
		name:          name,
	}, nil
}

func virtualMachineStart(opts *virtualMachineStartOptions) error {
	_, err := newEngine(opts.globalOptions)
	if err != nil {
		return err
	}

	return nil
}
