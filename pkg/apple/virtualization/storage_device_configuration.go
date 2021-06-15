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
