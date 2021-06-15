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
#include "linux_boot_loader.h"
*/
import "C"
import (
	"unsafe"
)

// LinuxBootLoader specify the location of the Linux kernel that serves as the guest operating system.
type LinuxBootLoader struct {
	BootLoader
	ptr unsafe.Pointer
}

// Pointer returns the objective-c pointer.
func (b *LinuxBootLoader) Pointer() unsafe.Pointer {
	return b.ptr
}

// SetCommandLine sets the command-line parameters to pass to the Linux kernel at boot time.
func (b *LinuxBootLoader) SetCommandLine(commandLine string) {
	C.VZLinuxBootLoader_setCommandLine(b.ptr, C.CString(commandLine))
}

// CommandLine returns the command-line parameters to pass to the Linux kernel at boot time.
func (b *LinuxBootLoader) CommandLine() string {
	return C.GoString(C.VZLinuxBootLoader_commandLine(b.ptr))
}

// SetInitialRamdiskURL sets the location of an optional RAM disk, which the boot loader maps into memory before it boots the Linux kernel.
func (b *LinuxBootLoader) SetInitialRamdiskURL(initialRamdiskURL string) {
	C.VZLinuxBootLoader_setInitialRamdiskURL(b.ptr, C.CString(initialRamdiskURL))
}

// NewLinuxBootLoader creates a boot loader that launches the Linux kernel at the specified URL.
//
// Docs: https://developer.apple.com/documentation/virtualization/vzlinuxbootloader
func NewLinuxBootLoader(kernelURL string) *LinuxBootLoader {
	return &LinuxBootLoader{
		ptr: C.VZLinuxBootLoader_init(C.CString(kernelURL)),
	}
}
