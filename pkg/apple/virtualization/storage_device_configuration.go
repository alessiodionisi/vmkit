package virtualization

/*
#include "storage_device.h"
*/
import "C"
import "unsafe"

type StorageDeviceConfiguration interface {
	Pointer() unsafe.Pointer
}

// VirtioBlockDeviceConfiguration create an emulated storage device in your virtual machine.
type VirtioBlockDeviceConfiguration struct {
	StorageDeviceConfiguration
	ptr unsafe.Pointer
}

// Pointer returns the objective-c pointer.
func (c *VirtioBlockDeviceConfiguration) Pointer() unsafe.Pointer {
	return c.ptr
}

// NewVirtioBlockDeviceConfiguration creates a block device configuration object that uses the specified storage medium.
//
// Docs: https://developer.apple.com/documentation/virtualization/vzvirtioblockdeviceconfiguration
func NewVirtioBlockDeviceConfiguration(attachment StorageDeviceAttachment) *VirtioBlockDeviceConfiguration {
	return &VirtioBlockDeviceConfiguration{
		ptr: C.VZVirtioBlockDeviceConfiguration_init(attachment.Pointer()),
	}
}
