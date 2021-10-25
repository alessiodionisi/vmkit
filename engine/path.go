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
