package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type sshOptions struct {
	*globalOptions
	name    string
	command bool
}

func newSSHCommand() *cobra.Command {
	cmd := &cobra.Command{
		Example: `  Quickly connect to a virtual machine via SSH on Unix systems:
    $(vmkit ssh vm1 --command)`,
		Args:  cobra.ExactArgs(1),
		Short: "SSH connection details of a running virtual machine",
		Use:   "ssh [name]",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			command, err := cmd.Flags().GetBool("command")
			if err != nil {
				return err
			}

			opts := &sshOptions{
				name:          args[0],
				globalOptions: globalOptions,
				command:       command,
			}

			if err := runSSH(opts); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().Bool("command", false, "print only the ssh command")

	return cmd
}

func runSSH(opts *sshOptions) error {
	eng, err := newEngine(opts.globalOptions)
	if err != nil {
		return err
	}

	vm := eng.FindVirtualMachine(opts.name)
	if vm == nil {
		return fmt.Errorf(`virtual machine "%s" not found`, opts.name)
	}

	sshConnectionDetails, err := vm.SSHConnectionDetails()
	if err != nil {
		return err
	}

	command := fmt.Sprintf(
		"ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -o IdentitiesOnly=yes -p %d -i %s %s@%s",
		sshConnectionDetails.Port,
		sshConnectionDetails.PrivateKey,
		sshConnectionDetails.Username,
		sshConnectionDetails.Host,
	)

	if opts.command {
		fmt.Println(command)

		return nil
	}

	fmt.Println("On Unix systems you can use this quick command to connect via SSH:")
	fmt.Printf("  %s\n\n", command)

	fmt.Println("Or you can use these parameters with your favourite SSH client:")
	fmt.Printf("  Username: %s\n", sshConnectionDetails.Username)
	fmt.Printf("  Host: %s\n", sshConnectionDetails.Host)
	fmt.Printf("  Port: %d\n", sshConnectionDetails.Port)
	fmt.Printf("  Private key: %s\n", sshConnectionDetails.PrivateKey)

	return nil
}
