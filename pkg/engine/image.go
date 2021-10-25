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
	"errors"
	"os"
	"path"
	"runtime"
)

type imageURLAndChecksum struct {
	checksum string
	url      string
}

type imageArchitecture struct {
	disk *imageURLAndChecksum
}

type Image struct {
	Description string
	Name        string
	Version     string

	arch   map[string]*imageArchitecture
	engine *Engine
	path   string
}

func (i *Image) makePath() error {
	return os.MkdirAll(i.path, 0755)
}

func (i *Image) Pulled() (bool, error) {
	if _, err := os.Stat(i.diskPath()); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (i *Image) diskPath() string {
	return path.Join(i.path, "disk.qcow2")
}

func (i *Image) Pull() error {
	i.engine.Printf("Pulling image \"%s:%s\"\n", i.Name, i.Version)

	if err := i.makePath(); err != nil {
		return err
	}

	arch, exist := i.arch[runtime.GOARCH]
	if !exist {
		return ErrUnsupportedArchitecture
	}

	if err := i.engine.downloadAndPrintProgress(arch.disk.url, i.diskPath()); err != nil {
		return err
	}

	// if err := img.engine.validateChecksum(arch.disk.checksum, img.diskPath()); err != nil {
	// 	return err
	// }

	return nil
}

func (e *Engine) FindImage(name string) *Image {
	image, exist := e.images[name]
	if !exist {
		return nil
	}

	return image
}

func (eng *Engine) ListImages() []*Image {
	images := make([]*Image, 0, len(eng.images))
	for _, img := range eng.images {
		images = append(images, img)
	}

	return images
}
