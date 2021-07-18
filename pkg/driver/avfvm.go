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
	"encoding/base64"
	"fmt"
	"os/exec"
	"strings"

	"github.com/adnsio/vmkit/pkg/config"
)

const (
	AVFVMExecutableName = "avfvm"
)

type AVFVM struct {
	executableName string
}

func (d *AVFVM) Command(opts *CommandOptions) (*exec.Cmd, error) {
	cmdArgs := []string{}

	if opts.Config.Spec.BootLoader.Linux == nil {
		return nil, fmt.Errorf("%w, linux boot loader is mandatory with avfvm", config.ErrInvalidBootLoaderConfiguration)
	}

	cmdArgs = append(cmdArgs, "--cpu-count", fmt.Sprint(opts.Config.Spec.CPU))
	cmdArgs = append(cmdArgs, "--linux-kernel", opts.Config.Spec.BootLoader.Linux.Kernel)
	cmdArgs = append(cmdArgs, "--linux-initial-ramdisk", opts.Config.Spec.BootLoader.Linux.InitialRamdisk)

	var linuxCommandLine []string

	if opts.Config.Spec.BootLoader.Linux.CommandLine != "" {
		linuxCommandLine = append(linuxCommandLine, opts.Config.Spec.BootLoader.Linux.CommandLine)
	}

	if opts.Config.Spec.CloudInit != nil {
		// CLOUD INIT - DATA SOURCE
		linuxCommandLine = append(linuxCommandLine, fmt.Sprintf("ds=nocloud;h=%s;i=%s", opts.Config.Metadata.Name, opts.Config.Metadata.Name))

		// CLOUD INIT - NETWORK CONFIGURATION
		if opts.Config.Spec.CloudInit.NetworkConfiguration != "" {
			linuxCommandLine = append(
				linuxCommandLine,
				fmt.Sprintf("network-config=%s",
					base64.StdEncoding.EncodeToString([]byte(opts.Config.Spec.CloudInit.NetworkConfiguration)),
				),
			)
		}

		// CLOUD INIT - USER DATA
		if opts.Config.Spec.CloudInit.UserData != "" {
			linuxCommandLine = append(linuxCommandLine, fmt.Sprintf("cc: %s end_cc", opts.Config.Spec.CloudInit.UserData))
		}
	}

	linuxCommandLineString := strings.Join(linuxCommandLine, " ")

	if !strings.Contains(linuxCommandLineString, "console=hvc0") {
		return nil, fmt.Errorf(
			`%w, linux command line "console=hvc0" is mandatory with avfvm`,
			config.ErrInvalidBootLoaderConfiguration,
		)
	}

	if opts.Config.Spec.CPU > 1 && !strings.Contains(linuxCommandLineString, "irqaffinity=0") {
		return nil, fmt.Errorf(
			`%w, linux command line "irqaffinity=0" is needed to fix sync problems with more than one cpu with avfvm`,
			config.ErrInvalidBootLoaderConfiguration,
		)
	}

	cmdArgs = append(cmdArgs, "--linux-command-line", linuxCommandLineString)

	for _, disk := range opts.Config.Spec.Disks {
		cmdArgs = append(cmdArgs, "--disk-image", disk.Path)
	}

	for _, network := range opts.Config.Spec.Networks {
		cmdArgs = append(cmdArgs, "--network", fmt.Sprintf("nat,macAddress=%s", network.MACAddress))
	}

	return exec.Command(
		d.executableName,
		cmdArgs...,
	), nil
}

func (d *AVFVM) supported() bool {
	path, err := exec.LookPath(d.executableName)
	if err != nil {
		return false
	}

	if path == "" {
		return false
	}

	return true
}

func NewAVFVM(
	executableName string,
) (Driver, error) {
	d := &AVFVM{
		executableName: executableName,
	}

	if !d.supported() {
		return nil, ErrNotSupported
	}

	return d, nil
}
