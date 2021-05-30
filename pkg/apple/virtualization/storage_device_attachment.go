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
