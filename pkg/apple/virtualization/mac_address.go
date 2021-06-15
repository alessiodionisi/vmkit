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
#include "mac_address.h"
*/
import "C"
import "unsafe"

// MACAddress is the media access control (MAC) address for a network interface in your virtual machine.
type MACAddress struct {
	ptr unsafe.Pointer
}

// Pointer returns the objective-c pointer.
func (m *MACAddress) Pointer() unsafe.Pointer {
	return m.ptr
}

// NewMacAddress creates a MAC address object from a specially formatted string.
func NewMacAddress(value string) *MACAddress {
	return &MACAddress{
		ptr: C.VZMACAddress_init(C.CString(value)),
	}
}
