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
#include "efi_boot_loader.h"
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
