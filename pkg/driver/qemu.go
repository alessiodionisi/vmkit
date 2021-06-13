package driver

import (
	"errors"
	"os/exec"

	"github.com/adnsio/vmkit/pkg/config"
)

type QEMU struct {
	executableName string
}

func (q *QEMU) supported() bool {
	path, err := exec.LookPath(q.executableName)
	if err != nil {
		return false
	}

	if path == "" {
		return false
	}

	return true
}

func (d *QEMU) Command(config *config.VirtualMachineV1Alpha1) (*exec.Cmd, error) {
	return nil, errors.New("TODO")
}

func NewQEMU(
	executableName string,
) (Driver, error) {
	d := &QEMU{
		executableName: executableName,
	}

	if !d.supported() {
		return nil, ErrNotSupported
	}

	return d, nil
}
