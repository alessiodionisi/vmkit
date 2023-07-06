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

type CommandOptionsDisk struct {
	Path     string
	ReadOnly bool
}

type CommandOptions struct {
	CPU          int
	Disks        []CommandOptionsDisk
	Memory       int
	MACAddress   string
	PortForwards map[string]string
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

func (q *QEMU) Command(opts CommandOptions) (*exec.Cmd, error) {
	cmdArgs := []string{
		"-machine", q.machine, // sets the machine
		"-m", fmt.Sprint(opts.Memory), // sets the memory available
		"-cpu", q.cpu, // sets the cpu
		"-smp", fmt.Sprint(opts.CPU), // sets the cpu topology
		"-accel", q.accelerator, // sets the accelerator
		"-device", "virtio-rng-pci", // random number generator device
		// "-rtc", "base=localtime", // sync RTC clock with host time
		"-serial", "mon:stdio", // use stdio for the serial input and output
		"-display", "none", // disable display
		"-nodefaults", // remove all defaults
	}

	if q.bios != "" {
		// sets the bios
		cmdArgs = append(cmdArgs, "-bios", q.bios)
	}

	for i, disk := range opts.Disks {
		blockDevOptions := []string{
			fmt.Sprintf("node-name=drive%d", i),
		}

		if strings.HasSuffix(disk.Path, ".qcow2") {
			blockDevOptions = append(
				blockDevOptions,
				"driver=qcow2",
				"file.driver=file",
				fmt.Sprintf("file.filename=%s", disk.Path),
			)
		} else {
			blockDevOptions = append(
				blockDevOptions,
				"driver=raw",
				"file.driver=file",
				fmt.Sprintf("file.filename=%s", disk.Path),
			)
		}

		if disk.ReadOnly {
			blockDevOptions = append(blockDevOptions, "read-only=on")
		}

		diskArgs := []string{
			"-device", fmt.Sprintf("virtio-blk,drive=drive%d", i), // create a virtio PCI block device
			"-blockdev", strings.Join(blockDevOptions, ","), // create a blockdev using the PCI block device
		}

		cmdArgs = append(cmdArgs, diskArgs...)
	}

	networkArgs := []string{
		"-device", fmt.Sprintf("virtio-net-pci,mac=%s,netdev=netdev0", opts.MACAddress), // create a virtio PCI network device
	}

	// configure netdev options
	netdevOptions := []string{
		"user",
		"id=netdev0",
	}

	// configure ports forwarding
	for vmPort, hostPort := range opts.PortForwards {
		netdevOptions = append(netdevOptions, fmt.Sprintf("hostfwd=tcp:127.0.0.1:%s-:%s", hostPort, vmPort))
	}

	// configure netdev
	networkArgs = append(networkArgs, "-netdev", strings.Join(netdevOptions, ","))

	cmdArgs = append(cmdArgs, networkArgs...)

	return exec.Command(
		q.executableName,
		cmdArgs...,
	), nil
}

func New(opts NewOptions) (*QEMU, error) {
	qemu := &QEMU{
		executableName: opts.ExecutableName,
	}

	path, err := exec.LookPath(qemu.executableName)
	if err != nil {
		return nil, ErrExecutableNotFound
	}

	if path == "" {
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
		qemu.cpu = "host"
		qemu.machine = "type=virt"

	case strings.Contains(qemu.executableName, "x86_64"):
		if runtime.GOARCH != "amd64" {
			return nil, ErrX8664Emulation
		}

		qemu.cpu = "host"
		qemu.machine = "type=q35"

	default:
		return nil, ErrUnsupportedArchitecture
	}

	return qemu, nil
}
