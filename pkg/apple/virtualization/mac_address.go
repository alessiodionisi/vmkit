package virtualization

/*
#include "mac_address.h"
*/
import "C"
import "unsafe"

type MACAddress struct {
	ptr unsafe.Pointer
}

func (m *MACAddress) Pointer() unsafe.Pointer {
	return m.ptr
}
