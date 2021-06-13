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
