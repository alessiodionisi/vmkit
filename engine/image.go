package engine

import "fmt"

type Image struct {
	Name string `json:"name"`

	e *Engine
}

// reloadImages reloads the images from the filesystem.
func (e *Engine) reloadImages() error {
	e.logDebug("Reloading images...")

	e.images = map[string]*Image{
		"fedora-39": {
			Name: "fedora-39",
			e:    e,
		},
	}

	e.logDebug("Images reloaded")

	return nil
}

// GetImage returns an image by name.
func (e *Engine) GetImage(name string) (*Image, error) {
	img, ok := e.images[name]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrImageNotFound, name)
	}

	return img, nil
}

// ListImages returns the images.
func (e *Engine) ListImages() map[string]*Image {
	return e.images
}
