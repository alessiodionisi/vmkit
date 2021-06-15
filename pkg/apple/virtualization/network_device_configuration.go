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
