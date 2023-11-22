package qemu

import (
	"fmt"
	"os/exec"
)

type QEMUImgCmd struct {
	executable string
}

func (q *QEMUImgCmd) Create(file string, size uint64) *exec.Cmd {
	cmdArgs := []string{
		"create",
		file,
		fmt.Sprintf("%d", size),
	}

	return exec.Command(
		q.executable,
		cmdArgs...,
	)
}

func (q *QEMUImgCmd) Resize(file string, size uint64) *exec.Cmd {
	cmdArgs := []string{
		"resize",
		file,
		fmt.Sprintf("%d", size),
	}

	return exec.Command(
		q.executable,
		cmdArgs...,
	)
}

func NewQEMUImgCmd(executable string) (*QEMUImgCmd, error) {
	executable, err := exec.LookPath(executable)
	if err != nil {
		return nil, fmt.Errorf("error looking for qemu-img: %w", err)
	}

	return &QEMUImgCmd{
		executable: executable,
	}, nil
}
