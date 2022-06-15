package main

import "github.com/spf13/cobra"

type commonCmdOptions struct {
}

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "vmkit",
	}

	rootCmd.PersistentFlags().StringP("server", "s", "", "the address and port of the server")

	rootCmd.AddCommand(newImageCmd())

	rootCmd.AddCommand(newCreateCmd())
	rootCmd.AddCommand(newDeleteCmd())
	rootCmd.AddCommand(newListCmd())
	rootCmd.AddCommand(newLogsCmd())
	rootCmd.AddCommand(newRestartCmd())
	rootCmd.AddCommand(newStartCmd())
	rootCmd.AddCommand(newStopCmd())
	rootCmd.AddCommand(newExecCmd())
	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(newEditCmd())

	return rootCmd
}
