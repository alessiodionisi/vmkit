package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var errUnableToDeleteInstance = fmt.Errorf("unable to delete instance")

type instanceDeleteOptions struct {
	*globalOptions

	name string
}

func newInstanceDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [name]",
		Args:  cobra.ExactArgs(1),
		Short: "Delete an instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			options, err := newInstanceDeleteOptions(globalOptions, cmd, args)
			if err != nil {
				return err
			}

			if err := instanceDelete(options); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func newInstanceDeleteOptions(globalOptions *globalOptions, cmd *cobra.Command, args []string) (*instanceDeleteOptions, error) {
	return &instanceDeleteOptions{
		globalOptions: globalOptions,
		name:          args[0],
	}, nil
}

func instanceDelete(opts *instanceDeleteOptions) error {
	e, err := newEngine(opts.globalOptions)
	if err != nil {
		return err
	}

	i, err := e.GetInstance(opts.name)
	if err != nil {
		return fmt.Errorf("%w: %w", errUnableToDeleteInstance, err)
	}

	if err := i.Delete(); err != nil {
		return fmt.Errorf("%w: %w", errUnableToDeleteInstance, err)
	}

	return nil
}
