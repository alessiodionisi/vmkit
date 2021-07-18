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

package virtualization

/*
#include "virtual_machine.h"
#include "foundation.h"
*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/rs/xid"
)

type VirtualMachineState int

func (s VirtualMachineState) String() string {
	switch s {
	case VirtualMachineStateStopped:
		return "stopped"
	case VirtualMachineStateRunning:
		return "running"
	case VirtualMachineStatePaused:
		return "paused"
	case VirtualMachineStateError:
		return "error"
	case VirtualMachineStateStarting:
		return "starting"
	case VirtualMachineStatePausing:
		return "pausing"
	case VirtualMachineStateResuming:
		return "resuming"
	default:
		return ""
	}
}

var (
	errorHandlers = map[string]func(err error){}
)

const (
	VirtualMachineStateStopped VirtualMachineState = iota
	VirtualMachineStateRunning
	VirtualMachineStatePaused
	VirtualMachineStateError
	VirtualMachineStateStarting
	VirtualMachineStatePausing
	VirtualMachineStateResuming
)

// VirtualMachine emulates a complete hardware machine of the same architecture as the underlying Mac computer.
type VirtualMachine struct {
	ptr      unsafe.Pointer
	queuePtr unsafe.Pointer
}

// IsSupported returns a value that indicates whether the system supports virtualization.
func IsSupported() bool {
	return bool(C.VZVirtualMachine_isSupported())
}

// CanStart returns a value that indicates whether you can start the virtual machine.
func (m *VirtualMachine) CanStart() bool {
	return bool(C.VZVirtualMachine_canStart(m.ptr, m.queuePtr))
}

// CanPause returns a value that indicates whether you can pause the virtual machine.
func (m *VirtualMachine) CanPause() bool {
	return bool(C.VZVirtualMachine_canPause(m.ptr, m.queuePtr))
}

// CanResume returns a value that indicates whether you can resume the virtual machine.
func (m *VirtualMachine) CanResume() bool {
	return bool(C.VZVirtualMachine_canResume(m.ptr, m.queuePtr))
}

// CanRequestStop returns a value that indicates whether you can ask the guest operating system to stop running.
func (m *VirtualMachine) CanRequestStop() bool {
	return bool(C.VZVirtualMachine_canRequestStop(m.ptr, m.queuePtr))
}

// State is the current execution state of the virtual machine.
func (m *VirtualMachine) State() VirtualMachineState {
	return VirtualMachineState(C.VZVirtualMachine_state(m.ptr, m.queuePtr))
}

// Start the virtual machine.
func (m *VirtualMachine) Start() error {
	errCh := make(chan error, 1)
	handlerID := xid.New().String()

	errorHandlers[handlerID] = func(err error) {
		errCh <- err
	}
	defer delete(errorHandlers, handlerID)

	C.VZVirtualMachine_start(m.ptr, m.queuePtr, C.CString(handlerID))

	return <-errCh
}

// Pause a running virtual machine.
func (m *VirtualMachine) Pause() error {
	C.VZVirtualMachine_pause(m.ptr, m.queuePtr)

	return nil
}

// Resume a virtual machine.
func (m *VirtualMachine) Resume() error {
	C.VZVirtualMachine_resume(m.ptr, m.queuePtr)

	return nil
}

// NewVirtualMachine creates the virtual machine and configures it with the specified data.
//
// Docs: https://developer.apple.com/documentation/virtualization/vzvirtualmachine
func NewVirtualMachine(configuration *VirtualMachineConfiguration) *VirtualMachine {
	queue := C.newDispatchQueue(C.CString(""))

	return &VirtualMachine{
		ptr:      C.VZVirtualMachine_init(configuration.Pointer(), queue),
		queuePtr: queue,
	}
}

//export startErrorHandler
func startErrorHandler(err unsafe.Pointer, handlerID *C.char) {
	hID := C.GoString(handlerID)

	if err != nil {
		errorHandlers[hID](errors.New(C.GoString(C.NSError_localizedDescription(err))))
	} else {
		errorHandlers[hID](nil)
	}
}
