package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var errUnableToDeleteVolume = fmt.Errorf("unable to delete volume")

type volumeDeleteOptions struct {
	*globalOptions

	name string
}

func newVolumeDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [name]",
		Args:  cobra.ExactArgs(1),
		Short: "Delete a volume",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			options, err := newVolumeDeleteOptions(globalOptions, cmd, args)
			if err != nil {
				return err
			}

			if err := volumeDelete(options); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func newVolumeDeleteOptions(globalOptions *globalOptions, _ *cobra.Command, args []string) (*volumeDeleteOptions, error) {
	return &volumeDeleteOptions{
		globalOptions: globalOptions,

		name: args[0],
	}, nil
}

func volumeDelete(opts *volumeDeleteOptions) error {
	e, err := newEngine(opts.globalOptions)
	if err != nil {
		return err
	}

	v, err := e.GetVolume(opts.name)
	if err != nil {
		return fmt.Errorf("%w: %w", errUnableToDeleteVolume, err)
	}

	if err := v.Delete(); err != nil {
		return fmt.Errorf("%w: %w", errUnableToDeleteVolume, err)
	}

	return nil
}
