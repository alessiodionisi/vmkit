// Virtual Machine manager that supports QEMU and Apple virtualization framework on macOS
// Copyright (C) 2021 VMKit Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/adnsio/vmkit/pkg/apple/virtualization"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "avfvm",
		Short: "Apple Virtualization.framework Virtual Machine",
		RunE:  run,
	}

	// virtual machine
	cmd.Flags().Uint64("cpu-count", 1, "number of cpu(s) you make available to the guest operating system")
	cmd.Flags().Uint64("memory-size", 1073741824, "amount of physical memory the guest operating system sees")

	// linux boot loader
	cmd.Flags().String("linux-kernel", "", "location of the kernel")
	cmd.Flags().String("linux-initial-ramdisk", "", "location of an optional ram disk, which the boot loader maps into memory before it boots the kernel")
	cmd.Flags().String("linux-command-line", "", "command-line parameters to pass to the kernel at boot time")

	// efi boot loader
	cmd.Flags().String("efi", "", "[EXPERIMENTAL] location of the efi image")
	cmd.Flags().String("efi-variable-store", "", "[EXPERIMENTAL] location of the efi variable store")

	// disk images
	cmd.Flags().StringArray("disk-image", []string{}, "disk image configuration(s)")

	cmd.Flags().StringArray("network", []string{"nat"}, "network interface configuration(s)\n  nat[,macAddress=mac]")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	linuxKernel, err := cmd.Flags().GetString("linux-kernel")
	if err != nil {
		return err
	}

	linuxInitialRamdisk, err := cmd.Flags().GetString("linux-initial-ramdisk")
	if err != nil {
		return err
	}

	linuxCommandLine, err := cmd.Flags().GetString("linux-command-line")
	if err != nil {
		return err
	}

	diskImages, err := cmd.Flags().GetStringArray("disk-image")
	if err != nil {
		return err
	}

	cpuCount, err := cmd.Flags().GetUint64("cpu-count")
	if err != nil {
		return err
	}

	memorySize, err := cmd.Flags().GetUint64("memory-size")
	if err != nil {
		return err
	}

	efi, err := cmd.Flags().GetString("efi")
	if err != nil {
		return err
	}

	efiVariableStore, err := cmd.Flags().GetString("efi-variable-store")
	if err != nil {
		return err
	}

	networks, err := cmd.Flags().GetStringArray("network")
	if err != nil {
		return err
	}

	if !virtualization.IsSupported() {
		return errors.New("virtualization is not supported")
	}

	// create the virtual machine configuration
	vmConfig := virtualization.NewVirtualMachineConfiguration()

	// == BOOT LOADER

	switch {
	// linux boot loader
	case linuxKernel != "":
		log.Printf(
			`boot loader: configuring linux kernel "%s"`,
			linuxKernel,
		)

		linuxBootLoader := virtualization.NewLinuxBootLoader(linuxKernel)

		if linuxInitialRamdisk != "" {
			log.Printf(
				`boot loader: configuring linux initial ram disk "%s"`,
				linuxInitialRamdisk,
			)
			linuxBootLoader.SetInitialRamdiskURL(linuxInitialRamdisk)
		}

		if linuxCommandLine != "" {
			log.Printf(
				`boot loader: configuring linux command line "%s"`,
				linuxCommandLine,
			)

			linuxBootLoader.SetCommandLine(linuxCommandLine)
		}

		vmConfig.SetBootLoader(linuxBootLoader)

	// efi boot loader
	case efi != "":
		log.Printf(
			`boot loader: configuring efi url "%s"`,
			efi,
		)

		efiBootLoader := virtualization.NewEFIBootLoader()
		efiBootLoader.SetEFIURL(efi)

		if efiVariableStore != "" {
			log.Printf(
				`boot loader: configuring efi variable store "%s"`,
				efiVariableStore,
			)

			efiVarStore, err := virtualization.NewEFIVariableStore(efiVariableStore)
			if err != nil {
				return fmt.Errorf("boot loader error: %w", err)
			}

			efiBootLoader.SetVariableStore(efiVarStore)
		}

		vmConfig.SetBootLoader(efiBootLoader)

	default:
		return errors.New("boot loader error: invalid configuration")
	}

	// == SERIAL PORT

	// create the serial port configuration
	log.Println("serial port: configuring with standard input and standard output")

	serialPortAttachment := virtualization.NewFileHandleSerialPortAttachment(
		os.Stdin,
		os.Stdout,
	)
	serialPortConfiguration := virtualization.NewVirtioConsoleDeviceSerialPortConfiguration(serialPortAttachment)

	vmConfig.SetSerialPorts([]virtualization.SerialPortConfiguration{
		serialPortConfiguration,
	})

	// == STORAGE

	// create an empty storage devices slice
	storageDevices := make([]virtualization.StorageDeviceConfiguration, 0)

	// parse the disk images
	for _, diskImage := range diskImages {
		log.Printf(
			`disk image: configuring "%s"`,
			diskImage,
		)

		attachment, err := virtualization.NewDiskImageStorageDeviceAttachment(
			diskImage,
			false,
		)
		if err != nil {
			return fmt.Errorf("disk image error: %w", err)
		}

		storageDevices = append(storageDevices, virtualization.NewVirtioBlockDeviceConfiguration(attachment))
	}

	vmConfig.SetStorageDevices(storageDevices)

	// == NETWORK

	// create an empty network devices slice
	networkDevices := make([]virtualization.NetworkDeviceConfiguration, 0)

	// configure the network
	for _, network := range networks {
		networkOptions := strings.Split(network, ",")
		networkType := networkOptions[0]
		networkOptions = networkOptions[1:]

		switch networkType {
		case "nat":
			log.Printf("network: configuring nat with options %s\n", networkOptions)

			networkDevice := virtualization.NewVirtioNetworkDeviceConfiguration(
				virtualization.NewNATNetworkDeviceAttachment(),
			)

			for _, networkOption := range networkOptions {
				nameAndValue := strings.Split(networkOption, "=")
				switch nameAndValue[0] {
				case "macAddress":
					networkDevice.SetMACAddress(virtualization.NewMacAddress(nameAndValue[1]))
				}
			}

			networkDevices = append(
				networkDevices,
				networkDevice,
			)
		}
	}

	vmConfig.SetNetworkDevices(networkDevices)

	// == RANDOMIZATION

	// Create a new entropy device
	log.Println("entropy device: configuring")
	entropyDevice := virtualization.NewVirtioEntropyDeviceConfiguration()

	vmConfig.SetEntropyDevices([]virtualization.NetworkDeviceConfiguration{
		entropyDevice,
	})

	// == VIRTUAL MACHINE

	// configure the virtual machine
	log.Printf(
		"virtual machine: configuring with %d cpu and %d bytes memory size",
		cpuCount,
		memorySize,
	)

	vmConfig.SetCPUCount(cpuCount)
	vmConfig.SetMemorySize(memorySize)

	// validate the configuration
	if err := vmConfig.Validate(); err != nil {
		return fmt.Errorf("validate error: %w", err)
	}

	// create the virtual machine
	vm := virtualization.NewVirtualMachine(vmConfig)

	// start the virtual machine
	log.Println("virtual machine: starting")
	if err := vm.Start(); err != nil {
		return fmt.Errorf("start error: %w", err)
	}

	for {
		if vm.State() == virtualization.VirtualMachineStateStopped {
			log.Println("virtual machine: stopped")
			os.Exit(0)
		}

		time.Sleep(1 * time.Second)
	}
}
