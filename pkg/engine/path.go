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

package engine

import (
	"path"
)

func (e *Engine) imagesPath() string {
	return path.Join(e.path, "image")
}

func (e *Engine) imagePath(name string) string {
	return path.Join(e.imagesPath(), name)
}

func (e *Engine) virtualMachinesPath() string {
	return path.Join(e.path, "virtual-machine")
}

func (e *Engine) virtualMachinePath(name string) string {
	return path.Join(e.virtualMachinesPath(), name)
}
