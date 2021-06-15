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

package driver

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/adnsio/vmkit/pkg/config"
)

type QEMU struct {
	executableName string
}

func (q *QEMU) supported() bool {
	path, err := exec.LookPath(q.executableName)
	if err != nil {
		return false
	}

	if path == "" {
		return false
	}

	return true
}

func (d *QEMU) Command(config *config.VirtualMachineV1Alpha1) (*exec.Cmd, error) {
	cmdArgs := []string{
		"-machine", "virt,highmem=off",
		"-cpu", "cortex-a72",
		"-accel", "hvf",
		"-rtc", "base=localtime",
		"-nographic",
	}

	cmdArgs = append(cmdArgs, "-smp", fmt.Sprint(config.Spec.CPU))
	cmdArgs = append(cmdArgs, "-m", config.Spec.Memory)

	for i, disk := range config.Spec.Disks {
		cmdArgs = append(cmdArgs, "-device", fmt.Sprintf("virtio-blk-pci,drive=drive%d", i))
		cmdArgs = append(cmdArgs, "-drive", fmt.Sprintf("if=none,media=disk,id=drive%d,file=%s,cache=writethrough", i, disk.Path))
	}

	for i, network := range config.Spec.Networks {
		cmdArgs = append(cmdArgs, "-device", fmt.Sprintf("virtio-net-pci,netdev=net%d,mac=%s", i, network.MACAddress))
		cmdArgs = append(cmdArgs, "-netdev", fmt.Sprintf("user,id=net%d", i))
	}

	fmt.Printf("%s %s", d.executableName, strings.Join(cmdArgs, " "))

	return exec.Command(
		d.executableName,
		cmdArgs...,
	), nil
}

func NewQEMU(
	executableName string,
) (Driver, error) {
	d := &QEMU{
		executableName: executableName,
	}

	if !d.supported() {
		return nil, ErrNotSupported
	}

	return d, nil
}
