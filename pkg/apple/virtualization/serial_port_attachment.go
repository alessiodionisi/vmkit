package virtualization

/*
#include "serial_port.h"
*/
import "C"
import (
	"os"
	"unsafe"
)

type SerialPortAttachment interface {
	Pointer() unsafe.Pointer
}

// FileHandleSerialPortAttachment configure a serial port using separate file handles for reading and writing data.
type FileHandleSerialPortAttachment struct {
	SerialPortAttachment
	ptr unsafe.Pointer
}

// Pointer returns the objective-c pointer.
func (a *FileHandleSerialPortAttachment) Pointer() unsafe.Pointer {
	return a.ptr
}

// NewFileHandleSerialPortAttachment creates a serial port attachment object from the specified file handles.
//
// Docs: https://developer.apple.com/documentation/virtualization/vzfilehandleserialportattachment
func NewFileHandleSerialPortAttachment(
	read *os.File,
	write *os.File,
) *FileHandleSerialPortAttachment {
	return &FileHandleSerialPortAttachment{
		ptr: C.VZFileHandleSerialPortAttachment_init(
			C.int(read.Fd()),
			C.int(write.Fd()),
		),
	}
}
