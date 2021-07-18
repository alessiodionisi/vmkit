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
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/adnsio/vmkit/pkg/driver"
	"github.com/cheggaaa/pb/v3"
)

type NewOptions struct {
	Writer io.Writer
	Driver driver.Driver
}

type Engine struct {
	virtualMachines map[string]*VirtualMachine
	images          map[string]*Image
	writer          io.Writer
	driver          driver.Driver
}

func (eng *Engine) FindVirtualMachine(name string) *VirtualMachine {
	virtualMachine, exist := eng.virtualMachines[name]
	if !exist {
		return nil
	}

	return virtualMachine
}

func (eng *Engine) ListVirtualMachines() []*VirtualMachine {
	virtualMachines := make([]*VirtualMachine, 0, len(eng.virtualMachines))
	for _, vm := range eng.virtualMachines {
		virtualMachines = append(virtualMachines, vm)
	}

	return virtualMachines
}

func (eng *Engine) FindImage(name string) *Image {
	image, exist := eng.images[name]
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

func (eng *Engine) validateChecksum(checksum string, name string) error {
	fmt.Fprintf(eng.writer, "Validating checksum of %s\n", name)

	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	fileChecksum := sha256.Sum256(fileBytes)
	hexFileChecksum := hex.EncodeToString(fileChecksum[:])

	if checksum != hexFileChecksum {
		return fmt.Errorf(`checksum "%s" is not valid, expected "%s"`, hexFileChecksum, checksum)
	}

	return nil
}

func (eng *Engine) downloadAndPrintProgress(url string, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	// fetch content length with an head request
	headResponse, err := http.Head(url)
	if err != nil {
		return err
	}
	defer headResponse.Body.Close()

	contentLength, err := strconv.Atoi(headResponse.Header.Get("content-length"))
	if err != nil {
		return err
	}

	fmt.Fprintf(eng.writer, "Downloading %s\n", url)

	// download the content
	progressBar := pb.Full.New(contentLength)
	progressBar.SetWriter(eng.writer)

	progressBar.Start()

	getResponse, err := http.Get(url)
	if err != nil {
		return err
	}
	defer getResponse.Body.Close()

	progressReader := progressBar.NewProxyReader(getResponse.Body)

	if _, err := io.Copy(file, progressReader); err != nil {
		return err
	}

	progressBar.Finish()

	return nil
}

func (eng *Engine) reloadVirtualMachines() error {
	virtualMachinesPath, err := eng.virtualMachinesPath()
	if err != nil {
		return err
	}

	virtualMachinePaths, err := os.ReadDir(virtualMachinesPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	virtualMachines := map[string]*VirtualMachine{}
	for _, vmPath := range virtualMachinePaths {
		if !vmPath.IsDir() {
			continue
		}

		virtualMachines[vmPath.Name()] = &VirtualMachine{
			Name: vmPath.Name(),
		}
	}

	eng.virtualMachines = virtualMachines

	return nil
}

func (eng *Engine) reloadImages() error {
	images := map[string]*Image{
		"ubuntu:hirsute": {
			Name:        "ubuntu",
			Version:     "hirsute",
			Description: "Ubuntu Server 21.04 (Hirsute Hippo)",
			engine:      eng,
			arch: map[string]*archData{
				"arm64": {
					kernel: &urlAndHash{
						url:      "http://cloud-images.ubuntu.com/hirsute/current/unpacked/hirsute-server-cloudimg-arm64-vmlinuz-generic",
						checksum: "6a101c5a63d472057ab4ae86c57485801eb025a8deb230b488054e23733bd90e",
					},
					initialRamDisk: &urlAndHash{
						url:      "http://cloud-images.ubuntu.com/hirsute/current/unpacked/hirsute-server-cloudimg-arm64-initrd-generic",
						checksum: "39b7807f99b2892dd4c376880b8e953da7286155e6a40938f3d5387b0dbdc39d",
					},
					disk: &urlAndHash{
						url:      "http://cloud-images.ubuntu.com/hirsute/current/hirsute-server-cloudimg-arm64.img",
						checksum: "d58892cd801cb6434b8b5899c1960766e0b9d456a01f7b7d5104de35617ff0f7",
					},
				},
				"amd64": {
					kernel: &urlAndHash{
						url:      "http://cloud-images.ubuntu.com/hirsute/current/unpacked/hirsute-server-cloudimg-amd64-vmlinuz-generic",
						checksum: "",
					},
					initialRamDisk: &urlAndHash{
						url:      "http://cloud-images.ubuntu.com/hirsute/current/unpacked/hirsute-server-cloudimg-amd64-initrd-generic",
						checksum: "",
					},
					disk: &urlAndHash{
						url:      "http://cloud-images.ubuntu.com/hirsute/current/hirsute-server-cloudimg-amd64.img",
						checksum: "",
					},
				},
			},
		},
		"ubuntu:focal": {
			Name:        "ubuntu",
			Version:     "focal",
			Description: "Ubuntu Server 20.04 LTS",
			engine:      eng,
			arch: map[string]*archData{
				"arm64": {
					kernel: &urlAndHash{
						url:      "http://cloud-images.ubuntu.com/focal/current/unpacked/focal-server-cloudimg-arm64-vmlinuz-generic",
						checksum: "",
					},
					initialRamDisk: &urlAndHash{
						url:      "http://cloud-images.ubuntu.com/focal/current/unpacked/focal-server-cloudimg-arm64-initrd-generic",
						checksum: "",
					},
					disk: &urlAndHash{
						url:      "http://cloud-images.ubuntu.com/focal/current/focal-server-cloudimg-arm64.img",
						checksum: "",
					},
				},
				"amd64": {
					kernel: &urlAndHash{
						url:      "http://cloud-images.ubuntu.com/focal/current/unpacked/focal-server-cloudimg-amd64-vmlinuz-generic",
						checksum: "",
					},
					initialRamDisk: &urlAndHash{
						url:      "http://cloud-images.ubuntu.com/focal/current/unpacked/focal-server-cloudimg-amd64-initrd-generic",
						checksum: "",
					},
					disk: &urlAndHash{
						url:      "http://cloud-images.ubuntu.com/focal/current/focal-server-cloudimg-amd64.img",
						checksum: "",
					},
				},
			},
		},
	}

	eng.images = images

	return nil
}

func New(opts *NewOptions) (*Engine, error) {
	engine := &Engine{
		writer: opts.Writer,
		driver: opts.Driver,
	}

	if err := engine.reloadImages(); err != nil {
		return nil, err
	}

	if err := engine.reloadVirtualMachines(); err != nil {
		return nil, err
	}

	return engine, nil
}
