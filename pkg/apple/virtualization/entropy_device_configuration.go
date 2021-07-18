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
