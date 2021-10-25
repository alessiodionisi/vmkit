package qemu

import "errors"

var (
	ErrARM64Emulation             = errors.New("qemu: you are trying to emulate arm64")
	ErrExecutableNotFound         = errors.New("qemu: executable not found")
	ErrUnsupportedArchitecture    = errors.New("qemu: unsupported architecture")
	ErrUnsupportedOperatingSystem = errors.New("qemu: unsupported operating system")
	ErrX8664Emulation             = errors.New("qemu: you are trying to emulate x86_64")
)
