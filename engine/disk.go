package engine

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Disk represents a disk in the engine.
type Disk struct {
	Name string `json:"name"`
	Size uint64 `json:"size"`

	// config     *v1beta1.Disk
	engine     *Engine
	file       string
	configFile string
}

// canApply checks if a disk configuration can be applied.
// func (d *Disk) canApply(config *v1beta1.Disk) error {
// 	changelog, err := diff.Diff(d.config, config)
// 	if err != nil {
// 		return fmt.Errorf("error diffing disk: %w", err)
// 	}

// 	sourceChangelog := changelog.Filter([]string{"Spec", "Source"})
// 	if len(sourceChangelog) > 0 {
// 		return errors.New("cannot change Spec.Source after creation")
// 	}

// 	return nil
// }

// apply applies a disk configuration.
// func (d *Disk) apply(config *v1beta1.Disk) error {
// 	if err := d.canApply(config); err != nil {
// 		return err
// 	}

// 	changelog, err := diff.Diff(d.config, config)
// 	if err != nil {
// 		return fmt.Errorf("error diffing disk: %w", err)
// 	}

// 	for _, change := range changelog {
// 		switch change.Path[0] {
// 		case "Spec":
// 			switch change.Path[1] {
// 			case "Size":
// 				d.Resize(config.Spec.Size)
// 			}
// 		}
// 	}

// 	return nil
// }

// Resize resizes the disk and updates the disk configuration.
func (d *Disk) Resize(size uint64) error {
	d.engine.logger.Debug("resizing disk", "disk", d.Name, "from", d.Size, "to", size)

	resizeCmd := d.engine.qemuImgCmd.Resize(d.file, d.Size)
	if err := resizeCmd.Run(); err != nil {
		return fmt.Errorf("error resizing disk: %w", err)
	}

	d.Size = size

	d.engine.logger.Debug("disk resized", "disk", d.Name, "from", d.Size, "to", size)

	return d.Write()
}

// Write writes the disk configuration to the filesystem.
func (d *Disk) Write() error {
	d.engine.logger.Debug("writing disk configuration", "disk", d.Name)

	configBytes, err := json.Marshal(d)
	if err != nil {
		return fmt.Errorf("error marshalling disk: %w", err)
	}

	if err := os.WriteFile(d.configFile, configBytes, 0644); err != nil {
		return fmt.Errorf("error writing disk: %w", err)
	}

	d.engine.logger.Debug("disk configuration written", "disk", d.Name)

	return nil
}

// createDisk it's used by Engine ApplyDisk function to create a disk.
// Don't call this function directly, use ApplyDisk instead.
// func (e *Engine) createDisk(parsedDisk *v1beta1.Disk) error {
// 	_, exist := e.disks[parsedDisk.Metadata.Name]
// 	if exist {
// 		return fmt.Errorf("error creating disk: disk %s already exist", parsedDisk.Metadata.Name)
// 	}

// 	e.logger.Debug("creating disk", "disk", parsedDisk.Metadata.Name)

// 	diskDir := path.Join(e.disksDir, parsedDisk.Metadata.Name)

// 	d := &Disk{
// 		Disk:       parsedDisk,
// 		engine:     e,
// 		file:       path.Join(diskDir, "disk.qcow2"),
// 		configFile: path.Join(diskDir, "config.yaml"),
// 	}

// 	if err := os.Mkdir(diskDir, 0755); err != nil {
// 		return fmt.Errorf("error creating disk directory: %w", err)
// 	}

// 	e.disks[d.Metadata.Name] = d

// 	if err := d.Write(); err != nil {
// 		return err
// 	}

// 	createCmd := e.qemuImgCmd.Create(d.file, d.Spec.Size)
// 	if err := createCmd.Run(); err != nil {
// 		return fmt.Errorf("error creating disk: %w", err)
// 	}

// 	e.logger.Debug("disk created", "disk", d.Metadata.Name)

// 	return nil
// }

// CanApplyDisk checks if a disk configuration can be applied.
// func (e *Engine) CanApplyDisk(config *v1beta1.Disk) (exist bool, err error) {
// 	disk, exist := e.disks[config.Metadata.Name]
// 	if !exist {
// 		return false, nil
// 	}

// 	if err := disk.canApply(config); err != nil {
// 		return true, err
// 	}

// 	return true, nil
// }

// ApplyDisk applies a disk configuration.
// func (e *Engine) ApplyDisk(config *v1beta1.Disk) (created bool, err error) {
// 	disk, exist := e.disks[config.Metadata.Name]
// 	if exist {
// 		if err := disk.canApply(config); err != nil {
// 			return false, err
// 		}
// 	}

// 	diskDir := path.Join(e.disksDir, config.Metadata.Name)

// 	disk = &Disk{
// 		config:     config,
// 		engine:     e,
// 		file:       path.Join(diskDir, "disk.qcow2"),
// 		configFile: path.Join(diskDir, "config.yaml"),
// 		status: status{
// 			phase: statusPhaseUnknown,
// 		},
// 	}

// 	e.disks[config.Metadata.Name] = disk

// 	return true, nil

// 	// e.logger.Debug("creating disk", "disk", disk.config.Metadata.Name)

// 	// if err := os.Mkdir(diskDir, 0755); err != nil {
// 	// 	disk.status.phase = statusPhaseFailed

// 	// 	return fmt.Errorf("error creating disk directory: %w", err)
// 	// }

// 	// if err := disk.Write(); err != nil {
// 	// 	disk.status.phase = statusPhaseFailed

// 	// 	return err
// 	// }

// 	// disk.status.phase = statusPhaseCreating

// 	// createCmd := e.qemuImgCmd.Create(disk.file, disk.config.Spec.Size)
// 	// if err := createCmd.Run(); err != nil {
// 	// 	disk.status.phase = statusPhaseFailed

// 	// 	return fmt.Errorf("error creating disk: %w", err)
// 	// }

// 	// disk.status.phase = statusPhaseAvailable

// 	// e.logger.Debug("disk created", "disk", disk.config.Metadata.Name)

// 	// return nil
// }

func (e *Engine) CreateDisk(name string, size uint64) (*Disk, error) {
	disk := &Disk{
		Name: name,
		Size: size,

		engine:     e,
		file:       filepath.Join(e.disksDir, name, "disk.qcow2"),
		configFile: filepath.Join(e.disksDir, name, "config.json"),
	}

	return disk, nil
}
