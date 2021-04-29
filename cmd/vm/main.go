package main

import (
	"log"
	"os"
	"time"

	"github.com/adnsio/vmkit/pkg/apple/virtualization"
)

func main() {
	// create the boot loader configuration
	bootLoader := virtualization.NewLinuxBootLoader("./bin/alpine/vmlinuz-lts")
	bootLoader.SetInitialRamdiskURL("./bin/alpine/initramfs-lts")
	bootLoader.SetCommandLine("console=hvc0")

	// log.Println("CommandLine", bootLoader.CommandLine())

	// create the serial port configuration
	serialPortAttachment := virtualization.NewFileHandleSerialPortAttachment(
		os.Stdin,
		os.Stdout,
	)
	serialPortConfiguration := virtualization.NewVirtioConsoleDeviceSerialPortConfiguration(serialPortAttachment)

	// create the iso disk image configuration
	isoDiskImageAttachment, err := virtualization.NewDiskImageStorageDeviceAttachment(
		"./bin/alpine/alpine-standard-3.13.5-aarch64.iso",
		true,
	)
	if err != nil {
		log.Fatal(err)
	}
	isoDiskImageConfiguration := virtualization.NewVirtioBlockDeviceConfiguration(isoDiskImageAttachment)

	// create the disk image configuration
	// diskImageAttachment, err := virtualization.NewDiskImageStorageDeviceAttachment(
	// 	"./bin/alpine/disk.img",
	// 	false,
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// diskImageConfiguration := virtualization.NewVirtioBlockDeviceConfiguration(diskImageAttachment)

	// create the network configuration
	networkAttachment := virtualization.NewNATNetworkDeviceAttachment()
	networkConfiguration := virtualization.NewVirtioNetworkDeviceConfiguration(networkAttachment)

	// create the virtual machine configuration
	vmConfig := virtualization.NewVirtualMachineConfiguration()
	vmConfig.SetBootLoader(bootLoader)
	vmConfig.SetCPUCount(4)
	vmConfig.SetMemorySize(4 * 1024 * 1024 * 1024)
	vmConfig.SetSerialPorts([]virtualization.SerialPortConfiguration{
		serialPortConfiguration,
	})
	vmConfig.SetStorageDevices([]virtualization.StorageDeviceConfiguration{
		// diskImageConfiguration,
		isoDiskImageConfiguration,
	})
	vmConfig.SetNetworkDevices([]virtualization.NetworkDeviceConfiguration{
		networkConfiguration,
	})

	// log.Println("MinimumAllowedCPUCount", vmConfig.MinimumAllowedCPUCount())
	// log.Println("MaximumAllowedCPUCount", vmConfig.MaximumAllowedCPUCount())
	// log.Println("MinimumAllowedMemorySize", vmConfig.MinimumAllowedMemorySize())
	// log.Println("MaximumAllowedMemorySize", vmConfig.MaximumAllowedMemorySize())
	// log.Println("CPUCount", vmConfig.CPUCount())
	// log.Println("MemorySize", vmConfig.MemorySize())

	// validate the configuration
	if err := vmConfig.Validate(); err != nil {
		log.Fatal(err)
	}

	// create the virtual machine
	vm := virtualization.NewVirtualMachine(vmConfig)

	// log.Println("CanStart", vm.CanStart())

	// start the virtual machine
	vm.Start()

	for {
		time.Sleep(1 * time.Second)
	}
}
