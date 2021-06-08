package virtualization

import "unsafe"

type BootLoader interface {
	Pointer() unsafe.Pointer
}
