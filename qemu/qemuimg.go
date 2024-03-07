package qemu

import (
	"fmt"
	"io"
	"os/exec"
)

type QEMUImg struct {
	debug      bool
	executable string
	writer     io.Writer
}

func (q *QEMUImg) Create(file string, size uint64) error {
	cmdArgs := []string{
		"create",
		"-f", "qcow2",
		file,
		fmt.Sprintf("%d", size),
	}

	cmd := exec.Command(
		q.executable,
		cmdArgs...,
	)

	if q.debug {
		cmd.Stdout = q.writer
		cmd.Stderr = q.writer
	}

	return cmd.Run()
}

func (q *QEMUImg) Resize(file string, size uint64) *exec.Cmd {
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

func NewQEMUImg(executable string, writer io.Writer, debug bool) (*QEMUImg, error) {
	executable, err := exec.LookPath(executable)
	if err != nil {
		return nil, fmt.Errorf("error looking for qemu-img: %w", err)
	}

	return &QEMUImg{
		debug:      debug,
		executable: executable,
		writer:     writer,
	}, nil
}
