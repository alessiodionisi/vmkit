package virtualization

/*
#include "configuration.h"
#include "foundation.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

// VirtualMachineConfiguration configure the environment of your virtual machine.
type VirtualMachineConfiguration struct {
	ptr unsafe.Pointer
}

// Pointer returns the objective-c pointer.
func (c *VirtualMachineConfiguration) Pointer() unsafe.Pointer {
	return c.ptr
}

// SetBootLoader sets the guest system to boot when the virtual machine starts.
func (c *VirtualMachineConfiguration) SetBootLoader(bootLoader BootLoader) {
	C.VZVirtualMachineConfiguration_setBootLoader(c.ptr, bootLoader.Pointer())
}

// SetCPUCount sets the number of CPUs you make available to the guest operating system.
func (c *VirtualMachineConfiguration) SetCPUCount(cpuCount uint64) {
	C.VZVirtualMachineConfiguration_setCPUCount(c.ptr, C.ulong(cpuCount))
}

// CPUCount returns the number of CPUs you make available to the guest operating system.
func (c *VirtualMachineConfiguration) CPUCount() uint64 {
	return uint64(C.VZVirtualMachineConfiguration_CPUCount(c.ptr))
}

// SetMemorySize sets the amount of physical memory the guest operating system sees.
func (c *VirtualMachineConfiguration) SetMemorySize(memorySize uint64) {
	C.VZVirtualMachineConfiguration_setMemorySize(c.ptr, C.ulonglong(memorySize))
}

// MemorySize returns the amount of physical memory the guest operating system sees.
func (c *VirtualMachineConfiguration) MemorySize() uint64 {
	return uint64(C.VZVirtualMachineConfiguration_memorySize(c.ptr))
}

// Validate the current configuration settings and reports any issues that might prevent the successful initialization of the virtual machine.
func (c *VirtualMachineConfiguration) Validate() error {
	var errPtr unsafe.Pointer

	if res := bool(C.VZVirtualMachineConfiguration_validateWithError(c.ptr, &errPtr)); !res {
		return errors.New(C.GoString(C.NSError_localizedDescription(errPtr)))
	}

	return nil
}

// SetSerialPorts sets the array of serial ports that you expose to the guest operating system.
func (c *VirtualMachineConfiguration) SetSerialPorts(serialPorts []SerialPortConfiguration) {
	serialPortsArray := C.NSMutableArray_arrayWithCapacity(C.ulong(len(serialPorts)))
	for _, sp := range serialPorts {
		C.NSMutableArray_addObject(serialPortsArray, sp.Pointer())
	}

	C.VZVirtualMachineConfiguration_setSerialPorts(c.ptr, serialPortsArray)
}

// SetStorageDevices sets the array of storage devices that you expose to the guest operating system.
func (c *VirtualMachineConfiguration) SetStorageDevices(storageDevices []StorageDeviceConfiguration) {
	storageDevicesArray := C.NSMutableArray_arrayWithCapacity(C.ulong(len(storageDevices)))
	for _, sd := range storageDevices {
		C.NSMutableArray_addObject(storageDevicesArray, sd.Pointer())
	}

	C.VZVirtualMachineConfiguration_setStorageDevices(c.ptr, storageDevicesArray)
}

// SetNetworkDevices sets the array of network devices that you expose to the guest operating system.
func (c *VirtualMachineConfiguration) SetNetworkDevices(networkDevices []NetworkDeviceConfiguration) {
	networkDevicesArray := C.NSMutableArray_arrayWithCapacity(C.ulong(len(networkDevices)))
	for _, nd := range networkDevices {
		C.NSMutableArray_addObject(networkDevicesArray, nd.Pointer())
	}

	C.VZVirtualMachineConfiguration_setNetworkDevices(c.ptr, networkDevicesArray)
}

// MinimumAllowedCPUCount returns the minimum number of CPUs you may configure for the virtual machine.
func (c VirtualMachineConfiguration) MinimumAllowedCPUCount() uint32 {
	panic(errors.New("unimplemented"))
}

// MaximumAllowedCPUCount returns the maximum number of CPUs you may configure for the virtual machine.
func (c VirtualMachineConfiguration) MaximumAllowedCPUCount() uint32 {
	panic(errors.New("unimplemented"))
}

// MinimumAllowedMemorySize returns the minimum amount of memory that you may configure for the virtual machine.
func (c VirtualMachineConfiguration) MinimumAllowedMemorySize() uint64 {
	panic(errors.New("unimplemented"))
}

// MaximumAllowedMemorySize returns the maximum amount of memory that you may configure for the virtual machine.
func (c VirtualMachineConfiguration) MaximumAllowedMemorySize() uint64 {
	panic(errors.New("unimplemented"))
}

// NewVirtualMachineConfiguration creates a new empty VirtualMachineConfiguration.
//
// Docs: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration
func NewVirtualMachineConfiguration() *VirtualMachineConfiguration {
	return &VirtualMachineConfiguration{
		ptr: C.VZVirtualMachineConfiguration_init(),
	}
}
