package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type pullOptions struct {
	name          string
	globalOptions *globalOptions
}

func newPullCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.ExactArgs(1),
		Short: "Pull an image",
		Use:   "pull [name]",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			opts := &pullOptions{
				name:          args[0],
				globalOptions: globalOptions,
			}

			if err := runPull(opts); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runPull(opts *pullOptions) error {
	eng, err := newEngine(opts.globalOptions)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	img := eng.FindImage(opts.name)
	if img == nil {
		return fmt.Errorf(`image "%s" not found`, opts.name)
	}

	return img.Pull()
}
