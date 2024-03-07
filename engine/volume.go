package engine

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"slices"
)

// Volume represents a volume in the engine.
type Volume struct {
	Name     string `json:"name"`
	Size     uint64 `json:"size"`
	Source   string `json:"source"`
	Instance string `json:"-"`

	dir         string
	engine      *Engine
	forceDelete bool
}

// Resize resizes the volume.
func (v *Volume) Resize(size uint64) error {
	v.engine.log("Resizing volume %s...", v.Name)

	// resizeCmd := v.engine.qemuImg.Resize(v.file, v.Size)
	// if err := resizeCmd.Run(); err != nil {
	// 	return fmt.Errorf("error resizing volume: %w", err)
	// }

	// v.Size = size

	// if err := v.Write(); err != nil {
	// 	return err
	// }

	v.engine.log("Volume %s resized", v.Name)

	return nil
}

// Attach attaches the volume to an instance.
func (v *Volume) Attach(instance string) error {
	if v.Instance != "" {
		return fmt.Errorf("%w: %s", ErrVolumeAlreadyAttached, v.Name)
	}

	v.engine.log("Attaching volume %s to instance %s...", v.Name, instance)

	i, err := v.engine.GetInstance(instance)
	if err != nil {
		return err
	}

	v.Instance = i.Name

	i.Volumes = append(i.Volumes, v.Name)

	if err := i.Write(); err != nil {
		return err
	}

	v.engine.log("Volume %s attached to instance %s", v.Name, i.Name)

	return nil
}

// Detach detaches the volume from the instance.
func (v *Volume) Detach() error {
	if v.Instance == "" {
		return fmt.Errorf("%w: %s", ErrVolumeNotAttached, v.Name)
	}

	v.engine.log("Detaching volume %s...", v.Name)

	i, err := v.engine.GetInstance(v.Instance)
	if err != nil {
		return err
	}

	for vi, vol := range i.Volumes {
		if vol == v.Name {
			if vi == 0 {
				return fmt.Errorf("%w: %s", ErrCannotDetachRootVolume, v.Name)
			}

			i.Volumes = slices.Delete(i.Volumes, vi, vi+1)

			if err := i.Write(); err != nil {
				return err
			}

			break
		}
	}

	v.Instance = ""

	v.engine.log("Volume %s detached", v.Name)

	return nil
}

// Write writes the volume configuration to the filesystem.
func (v *Volume) Write() error {
	v.engine.logDebug("Writing volume %s configuration...", v.Name)

	configBytes, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("error marshalling volume: %w", err)
	}

	if err := os.WriteFile(filepath.Join(v.dir, "config.json"), configBytes, 0644); err != nil {
		return fmt.Errorf("error writing volume: %w", err)
	}

	v.engine.logDebug("Volume %s configuration written", v.Name)

	return nil
}

// Delete deletes the volume.
func (v *Volume) Delete() error {
	v.engine.log("Deleting volume %s...", v.Name)

	if v.Instance != "" && !v.forceDelete {
		return fmt.Errorf("%w: %s", ErrVolumeAttached, v.Name)
	}

	if err := os.RemoveAll(v.dir); err != nil {
		return fmt.Errorf("error deleting volume directory: %w", err)
	}

	delete(v.engine.volumes, v.Name)

	v.engine.log("Volume %s deleted", v.Name)

	return nil
}

// reloadVolumes reloads the volumes from the filesystem.
func (e *Engine) reloadVolumes() error {
	e.logDebug("Reloading volumes...")

	e.volumes = make(map[string]*Volume)

	volFiles, err := os.ReadDir(e.volumeDir)
	if err != nil {
		return fmt.Errorf("error reading volumes directory: %w", err)
	}

	for _, f := range volFiles {
		if !f.IsDir() {
			continue
		}

		vol := &Volume{
			dir:    path.Join(e.volumeDir, f.Name()),
			engine: e,
		}

		configBytes, err := os.ReadFile(filepath.Join(vol.dir, "config.json"))
		if err != nil {
			return fmt.Errorf("error reading volume configuration: %w", err)
		}

		if err := json.Unmarshal(configBytes, vol); err != nil {
			return fmt.Errorf("error unmarshalling volume configuration: %w", err)
		}

		e.volumes[vol.Name] = vol
	}

	e.logDebug("Volumes reloaded")

	return nil
}

// CreateVolumeOptions represents the options to create a volume.
type CreateVolumeOptions struct {
	// Image is the name of the image to use as source.
	// Set this field to create a volume from an image.
	Image string
	// Name is the name of the volume.
	// This field is required.
	Name string
	// Size is the size of the volume in bytes.
	// This field is required.
	Size uint64
	// Source is the path to the directory to use as source.
	// Set this field to create a shared volume.
	Source string
}

// CreateVolume creates a new volume.
func (e *Engine) CreateVolume(opts *CreateVolumeOptions) (*Volume, error) {
	if opts.Name == "" {
		return nil, fmt.Errorf("%w: name", ErrRequiredFieldNotSet)
	}

	if opts.Size == 0 {
		return nil, fmt.Errorf("%w: size", ErrRequiredFieldNotSet)
	}

	if e.volumes[opts.Name] != nil {
		return nil, fmt.Errorf("%w: %s", ErrVolumeAlreadyExist, opts.Name)
	}

	if opts.Source != "" && opts.Image != "" {
		return nil, fmt.Errorf("%w: source and image", ErrInvalidFieldCombination)
	}

	var img *Image
	if opts.Image != "" {
		var err error
		img, err = e.GetImage(opts.Image)
		if err != nil {
			return nil, err
		}
	}

	e.log("Creating volume %s...", opts.Name)

	vol := &Volume{
		Name:   opts.Name,
		Size:   opts.Size,
		Source: opts.Source,

		dir:    path.Join(e.volumeDir, opts.Name),
		engine: e,
	}

	if err := os.Mkdir(vol.dir, 0755); err != nil {
		return nil, fmt.Errorf("error creating volume directory: %w", err)
	}

	switch {
	case opts.Source != "":
		panic("not implemented")

	case img != nil:
		break

	default:
		if err := e.qemuImg.Create(filepath.Join(vol.dir, "data.qcow2"), vol.Size); err != nil {
			return nil, fmt.Errorf("error creating volume with qemu-img: %w", err)
		}
	}

	if err := vol.Write(); err != nil {
		return nil, err
	}

	e.volumes[opts.Name] = vol

	e.log("Volume %s created", vol.Name)

	return vol, nil
}

// GetVolume returns a volume by name.
func (e *Engine) GetVolume(name string) (*Volume, error) {
	vol, ok := e.volumes[name]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrVolumeNotFound, name)
	}

	return vol, nil
}

// ListVolumes returns the volumes.
func (e *Engine) ListVolumes() map[string]*Volume {
	return e.volumes
}
