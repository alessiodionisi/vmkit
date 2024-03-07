package engine

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/alessiodionisi/vmkit/qemu"
)

// Engine is the main engine of VMKit.
type Engine struct {
	dataDir     string
	debug       bool
	images      map[string]*Image
	instanceDir string
	instances   map[string]*Instance
	logWriter   io.Writer
	qemuCmd     *qemu.QEMUCmd
	qemuImg     *qemu.QEMUImg
	volumeDir   string
	volumes     map[string]*Volume
}

type NewOptions struct {
	DataDir   string
	Debug     bool
	LogWriter io.Writer
}

// New creates a new Engine.
func New(opts NewOptions) (*Engine, error) {
	volumeDir := path.Join(opts.DataDir, "volume")
	instanceDir := path.Join(opts.DataDir, "instance")

	qemuImgCmd, err := qemu.NewQEMUImg("qemu-img", opts.LogWriter, opts.Debug)
	if err != nil {
		return nil, fmt.Errorf("error creating qemu-img cmd: %w", err)
	}

	qemuCmd, err := qemu.NewQEMUCmd("qemu-system-aarch64")
	if err != nil {
		return nil, fmt.Errorf("error creating qemu cmd: %w", err)
	}

	e := &Engine{
		dataDir:     opts.DataDir,
		debug:       opts.Debug,
		images:      make(map[string]*Image),
		instanceDir: instanceDir,
		instances:   make(map[string]*Instance),
		logWriter:   opts.LogWriter,
		qemuCmd:     qemuCmd,
		qemuImg:     qemuImgCmd,
		volumeDir:   volumeDir,
		volumes:     make(map[string]*Volume),
	}

	if _, err := os.Stat(e.dataDir); errors.Is(err, os.ErrNotExist) {
		e.logDebug(`Creating data directory "%s"...`, e.dataDir)

		if err := os.Mkdir(e.dataDir, 0755); err != nil {
			return nil, fmt.Errorf("error creating data directory: %w", err)
		}
	}

	if _, err := os.Stat(e.volumeDir); errors.Is(err, os.ErrNotExist) {
		e.logDebug(`Creating volume directory "%s"...`, e.volumeDir)

		if err := os.Mkdir(e.volumeDir, 0755); err != nil {
			return nil, fmt.Errorf("error creating volume directory: %w", err)
		}
	}

	if _, err := os.Stat(e.instanceDir); errors.Is(err, os.ErrNotExist) {
		e.logDebug(`Creating instance directory "%s"...`, e.instanceDir)

		if err := os.Mkdir(e.instanceDir, 0755); err != nil {
			return nil, fmt.Errorf("error creating instance directory: %w", err)
		}
	}

	if err := e.reloadImages(); err != nil {
		return nil, fmt.Errorf("error reloading images: %w", err)
	}

	if err := e.reloadVolumes(); err != nil {
		return nil, fmt.Errorf("error reloading volumes: %w", err)
	}

	if err := e.reloadInstances(); err != nil {
		return nil, fmt.Errorf("error reloading instances: %w", err)
	}

	return e, nil
}
