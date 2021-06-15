package engine

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"

	"github.com/adnsio/vmkit/pkg/config"
	"github.com/adnsio/vmkit/pkg/driver"
)

var (
	ErrVirtualMachineAlreadyRunning       = errors.New(`virtual machine is already running`)
	ErrVirtualMachineInvalidConfiguration = errors.New(`virtual machine configuration is invalid`)
	ErrVirtualMachineIsStopped            = errors.New(`virtual machine is stopped`)
	ErrVirtualMachineNotFound             = errors.New(`virtual machine not found`)
)

type VirtualMachinesStatus string

const (
	VirtualMachinesStatusStopped VirtualMachinesStatus = "stopped"
	VirtualMachinesStatusRunning VirtualMachinesStatus = "running"
)

type VirtualMachine struct {
	Config       *config.VirtualMachineV1Alpha1
	command      *exec.Cmd
	errorCh      chan error
	stdout       *bytes.Buffer
	readedStdout []byte
}

func (v *VirtualMachine) Status() VirtualMachinesStatus {
	if v.command == nil {
		return VirtualMachinesStatusStopped
	}

	if v.command.ProcessState != nil && v.command.ProcessState.Exited() {
		return VirtualMachinesStatusStopped
	}

	return VirtualMachinesStatusRunning
}

func (v *VirtualMachine) Logs() (string, error) {
	if v.Status() == VirtualMachinesStatusStopped {
		return "", ErrVirtualMachineIsStopped
	}

	if v.command == nil {
		return "", nil
	}

	stdoutBytes, err := io.ReadAll(v.stdout)
	if err != nil {
		return "", err
	}

	if len(stdoutBytes) == 0 {
		return string(v.readedStdout), nil
	}

	v.readedStdout = stdoutBytes

	return string(v.readedStdout), nil
}

type CreateVirtualMachineOptions struct {
	Config *config.VirtualMachineV1Alpha1
}

type StartVirtualMachineOptions struct {
	Name string
}

type Engine struct {
	virtualMachines []*VirtualMachine
	driver          driver.Driver
}

func (e *Engine) StartVirtualMachine(opts *StartVirtualMachineOptions) error {
	vm, err := e.FindVirtualMachine(opts.Name)
	if err != nil {
		return err
	}

	if vm.Status() == VirtualMachinesStatusRunning {
		return ErrVirtualMachineAlreadyRunning
	}

	cmd, err := e.driver.Command(vm.Config)
	if err != nil {
		return err
	}

	vm.stdout = &bytes.Buffer{}
	vm.readedStdout = []byte{}

	cmd.Stderr = vm.stdout
	cmd.Stdout = vm.stdout

	vm.errorCh = make(chan error, 1)

	go func() {
		if err := cmd.Start(); err != nil {
			vm.errorCh <- err
		}

		if err := cmd.Wait(); err != nil {
			vm.errorCh <- err
		}
	}()

	vm.command = cmd

	return nil
}

func (e *Engine) FindVirtualMachine(name string) (*VirtualMachine, error) {
	for _, vm := range e.virtualMachines {
		if vm.Config.Metadata.Name == name {
			return vm, nil
		}
	}

	return nil, ErrVirtualMachineNotFound
}

func (e *Engine) ReloadVirtualMachines() error {
	vmsPath, err := virtualMachinesPath()
	if err != nil {
		return err
	}

	vmFiles, err := os.ReadDir(vmsPath)
	if err != nil {
		return err
	}

	vms := make([]*VirtualMachine, len(vmFiles))

	for i, vmFile := range vmFiles {
		if vmFile.IsDir() {
			continue
		}

		fileBytes, err := os.ReadFile(path.Join(vmsPath, vmFile.Name()))
		if err != nil {
			return err
		}

		vmCfg, err := config.Unmarshal(fileBytes)
		if err != nil {
			return err
		}

		vms[i] = &VirtualMachine{
			Config: vmCfg,
		}
	}

	e.virtualMachines = vms

	return nil
}

func (e *Engine) VirtualMachinesList() []*VirtualMachine {
	return e.virtualMachines
}

func NewEngine(driver driver.Driver) (*Engine, error) {
	return &Engine{
		driver: driver,
	}, nil
}

func configurationPath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.vmkit", homePath), nil
}

func virtualMachinesPath() (string, error) {
	configPath, err := configurationPath()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/virtual-machine", configPath), nil
}
