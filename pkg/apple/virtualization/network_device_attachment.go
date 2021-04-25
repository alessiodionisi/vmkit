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
