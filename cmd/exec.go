package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type execOptions struct {
	*globalOptions
	name    string
	command string
}

func newExecCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args: cobra.MinimumNArgs(2),
		Example: `  Execute "uname -a" in the virtual machine:
    vmkit exec vm1 -- uname -a`,
		Short: "Execute a command in a running virtual machine",
		Use:   "exec [name] [command]",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			opts := &execOptions{
				name:          args[0],
				command:       strings.Join(args[1:], " "),
				globalOptions: globalOptions,
			}

			if err := runExec(opts); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().Bool("command", false, "prints the ssh command to connect")

	return cmd
}

func runExec(opts *execOptions) error {
	eng, err := newEngine(opts.globalOptions)
	if err != nil {
		return err
	}

	vm := eng.FindVirtualMachine(opts.name)
	if vm == nil {
		return fmt.Errorf(`virtual machine "%s" not found`, opts.name)
	}

	return vm.Exec(opts.command)
}
