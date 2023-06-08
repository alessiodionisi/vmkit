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

	"github.com/alessiodionisi/vmkit/qemu"
	"github.com/cheggaaa/pb/v3"
)

type NewOptions struct {
	QEMUExecutableName string
	Path               string
	Writer             io.Writer
}

type CreateVirtualMachineOptions struct {
	CPU    int
	Image  string
	Memory int
	Name   string
}

type Engine struct {
	qemu            *qemu.QEMU
	images          map[string]*Image
	path            string
	virtualMachines map[string]*VirtualMachine
	writer          io.Writer
}

func (e *Engine) validateChecksum(checksum string, name string) error {
	e.Printf("Validating checksum of %s\n", name)

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
		return fmt.Errorf(`%w. value: "%s", expected: "%s"`, ErrInvalidChecksum, hexFileChecksum, checksum)
	}

	return nil
}

func (e *Engine) downloadAndPrintProgress(url string, name string) error {
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

	e.Printf("Downloading %s\n", url)

	// download the content
	progressBar := pb.Full.New(contentLength)
	progressBar.SetWriter(e.writer)

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

func (e *Engine) reloadVirtualMachines() error {
	virtualMachinesPath := e.virtualMachinesPath()

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

			engine: e,
			path:   e.virtualMachinePath(vmPath.Name()),
			Config: &VirtualMachineConfig{},
		}

		if err := virtualMachines[vmPath.Name()].loadConfigFile(); err != nil {
			return err
		}
	}

	e.virtualMachines = virtualMachines

	return nil
}

func (e *Engine) reloadImages() error {
	images := map[string]*Image{
		"debian:bullseye": {
			arch: map[string]*imageArchitecture{
				"arm64": {
					disk: &imageURLAndChecksum{
						url: "https://cloud.debian.org/images/cloud/bullseye/latest/debian-11-generic-arm64.qcow2",
					},
				},
				"amd64": {
					disk: &imageURLAndChecksum{
						url: "https://cloud.debian.org/images/cloud/bullseye/latest/debian-11-generic-amd64.qcow2",
					},
				},
			},
			engine: e,
			path:   e.imagePath("debian-bullseye"),

			Description: "Debian 11 (Bullseye)",
			Name:        "debian",
			Version:     "bullseye",
		},
		"debian:buster": {
			arch: map[string]*imageArchitecture{
				"arm64": {
					disk: &imageURLAndChecksum{
						url: "https://cloud.debian.org/images/cloud/buster/latest/debian-10-generic-arm64.qcow2",
					},
				},
				"amd64": {
					disk: &imageURLAndChecksum{
						url: "https://cloud.debian.org/images/cloud/buster/latest/debian-10-generic-amd64.qcow2",
					},
				},
			},
			engine: e,
			path:   e.imagePath("debian-buster"),

			Description: "Debian 10 (Buster)",
			Name:        "debian",
			Version:     "buster",
		},
		"ubuntu:jammy": {
			arch: map[string]*imageArchitecture{
				"arm64": {
					disk: &imageURLAndChecksum{
						url: "http://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-arm64.img",
					},
				},
				"amd64": {
					disk: &imageURLAndChecksum{
						url: "http://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64.img",
					},
				},
			},
			engine: e,
			path:   e.imagePath("ubuntu-jammy"),

			Description: "Ubuntu Server 22.04 (Jammy Jellyfish)",
			Name:        "ubuntu",
			Version:     "jammy",
		},
		"ubuntu:impish": {
			arch: map[string]*imageArchitecture{
				"arm64": {
					disk: &imageURLAndChecksum{
						url: "http://cloud-images.ubuntu.com/impish/current/impish-server-cloudimg-arm64.img",
					},
				},
				"amd64": {
					disk: &imageURLAndChecksum{
						url: "http://cloud-images.ubuntu.com/impish/current/impish-server-cloudimg-amd64.img",
					},
				},
			},
			engine: e,
			path:   e.imagePath("ubuntu-impish"),

			Description: "Ubuntu Server 21.10 (Impish Indri)",
			Name:        "ubuntu",
			Version:     "impish",
		},
		"ubuntu:hirsute": {
			arch: map[string]*imageArchitecture{
				"arm64": {
					disk: &imageURLAndChecksum{
						url: "http://cloud-images.ubuntu.com/hirsute/current/hirsute-server-cloudimg-arm64.img",
					},
				},
				"amd64": {
					disk: &imageURLAndChecksum{
						url: "http://cloud-images.ubuntu.com/hirsute/current/hirsute-server-cloudimg-amd64.img",
					},
				},
			},
			engine: e,
			path:   e.imagePath("ubuntu-hirsute"),

			Description: "Ubuntu Server 21.04 (Hirsute Hippo)",
			Name:        "ubuntu",
			Version:     "hirsute",
		},
		"ubuntu:focal": {
			arch: map[string]*imageArchitecture{
				"arm64": {
					disk: &imageURLAndChecksum{
						url: "http://cloud-images.ubuntu.com/focal/current/focal-server-cloudimg-arm64.img",
					},
				},
				"amd64": {
					disk: &imageURLAndChecksum{
						url: "http://cloud-images.ubuntu.com/focal/current/focal-server-cloudimg-amd64.img",
					},
				},
			},
			engine: e,
			path:   e.imagePath("ubuntu-focal"),

			Description: "Ubuntu Server 20.04 (Focal Fossa) LTS",
			Name:        "ubuntu",
			Version:     "focal",
		},
	}

	e.images = images

	return nil
}

func (e *Engine) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(e.writer, format, a...)
}

func New(opts *NewOptions) (*Engine, error) {
	engine := &Engine{
		path:   opts.Path,
		writer: opts.Writer,
	}

	var err error
	engine.qemu, err = qemu.New(&qemu.NewOptions{
		ExecutableName: opts.QEMUExecutableName,
	})
	if err != nil {
		return nil, err
	}

	if err := engine.reloadImages(); err != nil {
		return nil, err
	}

	if err := engine.reloadVirtualMachines(); err != nil {
		return nil, err
	}

	return engine, nil
}
