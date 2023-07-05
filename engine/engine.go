package engine

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
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
	CPU      int
	Image    string
	Memory   int
	Name     string
	DiskSize int
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

	fileChecksum := sha512.Sum512(fileBytes)
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
		"debian:bookworm": {
			Archs: map[string]ImageArch{
				"arm64": {
					URL:      "https://cloud.debian.org/images/cloud/bookworm/20230612-1409/debian-12-generic-arm64-20230612-1409.qcow2",
					Checksum: "a55a6e507c4f1d0dcadb3db056515bea210f0163f3257d7eea94a31d096f708bcfc2db89fb7fd4571b3209aab897550ab10a46505025ed66d2cafe4458407e29",
				},
				"amd64": {
					URL:      "https://cloud.debian.org/images/cloud/bookworm/20230612-1409/debian-12-generic-amd64-20230612-1409.qcow2",
					Checksum: "61358292dbec302446a272d5011019091ca78e3fe8878b2d67d31b32e0661306c53a72f793f109394daf937a3db7b2db34422d504e07fdbb300a7bf87109fcf1",
				},
			},
			engine:  e,
			path:    e.imagePath("debian-bookworm"),
			sshUser: "debian",

			Description: "Debian 12 (Bookworm)",
			Name:        "debian",
			Version:     "bookworm",
		},
		"debian:bullseye": {
			Archs: map[string]ImageArch{
				"arm64": {
					URL:      "https://cloud.debian.org/images/cloud/bullseye/20230601-1398/debian-11-generic-arm64-20230601-1398.qcow2",
					Checksum: "8ae9cd402c612359832e92c3ffdb357472b0b0a1e7e25926d6eb326aabaf77e62e97a5809109a194c5cbbf4952f5dca040c8eb83d3f54e60810cd1e964290dd1",
				},
				"amd64": {
					URL:      "https://cloud.debian.org/images/cloud/bullseye/20230601-1398/debian-11-generic-amd64-20230601-1398.qcow2",
					Checksum: "9c052590934349dc207b03f77eeef1f32dca77cfedb1e1cbd6f689eaf507fe997104ac132f38e154e0d6d5f8020e1b90953d50fe14a10864b6c4e773fcae2372",
				},
			},
			engine:  e,
			path:    e.imagePath("debian-bullseye"),
			sshUser: "debian",

			Description: "Debian 11 (Bullseye)",
			Name:        "debian",
			Version:     "bullseye",
		},
		"debian:buster": {
			Archs: map[string]ImageArch{
				"arm64": {
					URL: "https://cloud.debian.org/images/cloud/buster/latest/debian-10-generic-arm64.qcow2",
				},
				"amd64": {
					URL: "https://cloud.debian.org/images/cloud/buster/latest/debian-10-generic-amd64.qcow2",
				},
			},
			engine:  e,
			path:    e.imagePath("debian-buster"),
			sshUser: "debian",

			Description: "Debian 10 (Buster)",
			Name:        "debian",
			Version:     "buster",
		},
		"ubuntu:jammy": {
			Archs: map[string]ImageArch{
				"arm64": {
					URL: "http://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-arm64.img",
				},
				"amd64": {
					URL: "http://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64.img",
				},
			},
			engine:  e,
			path:    e.imagePath("ubuntu-jammy"),
			sshUser: "ubuntu",

			Description: "Ubuntu Server 22.04 (Jammy Jellyfish)",
			Name:        "ubuntu",
			Version:     "jammy",
		},
		"ubuntu:impish": {
			Archs: map[string]ImageArch{
				"arm64": {
					URL: "http://cloud-images.ubuntu.com/impish/current/impish-server-cloudimg-arm64.img",
				},
				"amd64": {
					URL: "http://cloud-images.ubuntu.com/impish/current/impish-server-cloudimg-amd64.img",
				},
			},
			engine:  e,
			path:    e.imagePath("ubuntu-impish"),
			sshUser: "ubuntu",

			Description: "Ubuntu Server 21.10 (Impish Indri)",
			Name:        "ubuntu",
			Version:     "impish",
		},
		"ubuntu:hirsute": {
			Archs: map[string]ImageArch{
				"arm64": {
					URL: "http://cloud-images.ubuntu.com/hirsute/current/hirsute-server-cloudimg-arm64.img",
				},
				"amd64": {
					URL: "http://cloud-images.ubuntu.com/hirsute/current/hirsute-server-cloudimg-amd64.img",
				},
			},
			engine:  e,
			path:    e.imagePath("ubuntu-hirsute"),
			sshUser: "ubuntu",

			Description: "Ubuntu Server 21.04 (Hirsute Hippo)",
			Name:        "ubuntu",
			Version:     "hirsute",
		},
		"ubuntu:focal": {
			Archs: map[string]ImageArch{
				"arm64": {
					URL: "http://cloud-images.ubuntu.com/focal/current/focal-server-cloudimg-arm64.img",
				},
				"amd64": {
					URL: "http://cloud-images.ubuntu.com/focal/current/focal-server-cloudimg-amd64.img",
				},
			},
			engine:  e,
			path:    e.imagePath("ubuntu-focal"),
			sshUser: "ubuntu",

			Description: "Ubuntu Server 20.04 (Focal Fossa) LTS",
			Name:        "ubuntu",
			Version:     "focal",
		},
	}

	keys := make([]string, 0, len(images))
	for k := range images {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sortedImages := make(map[string]*Image, len(images))
	for _, k := range keys {
		sortedImages[k] = images[k]
	}

	e.images = sortedImages

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
	engine.qemu, err = qemu.New(qemu.NewOptions{
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
