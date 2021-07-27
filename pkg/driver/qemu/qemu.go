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

package qemu

import (
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"strings"

	"github.com/adnsio/vmkit/pkg/driver"
)

const (
	Aarch64ExecutableName = "qemu-system-aarch64"
	X86_64ExecutableName  = "qemu-system-x86_64"
)

type NewOptions struct {
	ExecutableName  string
	OVMFBiosPath    string
	QEMUEFIBiosPath string
	Writer          io.Writer
}

type QEMU struct {
	accelerator    string
	biosPath       string
	cpu            string
	executableName string
	machine        string
	writer         io.Writer
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

func (d *QEMU) Command(opts *driver.CommandOptions) (*exec.Cmd, error) {
	cmdArgs := []string{
		"-accel", d.accelerator, // enable apple hypervisor.framework acceleration
		"-cpu", d.cpu, // sets the emulated cpu
		"-device", "qemu-xhci", // adds a PCI bus for USB devices
		"-m", fmt.Sprint(opts.Memory), // sets the memory available
		"-machine", d.machine, // sets the emulated machine with highmem=off
		"-nographic",             // use stdio for the serial input and output
		"-rtc", "base=localtime", // sync RTC clock with host time
		"-smp", fmt.Sprint(opts.CPU), // sets the number of CPUs
	}

	if d.biosPath != "" {
		// sets the bios
		cmdArgs = append(cmdArgs, "-bios", d.biosPath)
	}

	for i, disk := range opts.Disks {
		diskArgs := []string{
			"-device", fmt.Sprintf("virtio-blk-pci,drive=drive%d", i), // create a virtio PCI block device
			"-drive", fmt.Sprintf("if=none,media=disk,id=drive%d,file=%s,cache=writethrough", i, disk), // sets the media as disk and load the file
		}

		cmdArgs = append(cmdArgs, diskArgs...)
	}

	if opts.SSHPortForward != 0 {
		networkArgs := []string{
			"-device", "virtio-net-pci,netdev=netdev0", // create a virtio PCI network device
			"-netdev", fmt.Sprintf("user,id=netdev0,hostfwd=tcp::%d-:22", opts.SSHPortForward), // configure port forwarding
		}

		cmdArgs = append(cmdArgs, networkArgs...)
	}

	if opts.CloudInitISO != "" {
		cloudInitArgs := []string{
			"-device", "usb-storage,drive=cloud-init,removable=true", // create a removable USB storage
			"-drive", fmt.Sprintf("if=none,media=cdrom,id=cloud-init,file=%s,cache=writethrough", opts.CloudInitISO), // sets the media as cdrom and load the ISO file
		}

		cmdArgs = append(cmdArgs, cloudInitArgs...)
	}

	return exec.Command(
		d.executableName,
		cmdArgs...,
	), nil
}

func New(opts *NewOptions) (*QEMU, error) {
	qemu := &QEMU{
		executableName: opts.ExecutableName,
		writer:         opts.Writer,
	}

	if !qemu.lookExecutable() {
		return nil, driver.ErrExecutableNotFound
	}

	switch {
	case strings.Contains(qemu.executableName, "aarch64"):
		if runtime.GOARCH != "arm64" {
			return nil, driver.ErrUnsupportedArchitecture
		}

		qemu.biosPath = opts.QEMUEFIBiosPath
		qemu.cpu = "cortex-a72"
		qemu.machine = "virt,highmem=off"

	case strings.Contains(qemu.executableName, "x86_64"):
		if runtime.GOARCH != "amd64" {
			return nil, driver.ErrUnsupportedArchitecture
		}

		qemu.biosPath = opts.OVMFBiosPath
		qemu.cpu = "host"
		qemu.machine = "q35"

	default:
		return nil, driver.ErrUnsupportedArchitecture
	}

	switch runtime.GOOS {
	case "darwin":
		qemu.accelerator = "hvf"

	case "linux":
		qemu.accelerator = "kvm"

	default:
		return nil, driver.ErrUnsupportedOperatingSystem
	}

	return qemu, nil
}
