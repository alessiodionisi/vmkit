package main

import (
	"fmt"

	"github.com/adnsio/vmkit/pkg/macaddress"
	"github.com/spf13/cobra"
)

func newMacAddressCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mac-address",
		Short: "Generate a random unicast locally administered address",
		RunE: func(cmd *cobra.Command, args []string) error {
			macAddress, err := macaddress.NewUnicastLocallyAdministeredMACAddress()
			if err != nil {
				return err
			}

			fmt.Println(macAddress)

			return nil
		},
	}

	return cmd
}
