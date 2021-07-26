// Spin up Linux VMs with QEMU and Apple virtualization framework
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
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/adnsio/vmkit/pkg/engine/bios/qemuefi"
)

func (eng *Engine) qemuEFIBiosPath() string {
	return path.Join(eng.biosPath(), "QEMU_EFI.fd")
}

func (eng *Engine) checkAndWriteBiosFiles() error {
	biosPath := eng.biosPath()
	_, err := os.Stat(biosPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := os.MkdirAll(biosPath, 0755); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	qemuEFIBiosPath := eng.qemuEFIBiosPath()
	_, err = os.Stat(qemuEFIBiosPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Fprintln(eng.writer, "Extracting qemu efi bios")

			if err := os.WriteFile(qemuEFIBiosPath, qemuefi.Bytes, 0644); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}
