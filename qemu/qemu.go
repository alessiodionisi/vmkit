package qemu

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

const (
	Aarch64ExecutableName = "qemu-system-aarch64"
	X8664ExecutableName   = "qemu-system-x86_64"
)

type CommandOptions struct {
	CDRoms         []string
	CPU            int
	Disks          []string
	Memory         int
	SSHPortForward int
}

type NewOptions struct {
	ExecutableName string
}

type QEMU struct {
	accelerator    string
	bios           string
	cpu            string
	executableName string
	machine        string
}

func (q *QEMU) lookExecutable() bool {
	path, err := exec.LookPath(q.executableName)
	if err != nil {
		return false
	}

	if path == "" {
		return false
	}

	return true
}

func (q *QEMU) Command(opts *CommandOptions) (*exec.Cmd, error) {
	cmdArgs := []string{
		"-accel", q.accelerator, // enable acceleration
		"-cpu", q.cpu, // sets the emulated cpu
		"-m", fmt.Sprint(opts.Memory), // sets the memory available
		"-machine", q.machine, // sets the emulated machine with highmem=off
		"-nographic",             // use stdio for the serial input and output
		"-rtc", "base=localtime", // sync RTC clock with host time
		"-smp", fmt.Sprint(opts.CPU), // sets the number of CPUs
		"-device", "qemu-xhci", // adds a bus for USB devices
	}

	if q.bios != "" {
		// sets the bios
		cmdArgs = append(cmdArgs, "-bios", q.bios)
	}

	for i, disk := range opts.Disks {
		diskArgs := []string{
			"-device", fmt.Sprintf("virtio-blk,drive=drive%d", i), // create a virtio PCI block device
			"-drive", fmt.Sprintf("if=none,media=disk,id=drive%d,file=%s,cache=writethrough", i, disk), // sets the media as disk and load the file
		}

		cmdArgs = append(cmdArgs, diskArgs...)
	}

	for i, cdrom := range opts.CDRoms {
		cdromArgs := []string{
			"-device", fmt.Sprintf("usb-storage,drive=cdrom%d,removable=true", i), // create a removable USB storage
			"-drive", fmt.Sprintf("if=none,media=cdrom,id=cdrom%d,file=%s,cache=writethrough", i, cdrom), // sets the media as cdrom and load the ISO file
		}

		cmdArgs = append(cmdArgs, cdromArgs...)
	}

	if opts.SSHPortForward != 0 {
		networkArgs := []string{
			"-device", "virtio-net-pci,netdev=netdev0", // create a virtio PCI network device
			"-netdev", fmt.Sprintf("user,id=netdev0,hostfwd=tcp::%d-:22", opts.SSHPortForward), // configure port forwarding
		}

		cmdArgs = append(cmdArgs, networkArgs...)
	}

	return exec.Command(
		q.executableName,
		cmdArgs...,
	), nil
}

func New(opts *NewOptions) (*QEMU, error) {
	qemu := &QEMU{
		executableName: opts.ExecutableName,
	}

	if !qemu.lookExecutable() {
		return nil, ErrExecutableNotFound
	}

	switch runtime.GOOS {
	case "darwin":
		qemu.accelerator = "hvf"

	case "linux":
		qemu.accelerator = "kvm"

	default:
		return nil, ErrUnsupportedOperatingSystem
	}

	switch {
	case strings.Contains(qemu.executableName, "aarch64"):
		if runtime.GOARCH != "arm64" {
			return nil, ErrARM64Emulation
		}

		qemu.bios = "edk2-aarch64-code.fd"
		qemu.cpu = "cortex-a72"
		qemu.machine = "virt,highmem=off"

	case strings.Contains(qemu.executableName, "x86_64"):
		if runtime.GOARCH != "amd64" {
			return nil, ErrX8664Emulation
		}

		qemu.cpu = "host"
		qemu.machine = "q35"

	default:
		return nil, ErrUnsupportedArchitecture
	}

	return qemu, nil
}
