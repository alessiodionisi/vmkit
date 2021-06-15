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
#include "foundation.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type StorageDeviceAttachment interface {
	Pointer() unsafe.Pointer
}

// DiskImageStorageDeviceAttachment manage the storage for a disk in the virtual machine. The guest operating system sees the storage as a disk.
type DiskImageStorageDeviceAttachment struct {
	StorageDeviceAttachment
	ptr unsafe.Pointer
}

// Pointer returns the objective-c pointer.
func (a *DiskImageStorageDeviceAttachment) Pointer() unsafe.Pointer {
	return a.ptr
}

// NewDiskImageStorageDeviceAttachment creates the attachment object from the specified disk image.
//
// Docs: https://developer.apple.com/documentation/virtualization/vzdiskimagestoragedeviceattachment
func NewDiskImageStorageDeviceAttachment(
	diskImageURL string,
	readOnly bool,
) (*DiskImageStorageDeviceAttachment, error) {
	var err unsafe.Pointer

	attachment := &DiskImageStorageDeviceAttachment{
		ptr: C.VZDiskImageStorageDeviceAttachment_init(
			C.CString(diskImageURL),
			C.bool(readOnly),
			&err,
		),
	}

	if err != nil {
		return nil, errors.New(C.GoString(C.NSError_localizedDescription(err)))
	}

	return attachment, nil
}
