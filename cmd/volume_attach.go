package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var errUnableToAttachVolume = fmt.Errorf("unable to attach volume")

type volumeAttachOptions struct {
	*globalOptions

	name     string
	instance string
}

func newVolumeAttachCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attach [name] [instance]",
		Args:  cobra.ExactArgs(2),
		Short: "Attach a volume to an instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			options, err := newVolumeAttachOptions(globalOptions, cmd, args)
			if err != nil {
				return err
			}

			if err := volumeAttach(options); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func newVolumeAttachOptions(globalOptions *globalOptions, _ *cobra.Command, args []string) (*volumeAttachOptions, error) {
	return &volumeAttachOptions{
		globalOptions: globalOptions,

		name:     args[0],
		instance: args[1],
	}, nil
}

func volumeAttach(opts *volumeAttachOptions) error {
	e, err := newEngine(opts.globalOptions)
	if err != nil {
		return err
	}

	v, err := e.GetVolume(opts.name)
	if err != nil {
		return fmt.Errorf("%w: %w", errUnableToAttachVolume, err)
	}

	if err := v.Attach(opts.instance); err != nil {
		return fmt.Errorf("%w: %w", errUnableToAttachVolume, err)
	}

	return nil
}
