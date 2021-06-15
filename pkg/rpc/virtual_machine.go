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

package rpc

import (
	"github.com/adnsio/vmkit/pkg/config"
	"github.com/adnsio/vmkit/pkg/engine"
)

type StartVirtualMachineOptions struct {
	Name string
}

type VirtualMachineLogsOptions struct {
	Name string
}

type VirtualMachine struct {
	Config *config.VirtualMachineV1Alpha1
	Status engine.VirtualMachinesStatus
}

type VirtualMachineReceiver struct {
	engine *engine.Engine
}

func (r *VirtualMachineReceiver) Start(opts *StartVirtualMachineOptions, _ *interface{}) error {
	return r.engine.StartVirtualMachine(&engine.StartVirtualMachineOptions{
		Name: opts.Name,
	})
}

func (r *VirtualMachineReceiver) List(_ *interface{}, reply *[]*VirtualMachine) error {
	virtualMachines := r.engine.VirtualMachinesList()

	var res = make([]*VirtualMachine, len(virtualMachines))
	for i, vm := range virtualMachines {
		res[i] = &VirtualMachine{
			Config: vm.Config,
			Status: vm.Status(),
		}
	}

	*reply = res

	return nil
}

func (r *VirtualMachineReceiver) Logs(opts *VirtualMachineLogsOptions, reply *string) error {
	virtualMachine, err := r.engine.FindVirtualMachine(opts.Name)
	if err != nil {
		return err
	}

	logs, err := virtualMachine.Logs()
	if err != nil {
		return err
	}

	*reply = logs

	return nil
}
