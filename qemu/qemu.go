package qemu

import (
	"fmt"
	"os/exec"
)

type QEMUCmd struct {
	executable string
}

func (q *QEMUCmd) Run() *exec.Cmd {
	cmdArgs := []string{}

	return exec.Command(
		q.executable,
		cmdArgs...,
	)
}

func NewQEMUCmd(executable string) (*QEMUCmd, error) {
	executable, err := exec.LookPath(executable)
	if err != nil {
		return nil, fmt.Errorf("error looking for qemu: %w", err)
	}

	return &QEMUCmd{
		executable: executable,
	}, nil
}
