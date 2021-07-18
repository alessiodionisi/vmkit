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
)

type urlAndHash struct {
	checksum string
	url      string
}

type archData struct {
	disk           *urlAndHash
	initialRamDisk *urlAndHash
	kernel         *urlAndHash
}

type Image struct {
	arch        map[string]*archData
	Description string
	engine      *Engine
	Name        string
	Version     string
}

func (img *Image) Pulled() (bool, error) {
	imgPath, err := img.engine.imagePath(img.NameAndVersion())
	if err != nil {
		return false, err
	}

	if _, err := os.Stat(path.Join(imgPath, "disk.img")); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, err
	}

	if _, err := os.Stat(path.Join(imgPath, "kernel")); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, err
	}

	if _, err := os.Stat(path.Join(imgPath, "initial-ram-disk")); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (img *Image) NameAndVersion() string {
	return fmt.Sprintf("%s:%s", img.Name, img.Version)
}

func (img *Image) Pull() error {
	imagePath, err := img.engine.imagePath(img.NameAndVersion())
	if err != nil {
		return err
	}

	if err := os.MkdirAll(imagePath, 0755); err != nil {
		return err
	}

	diskPath := path.Join(imagePath, "disk.img")
	if err := img.engine.downloadAndPrintProgress(img.arch["arm64"].disk.url, diskPath); err != nil {
		return err
	}

	if err := img.engine.validateChecksum(img.arch["arm64"].disk.checksum, diskPath); err != nil {
		return err
	}

	kernelPath := path.Join(imagePath, "kernel")
	if err := img.engine.downloadAndPrintProgress(img.arch["arm64"].kernel.url, kernelPath); err != nil {
		return err
	}

	if err := img.engine.validateChecksum(img.arch["arm64"].kernel.checksum, kernelPath); err != nil {
		return err
	}

	initialRamDiskPath := path.Join(imagePath, "initial-ram-disk")
	if err := img.engine.downloadAndPrintProgress(img.arch["arm64"].initialRamDisk.url, initialRamDiskPath); err != nil {
		return err
	}

	if err := img.engine.validateChecksum(img.arch["arm64"].initialRamDisk.checksum, initialRamDiskPath); err != nil {
		return err
	}

	return nil
}
