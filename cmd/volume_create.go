package cmd

import (
	"fmt"
	"os"

	"github.com/alessiodionisi/vmkit/bytesize"
	"github.com/alessiodionisi/vmkit/engine"
	"github.com/spf13/cobra"
)

type volumeCreateOptions struct {
	*globalOptions

	name   string
	size   string
	source string
}

func newVolumeCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [name]",
		Args:  cobra.ExactArgs(1),
		Short: "Create a new volume",
		Example: `  Create a volume of 20 gigabytes:
    vmkit volume create some-volume --size 20GB

  Create a volume from a directory, shared using virtiofs:
    vmkit volume create some-volume --source /path/to/directory`,
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			options, err := newVolumeCreateOptions(globalOptions, cmd, args)
			if err != nil {
				return err
			}

			if err := volumeCreate(options); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().String("size", "", "Size of the volume")
	cmd.Flags().String("source", "", "Path to the directory to use as source")

	return cmd
}

func newVolumeCreateOptions(globalOptions *globalOptions, cmd *cobra.Command, args []string) (*volumeCreateOptions, error) {
	size, err := cmd.Flags().GetString("size")
	if err != nil {
		return nil, err
	}

	source, err := cmd.Flags().GetString("source")
	if err != nil {
		return nil, err
	}

	return &volumeCreateOptions{
		globalOptions: globalOptions,
		name:          args[0],
		size:          size,
		source:        source,
	}, nil
}

func volumeCreate(opts *volumeCreateOptions) error {
	e, err := newEngine(opts.globalOptions)
	if err != nil {
		return err
	}

	size, err := bytesize.Parse(opts.size)
	if err != nil {
		return fmt.Errorf("unable to parse size: %w", err)
	}

	_, err = e.CreateVolume(&engine.CreateVolumeOptions{
		Name:   opts.name,
		Size:   size,
		Source: opts.source,
	})

	if err != nil {
		return fmt.Errorf("unable to create volume: %w", err)
	}

	return nil
}
