package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newMacAddressCommand() *cobra.Command {
	cmd := &cobra.Command{
		Short: "Generate a valid, random, locally administered, unicast MAC address",
		Use:   "macaddress",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			if err := runMacAddress(globalOptions); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runMacAddress(opts *globalOptions) error {
	eng, err := newEngine(opts)
	if err != nil {
		return err
	}

	macAddress, err := eng.RandomLocallyAdministeredMacAddress()
	if err != nil {
		return err
	}

	fmt.Println(macAddress)

	return nil
}
