package virtualization

/*
#include "virtual_machine.h"
#include "foundation.h"
*/
import "C"
import (
	"unsafe"
)

// VirtualMachine emulates a complete hardware machine of the same architecture as the underlying Mac computer.
type VirtualMachine struct {
	ptr      unsafe.Pointer
	queuePtr unsafe.Pointer
}

// CanStart returns a value that indicates whether you can start the virtual machine.
func (m *VirtualMachine) CanStart() bool {
	return bool(C.VZVirtualMachine_canStart(m.ptr, m.queuePtr))
}

// Start the virtual machine and notifies the specified completion handler if startup was successful.
func (m *VirtualMachine) Start() {
	C.VZVirtualMachine_startWithCompletionHandler(m.ptr, m.queuePtr)
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
