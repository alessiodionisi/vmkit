package driver

import (
	"errors"
	"os/exec"

	"github.com/adnsio/vmkit/pkg/config"
)

var (
	ErrNotSupported = errors.New("driver not supported")
)

type DriverType string

const (
	DriverTypeAVFVM DriverType = "avfvm"
	DriverTypeQEMU  DriverType = "qemu"
)

type Driver interface {
	Command(config *config.VirtualMachineV1Alpha1) (*exec.Cmd, error)
}
