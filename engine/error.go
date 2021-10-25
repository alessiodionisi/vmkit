package engine

import "errors"

var (
	ErrImageNotFound                = errors.New("engine: image not found")
	ErrInvalidChecksum              = errors.New("engine: invalid checksum")
	ErrInvalidSSHPort               = errors.New("engine: invalid ssh port")
	ErrUnsupportedArchitecture      = errors.New("engine: unsupported architecture")
	ErrVirtualMachineAlreadyExist   = errors.New("engine: virtual machine already exist")
	ErrVirtualMachineAlreadyRunning = errors.New("engine: virtual machine is already running")
	ErrVirtualMachineNotRunning     = errors.New("engine: virtual machine is not running")
)
