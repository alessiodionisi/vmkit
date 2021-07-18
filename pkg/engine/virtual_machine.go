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

package engine

import (
	"errors"
	"os"
	"path"
	"strconv"
	"strings"
	"syscall"

	"github.com/adnsio/vmkit/pkg/driver"
)

var (
	ErrVirtualMachineAlreadyRunning = errors.New("virtual machine is already running")
	ErrVirtualMachineNotRunning     = errors.New("virtual machine is not running")
)

type VirtualMachineStatus string

const (
	VirtualMachineStatusStopped VirtualMachineStatus = "stopped"
	VirtualMachineStatusRunning VirtualMachineStatus = "running"
	VirtualMachineStatusError   VirtualMachineStatus = "error"
)

type VirtualMachine struct {
	engine *Engine
	Name   string
}

func (vm *VirtualMachine) Status() (VirtualMachineStatus, error) {
	virtualMachinePath, err := vm.engine.virtualMachinePath(vm.Name)
	if err != nil {
		return VirtualMachineStatusError, err
	}

	// read pid file
	pidFileBytes, err := os.ReadFile(path.Join(virtualMachinePath, "pid"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return VirtualMachineStatusStopped, nil
		}

		return VirtualMachineStatusError, err
	}

	pid, err := strconv.Atoi(strings.ReplaceAll(string(pidFileBytes), "\n", ""))
	if err != nil {
		return VirtualMachineStatusError, err
	}

	// find process
	proc, err := os.FindProcess(pid)
	if err != nil {
		return VirtualMachineStatusError, err
	}

	// check if is running with a signal
	if err := proc.Signal(syscall.Signal(0)); err != nil {
		return VirtualMachineStatusStopped, nil
	}

	return VirtualMachineStatusRunning, nil
}

func (vm *VirtualMachine) Start() error {
	// get status
	status, err := vm.Status()
	if err != nil {
		return err
	}

	if status == VirtualMachineStatusRunning {
		return ErrVirtualMachineAlreadyRunning
	}

	// start command
	cmd, err := vm.engine.driver.Command(&driver.CommandOptions{})
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	// write pid to file
	virtualMachinePath, err := vm.engine.virtualMachinePath(vm.Name)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(virtualMachinePath, 0755); err != nil {
		return err
	}

	if err := os.WriteFile(
		path.Join(virtualMachinePath, "pid"),
		[]byte(strconv.Itoa(cmd.Process.Pid)),
		0666,
	); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) Stop() error {
	status, err := vm.Status()
	if err != nil {
		return err
	}

	if status == VirtualMachineStatusRunning {
		return ErrVirtualMachineNotRunning
	}

	return nil
}

func (vm *VirtualMachine) Remove() error {
	return nil
}
