package cmd

import (
	"github.com/spf13/cobra"
)

func newVolumeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "volume",
		Short: "Manage volumes",
	}

	cmd.AddCommand(newVolumeAttachCommand())
	cmd.AddCommand(newVolumeCreateCommand())
	cmd.AddCommand(newVolumeDeleteCommand())
	cmd.AddCommand(newVolumeDetachCommand())
	cmd.AddCommand(newVolumeListCommand())

	return cmd
}
