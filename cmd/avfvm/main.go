// Spin up Linux VMs with QEMU and Apple virtualization framework
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
	"os/signal"
	"strings"
	"time"

	"github.com/adnsio/vmkit/pkg/apple/virtualization"
	"github.com/spf13/cobra"
)

type runOptions struct {
	CPUCount            uint
	DiskImages          []string
	EFI                 string
	EFIVariableStore    string
	LinuxCommandLine    string
	LinuxInitialRamdisk string
	LinuxKernel         string
	MemorySize          uint
	Networks            []string
}

func main() {
	cmd := &cobra.Command{
		Use:   "avfvm",
		Short: "Apple Virtualization.framework Virtual Machine",
		RunE: func(cmd *cobra.Command, args []string) error {
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

			cpuCount, err := cmd.Flags().GetUint("cpu-count")
			if err != nil {
				return err
			}

			memorySize, err := cmd.Flags().GetUint("memory-size")
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

			if err := run(&runOptions{
				CPUCount:            cpuCount,
				DiskImages:          diskImages,
				EFI:                 efi,
				EFIVariableStore:    efiVariableStore,
				LinuxCommandLine:    linuxCommandLine,
				LinuxInitialRamdisk: linuxInitialRamdisk,
				LinuxKernel:         linuxKernel,
				MemorySize:          memorySize,
				Networks:            networks,
			}); err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			return nil
		},
	}

	// virtual machine
	cmd.Flags().UintP("cpu-count", "c", 1, "number of cpu(s) you make available to the guest operating system")
	cmd.Flags().UintP("memory-size", "m", 1073741824, "amount of physical memory the guest operating system sees")

	// linux boot loader
	cmd.Flags().String("linux-kernel", "", "location of the kernel")
	cmd.Flags().String("linux-initial-ramdisk", "", "location of an optional ram disk, which the boot loader maps into memory before it boots the kernel")
	cmd.Flags().String("linux-command-line", "", "command-line parameters to pass to the kernel at boot time")

	// efi boot loader
	cmd.Flags().String("efi", "", "[EXPERIMENTAL] location of the efi image")
	cmd.Flags().String("efi-variable-store", "", "[EXPERIMENTAL] location of the efi variable store")

	// disk images
	cmd.Flags().StringArrayP("disk-image", "d", []string{}, "disk image configuration(s)")

	// networks
	cmd.Flags().StringArrayP("network", "n", []string{"nat"}, "network interface configuration(s)\n  nat[,macAddress=mac]")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(opts *runOptions) error {
	if !virtualization.IsSupported() {
		return errors.New("virtualization is not supported")
	}

	// create the virtual machine configuration
	vmConfig := virtualization.NewVirtualMachineConfiguration()

	// == BOOT LOADER

	switch {
	// linux boot loader
	case opts.LinuxKernel != "":
		log.Printf(
			`boot loader: configuring linux kernel "%s"`,
			opts.LinuxKernel,
		)

		linuxBootLoader := virtualization.NewLinuxBootLoader(opts.LinuxKernel)

		if opts.LinuxInitialRamdisk != "" {
			log.Printf(
				`boot loader: configuring linux initial ram disk "%s"`,
				opts.LinuxInitialRamdisk,
			)
			linuxBootLoader.SetInitialRamdiskURL(opts.LinuxInitialRamdisk)
		}

		if opts.LinuxCommandLine != "" {
			log.Printf(
				`boot loader: configuring linux command line "%s"`,
				opts.LinuxCommandLine,
			)

			linuxBootLoader.SetCommandLine(opts.LinuxCommandLine)
		}

		vmConfig.SetBootLoader(linuxBootLoader)

	// efi boot loader
	case opts.EFI != "":
		log.Printf(
			`boot loader: configuring efi url "%s"`,
			opts.EFI,
		)

		efiBootLoader := virtualization.NewEFIBootLoader()
		efiBootLoader.SetEFIURL(opts.EFI)

		if opts.EFIVariableStore != "" {
			log.Printf(
				`boot loader: configuring efi variable store "%s"`,
				opts.EFIVariableStore,
			)

			efiVarStore, err := virtualization.NewEFIVariableStore(opts.EFIVariableStore)
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
	for _, diskImage := range opts.DiskImages {
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
	for _, network := range opts.Networks {
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
		opts.CPUCount,
		opts.MemorySize,
	)

	vmConfig.SetCPUCount(uint64(opts.CPUCount))
	vmConfig.SetMemorySize(uint64(opts.MemorySize))

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

	go func() {
		interruptCh := make(chan os.Signal, 1)
		signal.Notify(interruptCh, os.Interrupt)

		<-interruptCh
		log.Println("siginal: got an interrupt")
		// TODO: request VM stop
		os.Exit(0)
	}()

	for {
		if vm.State() == virtualization.VirtualMachineStateStopped {
			log.Println("virtual machine: stopped")
			os.Exit(0)
		}

		time.Sleep(1 * time.Second)
	}
}
