package engine

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

type InstanceState uint

func (s InstanceState) String() string {
	return [...]string{"Unknown", "Stopped", "Running"}[s]
}

const (
	InstanceStateUnknown InstanceState = iota
	InstanceStateStopped
	InstanceStateRunning
)

// Instance represents a virtual machine instance in the engine.
type Instance struct {
	VCPU    uint     `json:"cpus"`
	Memory  uint64   `json:"memory"`
	Name    string   `json:"name"`
	Volumes []string `json:"volumes"`

	dir    string
	engine *Engine
}

// Start starts the instance.
func (i *Instance) Start() error {
	i.engine.log("Starting instance %s...", i.Name)

	i.engine.log("Instance %s started", i.Name)

	return nil
}

// State returns the state of the instance.
func (i *Instance) State() InstanceState {
	return InstanceStateStopped
}

// Delete deletes the instance.
func (i *Instance) Delete() error {
	i.engine.log("Deleting instance %s...", i.Name)

	for _, vol := range i.Volumes {
		v, err := i.engine.GetVolume(vol)
		if err != nil {
			return err
		}

		v.forceDelete = true

		if err := v.Delete(); err != nil {
			return err
		}
	}

	if err := os.RemoveAll(i.dir); err != nil {
		return fmt.Errorf("error deleting instance directory: %w", err)
	}

	delete(i.engine.instances, i.Name)

	i.engine.log("Instance %s deleted", i.Name)

	return nil
}

// Write writes the instance configuration to the filesystem.
func (i *Instance) Write() error {
	i.engine.logDebug("Writing instance %s configuration...", i.Name)

	configBytes, err := json.Marshal(i)
	if err != nil {
		return fmt.Errorf("error marshalling instance: %w", err)
	}

	if err := os.WriteFile(filepath.Join(i.dir, "config.json"), configBytes, 0644); err != nil {
		return fmt.Errorf("error writing instance: %w", err)
	}

	i.engine.logDebug("Instance %s configuration written", i.Name)

	return nil
}

// reloadInstances reloads the instances from the filesystem.
func (e *Engine) reloadInstances() error {
	e.logDebug("Reloading instances...")

	e.instances = make(map[string]*Instance)

	files, err := os.ReadDir(e.instanceDir)
	if err != nil {
		return fmt.Errorf("error reading instances directory: %w", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		ist := &Instance{
			dir:    path.Join(e.instanceDir, file.Name()),
			engine: e,
		}

		configBytes, err := os.ReadFile(filepath.Join(ist.dir, "config.json"))
		if err != nil {
			return fmt.Errorf("error reading instance configuration: %w", err)
		}

		if err := json.Unmarshal(configBytes, ist); err != nil {
			return fmt.Errorf("error unmarshalling instance configuration: %w", err)
		}

		for _, vol := range ist.Volumes {
			e.volumes[vol].Instance = ist.Name
		}

		e.instances[ist.Name] = ist
	}

	e.logDebug("Instances reloaded")

	return nil
}

// CreateInstanceOptions represents the options for creating an instance.
type CreateInstanceOptions struct {
	// VCPU is the number of vCPU to allocate to the instance.
	// This field is required.
	VCPU uint
	// Image is the name of the image to use as root volume.
	// This field is required.
	Image string
	// Memory is the amount of memory to allocate to the instance in bytes.
	// This field is required.
	Memory uint64
	// Name is the name of the instance.
	// This field is required.
	Name string
	// UserData is the path to the user data file to use for cloud-init.
	UserData string
	// Volume is the size of the root volume in bytes.
	// This field is required.
	Volume uint64
}

// CreateInstance creates a new instance.
func (e *Engine) CreateInstance(opts *CreateInstanceOptions) (*Instance, error) {
	if opts.VCPU == 0 {
		return nil, fmt.Errorf("%w: vcpus", ErrRequiredFieldNotSet)
	}

	if opts.Image == "" {
		return nil, fmt.Errorf("%w: image", ErrRequiredFieldNotSet)
	}

	if opts.Memory == 0 {
		return nil, fmt.Errorf("%w: memory", ErrRequiredFieldNotSet)
	}

	if opts.Name == "" {
		return nil, fmt.Errorf("%w: name", ErrRequiredFieldNotSet)
	}

	if opts.Volume == 0 {
		return nil, fmt.Errorf("%w: volume", ErrRequiredFieldNotSet)
	}

	if e.instances[opts.Name] != nil {
		return nil, fmt.Errorf("%w: %s", ErrInstanceAlreadyExist, opts.Name)
	}

	e.log("Creating instance %s...", opts.Name)

	vol, err := e.CreateVolume(&CreateVolumeOptions{
		Image: opts.Image,
		Name:  opts.Name,
		Size:  opts.Volume,
	})
	if err != nil {
		return nil, err
	}

	ins := &Instance{
		VCPU:   opts.VCPU,
		Memory: opts.Memory,
		Name:   opts.Name,

		dir:    path.Join(e.instanceDir, opts.Name),
		engine: e,
	}

	if err := os.Mkdir(ins.dir, 0755); err != nil {
		return nil, fmt.Errorf("error creating instance directory: %w", err)
	}

	if err := ins.Write(); err != nil {
		return nil, err
	}

	e.instances[ins.Name] = ins

	if err := vol.Attach(ins.Name); err != nil {
		return nil, err
	}

	e.log("Instance %s created", ins.Name)

	return ins, nil
}

// GetInstance returns the instance by name.
func (e *Engine) GetInstance(name string) (*Instance, error) {
	ins, ok := e.instances[name]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrInstanceNotFound, name)
	}

	return ins, nil
}

// ListInstances returns the instances.
func (e *Engine) ListInstances() map[string]*Instance {
	return e.instances
}
