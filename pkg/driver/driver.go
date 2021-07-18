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
	"os/exec"

	"github.com/adnsio/vmkit/pkg/config"
)

type CommandOptions struct {
	Config       *config.VirtualMachineV1Alpha1
	CloudInitISO string
}

type DriverType string

const (
	DriverTypeAVFVM DriverType = "avfvm"
	DriverTypeQEMU  DriverType = "qemu"
)

type Driver interface {
	Command(opts *CommandOptions) (*exec.Cmd, error)
}
