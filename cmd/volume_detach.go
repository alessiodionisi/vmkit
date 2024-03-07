package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var errUnableToDetachVolume = fmt.Errorf("unable to detach volume")

type volumeDetachOptions struct {
	*globalOptions

	name     string
	instance string
}

func newVolumeDetachCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "detach [name]",
		Args:  cobra.ExactArgs(1),
		Short: "Detach a volume",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			options, err := newVolumeDetachOptions(globalOptions, cmd, args)
			if err != nil {
				return err
			}

			if err := volumeDetach(options); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func newVolumeDetachOptions(globalOptions *globalOptions, _ *cobra.Command, args []string) (*volumeDetachOptions, error) {
	return &volumeDetachOptions{
		globalOptions: globalOptions,

		name: args[0],
	}, nil
}

func volumeDetach(opts *volumeDetachOptions) error {
	e, err := newEngine(opts.globalOptions)
	if err != nil {
		return err
	}

	v, err := e.GetVolume(opts.name)
	if err != nil {
		return fmt.Errorf("%w: %w", errUnableToDetachVolume, err)
	}

	if err := v.Detach(); err != nil {
		return fmt.Errorf("%w: %w", errUnableToDetachVolume, err)
	}

	return nil
}
