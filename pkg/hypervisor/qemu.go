package hypervisor

import (
	"os/exec"
)

type QEMU struct {
}

func (q *QEMU) IsSupported() bool {
	path, err := exec.LookPath(qemuExecutableName)
	if err != nil {
		return false
	}

	if path == "" {
		return false
	}

	return true
}

func NewQEMU() (Hypervisor, error) {
	return &QEMU{}, nil
}
