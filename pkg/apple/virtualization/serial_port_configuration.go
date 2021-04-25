package virtualization

/*
#include "serial_port.h"
*/
import "C"
import (
	"unsafe"
)

type SerialPortConfiguration interface {
	Pointer() unsafe.Pointer
}

// VirtioConsoleDeviceSerialPortConfiguration enables serial communication between the guest operating system and host computer through the Virtio interface.
type VirtioConsoleDeviceSerialPortConfiguration struct {
	SerialPortConfiguration
	ptr unsafe.Pointer
}

// Pointer returns the objective-c pointer.
func (c *VirtioConsoleDeviceSerialPortConfiguration) Pointer() unsafe.Pointer {
	return c.ptr
}

// NewVirtioConsoleDeviceSerialPortConfiguration creates a serial port configuration object.
//
// Docs: https://developer.apple.com/documentation/virtualization/vzvirtioconsoledeviceserialportconfiguration
func NewVirtioConsoleDeviceSerialPortConfiguration(attachment SerialPortAttachment) *VirtioConsoleDeviceSerialPortConfiguration {
	return &VirtioConsoleDeviceSerialPortConfiguration{
		ptr: C.VZVirtioConsoleDeviceSerialPortConfiguration_init(attachment.Pointer()),
	}
}
