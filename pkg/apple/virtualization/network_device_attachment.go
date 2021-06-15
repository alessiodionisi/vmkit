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

type NetworkDeviceAttachment interface {
	Pointer() unsafe.Pointer
}

// NATNetworkDeviceAttachment is a device that routes network requests through the host computer and performs network address translation on the resulting packets.
type NATNetworkDeviceAttachment struct {
	NetworkDeviceAttachment
	ptr unsafe.Pointer
}

// Pointer returns the objective-c pointer.
func (a *NATNetworkDeviceAttachment) Pointer() unsafe.Pointer {
	return a.ptr
}

// NewNATNetworkDeviceAttachment creates an attachment that performs network address translation on the guest system’s network packets.
func NewNATNetworkDeviceAttachment() *NATNetworkDeviceAttachment {
	return &NATNetworkDeviceAttachment{
		ptr: C.VZNATNetworkDeviceAttachment_init(),
	}
}
