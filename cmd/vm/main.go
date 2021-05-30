package main

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/adnsio/vmkit/pkg/apple/virtualization"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:  "vmkit-vm",
		RunE: run,
	}

	// virtual machine
	cmd.Flags().Uint64("cpu-count", 1, "")
	cmd.Flags().Uint64("memory-size", 1073741824, "")

	// linux boot loader
	cmd.Flags().String("linux-kernel", "", "")
	cmd.Flags().String("linux-initial-ramdisk", "", "")
	cmd.Flags().String("linux-command-line", "", "")

	// efi boot loader
	cmd.Flags().String("efi-url", "", "")
	cmd.Flags().String("efi-variable-store-url", "", "")

	cmd.Flags().StringArray("disk-image", []string{}, "")

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

	efiURL, err := cmd.Flags().GetString("efi-url")
	if err != nil {
		return err
	}

	efiVariableStoreURL, err := cmd.Flags().GetString("efi-variable-store-url")
	if err != nil {
		return err
	}

	if !virtualization.IsSupported() {
		return errors.New("virtualization is not supported")
	}

	// create the virtual machine configuration
	vmConfig := virtualization.NewVirtualMachineConfiguration()

	// create the boot loader configuration
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
	case efiURL != "":
		log.Printf(
			`boot loader: configuring efi url "%s"`,
			efiURL,
		)

		efiBootLoader := virtualization.NewEFIBootLoader()
		efiBootLoader.SetEFIURL(efiURL)

		if efiVariableStoreURL != "" {
			log.Printf(
				`boot loader: configuring efi variable store "%s"`,
				efiVariableStoreURL,
			)

			efiVariableStore, err := virtualization.NewEFIVariableStore(efiVariableStoreURL)
			if err != nil {
				return err
			}

			efiBootLoader.SetVariableStore(efiVariableStore)
		}

		vmConfig.SetBootLoader(efiBootLoader)

	default:
		return errors.New("boot loader: invalid configuration")
	}

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
			return err
		}

		storageDevices = append(storageDevices, virtualization.NewVirtioBlockDeviceConfiguration(attachment))
	}

	vmConfig.SetStorageDevices(storageDevices)

	// create the network configuration
	log.Println("network: configuring nat")

	networkAttachment := virtualization.NewNATNetworkDeviceAttachment()
	networkConfiguration := virtualization.NewVirtioNetworkDeviceConfiguration(networkAttachment)

	vmConfig.SetNetworkDevices([]virtualization.NetworkDeviceConfiguration{
		networkConfiguration,
	})

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
		return err
	}

	// create the virtual machine
	vm := virtualization.NewVirtualMachine(vmConfig)

	// start the virtual machine
	log.Println("virtual machine: starting")
	vm.Start()

	for {
		log.Println("================")
		log.Printf("state: %v", vm.State().String())
		log.Printf("canStart: %v", vm.CanStart())
		log.Printf("canPause: %v", vm.CanPause())
		log.Printf("canResume: %v", vm.CanResume())
		log.Printf("canRequestStop: %v", vm.CanRequestStop())

		time.Sleep(5 * time.Second)
	}
}
