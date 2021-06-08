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
