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

func (d *QEMU) Command(opts *CommandOptions) (*exec.Cmd, error) {
	cmdArgs := []string{
		"-machine", "virt,highmem=off",
		"-cpu", "cortex-a72",
		"-accel", "hvf",
		"-rtc", "base=localtime",
		"-nographic",
		"-device", "qemu-xhci",
	}

	switch {
	case opts.Config.Spec.BootLoader.EFI != nil:
		cmdArgs = append(cmdArgs, "-bios", opts.Config.Spec.BootLoader.EFI.Path)
	default:
		return nil, fmt.Errorf("%w, qemu supports only efi boot loader", config.ErrInvalidBootLoaderConfiguration)
	}

	cmdArgs = append(cmdArgs, "-smp", fmt.Sprint(opts.Config.Spec.CPU))
	cmdArgs = append(cmdArgs, "-m", opts.Config.Spec.Memory)

	for i, disk := range opts.Config.Spec.Disks {
		cmdArgs = append(cmdArgs, "-device", fmt.Sprintf("virtio-blk-pci,drive=drive%d", i))
		cmdArgs = append(cmdArgs, "-drive", fmt.Sprintf("if=none,media=disk,id=drive%d,file=%s,cache=writethrough", i, disk.Path))
	}

	for i, network := range opts.Config.Spec.Networks {
		device := []string{
			"virtio-net-pci",
			fmt.Sprintf("netdev=netdev%d", i),
		}

		if network.MACAddress != "" {
			device = append(device, fmt.Sprintf("mac=%s", network.MACAddress))
		}

		cmdArgs = append(cmdArgs, "-device", strings.Join(device, ","))
		cmdArgs = append(cmdArgs, "-netdev", fmt.Sprintf("user,id=netdev%d,hostfwd=tcp::2222-:22", i))
	}

	if opts.CloudInitISO != "" {
		cmdArgs = append(cmdArgs, "-device", "usb-storage,drive=cloud-init,removable=true")
		cmdArgs = append(cmdArgs, "-drive", fmt.Sprintf("if=none,media=cdrom,id=cloud-init,file=%s,cache=writethrough", opts.CloudInitISO))
	}

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
