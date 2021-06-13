package virtualization

/*
#include "network_device.h"
*/
import "C"
import "unsafe"

type NetworkDeviceConfiguration interface {
	Pointer() unsafe.Pointer
}

// VirtioNetworkDeviceConfiguration requests the creation of a network device for the guest system.
type VirtioNetworkDeviceConfiguration struct {
	NetworkDeviceConfiguration
	ptr unsafe.Pointer
}

// Pointer returns the objective-c pointer.
func (c *VirtioNetworkDeviceConfiguration) Pointer() unsafe.Pointer {
	return c.ptr
}

// SetMACAddress sets the media access control (MAC) address to assign to the network device.
func (c *VirtioNetworkDeviceConfiguration) SetMACAddress(macAddress *MACAddress) {
	C.VZVirtioNetworkDeviceConfiguration_setMACAddress(c.ptr, macAddress.Pointer())
}

// NewVirtioNetworkDeviceConfiguration creates a network device configuration object for you to configure.
//
// Docs: https://developer.apple.com/documentation/virtualization/vzvirtionetworkdeviceconfiguration
func NewVirtioNetworkDeviceConfiguration(attachment NetworkDeviceAttachment) *VirtioNetworkDeviceConfiguration {
	return &VirtioNetworkDeviceConfiguration{
		ptr: C.VZVirtioNetworkDeviceConfiguration_init(attachment.Pointer()),
	}
}
