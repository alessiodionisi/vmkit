package virtualization

/*
#include "entropy_device_configuration.h"
*/
import "C"
import "unsafe"

// EntropyDeviceConfiguration is the common behaviors for entropy devices.
type EntropyDeviceConfiguration interface {
	Pointer() unsafe.Pointer
}

// VirtioEntropyDeviceConfiguration expose a source of entropy for the guest operating system’s random-number generator.
type VirtioEntropyDeviceConfiguration struct {
	EntropyDeviceConfiguration
	ptr unsafe.Pointer
}

// Pointer returns the objective-c pointer.
func (c *VirtioEntropyDeviceConfiguration) Pointer() unsafe.Pointer {
	return c.ptr
}

// NewVirtioEntropyDeviceConfiguration creates an entropy device configuration object.
//
// Docs: https://developer.apple.com/documentation/virtualization/vzvirtioentropydeviceconfiguration
func NewVirtioEntropyDeviceConfiguration() *VirtioEntropyDeviceConfiguration {
	return &VirtioEntropyDeviceConfiguration{
		ptr: C.VZVirtioEntropyDeviceConfiguration_init(),
	}
}
