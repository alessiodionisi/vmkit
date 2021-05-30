package virtualization

/*
#include "boot_loader_private.h"
#include "foundation.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

// EFIBootLoader
type EFIBootLoader struct {
	BootLoader
	ptr unsafe.Pointer
}

// Pointer returns the objective-c pointer.
func (b *EFIBootLoader) Pointer() unsafe.Pointer {
	return b.ptr
}

// SetEFIURL
func (b *EFIBootLoader) SetEFIURL(efiURL string) {
	C.VZEFIBootLoader_setEfiURL(b.ptr, C.CString(efiURL))
}

// SetVariableStore
func (b *EFIBootLoader) SetVariableStore(variableStore *EFIVariableStore) {
	C.VZEFIBootLoader_setVariableStore(b.ptr, variableStore.ptr)
}

// NewEFIBootLoader
func NewEFIBootLoader() *EFIBootLoader {
	return &EFIBootLoader{
		ptr: C.VZEFIBootLoader_init(),
	}
}

// EFIVariableStore
type EFIVariableStore struct {
	ptr unsafe.Pointer
}

// NewEFIVariableStore
func NewEFIVariableStore(url string) (*EFIVariableStore, error) {
	var errPtr unsafe.Pointer

	variableStore := &EFIVariableStore{
		ptr: C.VZEFIVariableStore_init(C.CString(url), &errPtr),
	}

	if errPtr != nil {
		return nil, errors.New(C.GoString(C.NSError_localizedDescription(errPtr)))
	}

	return variableStore, nil
}
