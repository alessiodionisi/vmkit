package qemu

import (
	"fmt"
	"io"
	"os/exec"
)

// QEMUImg is a wrapper around the qemu-img command.
type QEMUImg struct {
	debug      bool
	executable string
	writer     io.Writer
}

// Create creates a new qcow2 file with the given size in bytes.
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

// Resize resizes the given file to the given size in bytes.
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

// NewQEMUImg creates a new QEMUImg instance.
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
