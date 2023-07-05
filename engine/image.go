package engine

import (
	"errors"
	"os"
	"path"
	"runtime"
)

type ImageArch struct {
	Checksum string
	URL      string
}

type Image struct {
	Description string
	Name        string
	Version     string
	Archs       map[string]ImageArch

	engine  *Engine
	path    string
	sshUser string
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

	arch, exist := i.Archs[runtime.GOARCH]
	if !exist {
		return ErrUnsupportedArchitecture
	}

	if err := i.engine.downloadAndPrintProgress(arch.URL, i.diskPath()); err != nil {
		return err
	}

	if arch.Checksum != "" {
		if err := i.engine.validateChecksum(arch.Checksum, i.diskPath()); err != nil {
			return err
		}
	}

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
