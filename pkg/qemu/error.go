// Spin up Linux VMs with QEMU
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

package qemu

import "errors"

var (
	ErrARM64Emulation             = errors.New("qemu: you are trying to emulate arm64")
	ErrExecutableNotFound         = errors.New("qemu: executable not found")
	ErrUnsupportedArchitecture    = errors.New("qemu: unsupported architecture")
	ErrUnsupportedOperatingSystem = errors.New("qemu: unsupported operating system")
	ErrX8664Emulation             = errors.New("qemu: you are trying to emulate x86_64")
)
