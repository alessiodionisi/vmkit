package engine

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/alessiodionisi/vmkit/qemu"
)

// Engine is the main engine of VMKit.
type Engine struct {
	logger     *slog.Logger
	vmkitDir   string
	disksDir   string
	disks      map[string]*Disk
	qemuCmd    *qemu.QEMUCmd
	qemuImgCmd *qemu.QEMUImgCmd
}

// Apply applies resource configurations.
// messageChan is used to receive messages about the status of the operation.
// func (e *Engine) Apply(messageChan chan<- string, data []byte) error {
// 	e.logger.Debug("applying resource configurations")

// 	dataParts := bytes.Split(data, []byte("---\n"))
// 	for _, dataPart := range dataParts {
// 		var parsedResource api.APIVersionAndKind
// 		if err := yaml.Unmarshal(dataPart, &parsedResource); err != nil {
// 			return fmt.Errorf("error unmarshalling resource: %w", err)
// 		}

// 		nAPIVersion := api.NormalizeAPIVersion(parsedResource.APIVersion)
// 		nKind := api.NormalizeKind(parsedResource.Kind)

// 		e.logger.Debug("applying resource", "api-version", nAPIVersion, "kind", nKind)

// 		switch nKind {
// 		case api.KindDisk:
// 			var parsedDisk v1beta1.Disk
// 			if err := yaml.Unmarshal(dataPart, &parsedDisk); err != nil {
// 				return fmt.Errorf("error unmarshalling disk: %w", err)
// 			}

// 			created, err := e.ApplyDisk(&parsedDisk)
// 			if err != nil {
// 				messageChan <- err.Error()
// 				continue
// 			}

// 			if created {
// 				messageChan <- fmt.Sprintf("%s/%s created", nKind, nAPIVersion)
// 				continue
// 			}

// 			messageChan <- fmt.Sprintf("%s/%s applied", nKind, nAPIVersion)

// 		default:
// 			messageChan <- fmt.Sprintf("unknow %s kind version %s", nKind, nAPIVersion)
// 		}
// 	}

// 	e.logger.Debug("resource configurations applied")

// 	close(messageChan)

// 	return nil
// }

// New creates a new Engine.
func New(
	logger *slog.Logger,
	vmkitDir string,
) (*Engine, error) {
	disksDir := path.Join(vmkitDir, "disks")
	if _, err := os.Stat(disksDir); errors.Is(err, os.ErrNotExist) {
		logger.Debug("creating disks directory", "path", disksDir)

		if err := os.Mkdir(disksDir, 0755); err != nil {
			return nil, fmt.Errorf("error creating disks directory: %w", err)
		}
	}

	qemuImgCmd, err := qemu.NewQEMUImgCmd("qemu-img")
	if err != nil {
		return nil, fmt.Errorf("error creating qemu-img cmd: %w", err)
	}

	qemuCmd, err := qemu.NewQEMUCmd("qemu-system-aarch64")
	if err != nil {
		return nil, fmt.Errorf("error creating qemu cmd: %w", err)
	}

	return &Engine{
		logger:     logger,
		vmkitDir:   vmkitDir,
		disksDir:   disksDir,
		disks:      make(map[string]*Disk),
		qemuImgCmd: qemuImgCmd,
		qemuCmd:    qemuCmd,
	}, nil
}
