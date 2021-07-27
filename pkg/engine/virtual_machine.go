// Spin up Linux VMs with QEMU and Apple virtualization framework
// Copyright (C) 2021 VMKit Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package engine

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"path"
	"strconv"
	"strings"
	"syscall"

	"github.com/adnsio/vmkit/pkg/driver"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

type SSHConnectionDetails struct {
	Host       string
	Port       int
	Username   string
	PrivateKey string
}

type VirtualMachineStatus string

const (
	VirtualMachineStatusStopped VirtualMachineStatus = "stopped"
	VirtualMachineStatusRunning VirtualMachineStatus = "running"
	VirtualMachineStatusError   VirtualMachineStatus = "error"
)

type vmConfig struct {
	CPU    int `json:"cpu"`
	Memory int `json:"memory"`
}

type VirtualMachine struct {
	config *vmConfig
	engine *Engine
	path   string

	Name string
}

func (vm *VirtualMachine) makePath() error {
	return os.MkdirAll(vm.path, 0755)
}

func (vm *VirtualMachine) writeFile(name string, bytes []byte) error {
	if err := vm.makePath(); err != nil {
		return err
	}

	return os.WriteFile(name, bytes, 0644)
}

func (vm *VirtualMachine) writeFilePerm(name string, bytes []byte, perm os.FileMode) error {
	if err := vm.makePath(); err != nil {
		return err
	}

	return os.WriteFile(name, bytes, perm)
}

func (vm *VirtualMachine) pidPath() string {
	return path.Join(vm.path, "pid")
}

func (vm *VirtualMachine) sshPortPath() string {
	return path.Join(vm.path, "ssh-port")
}

func (vm *VirtualMachine) diskPath() string {
	return path.Join(vm.path, "disk.img")
}

func (vm *VirtualMachine) cloudInitPath() string {
	return path.Join(vm.path, "cloud-init.iso")
}

func (vm *VirtualMachine) configPath() string {
	return path.Join(vm.path, "config.json")
}

func (vm *VirtualMachine) privateKeyPath() string {
	return path.Join(vm.path, "key.pem")
}

func (vm *VirtualMachine) publicKeyPath() string {
	return path.Join(vm.path, "key.pub")
}

// Status returns the status of the virtual machine
func (vm *VirtualMachine) Status() (VirtualMachineStatus, error) {
	proc, err := vm.findProcess()
	if err != nil {
		return VirtualMachineStatusStopped, err
	}

	if proc == nil {
		return VirtualMachineStatusStopped, nil
	}

	// check if is running with a signal
	if err := proc.Signal(syscall.Signal(0)); err != nil {
		return VirtualMachineStatusStopped, nil
	}

	return VirtualMachineStatusRunning, nil
}

// Start starts the virtual machine
func (vm *VirtualMachine) Start() error {
	fmt.Fprintf(vm.engine.writer, "Starting virtual machine \"%s\"\n", vm.Name)

	// get the status
	status, err := vm.Status()
	if err != nil {
		return err
	}

	// check the status
	if status == VirtualMachineStatusRunning {
		return ErrVirtualMachineAlreadyRunning
	}

	// start a tcp listener to find an unused port
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return err
	}

	// get the assigned port
	sshPort := listener.Addr().(*net.TCPAddr).Port

	// stop the tcp lister
	if err := listener.Close(); err != nil {
		return err
	}

	if err := vm.engine.checkAndWriteBiosFiles(); err != nil {
		return err
	}

	// use the driver to create the start command
	cmd, err := vm.engine.driver.Command(&driver.CommandOptions{
		Disks: []string{
			vm.diskPath(),
		},
		CloudInitISO:   vm.cloudInitPath(),
		CPU:            vm.config.CPU,
		Memory:         vm.config.Memory,
		SSHPortForward: sshPort,
	})
	if err != nil {
		return err
	}

	// fmt.Fprintf(vm.engine.writer, "Running command: %s\n", strings.Join(cmd.Args, " "))

	// start the command
	if err := cmd.Start(); err != nil {
		return err
	}

	// write the process id to disk
	if err := vm.writeFile(vm.pidPath(), []byte(strconv.Itoa(cmd.Process.Pid))); err != nil {
		return err
	}

	// write the ssh port to disk
	if err := vm.writeFile(vm.sshPortPath(), []byte(strconv.Itoa(sshPort))); err != nil {
		return err
	}

	return nil
}

// writeConfigFile writes the config file
func (vm *VirtualMachine) writeConfigFile() error {
	configBytes, err := json.Marshal(vm.config)
	if err != nil {
		return err
	}

	return vm.writeFile(vm.configPath(), configBytes)
}

// loadConfigFile loads the config file
func (vm *VirtualMachine) loadConfigFile() error {
	configBytes, err := os.ReadFile(vm.configPath())
	if err != nil {
		return err
	}

	return json.Unmarshal(configBytes, vm.config)
}

func (vm *VirtualMachine) sshPort() (int, error) {
	sshPortFileBytes, err := os.ReadFile(vm.sshPortPath())
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return 0, nil
		}

		return 0, err
	}

	sshPort, err := strconv.Atoi(strings.ReplaceAll(string(sshPortFileBytes), "\n", ""))
	if err != nil {
		return 0, err
	}

	return sshPort, nil
}

// findProcess returns the running process if exist
func (vm *VirtualMachine) findProcess() (*os.Process, error) {
	pidFileBytes, err := os.ReadFile(vm.pidPath())
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}

		return nil, err
	}

	pid, err := strconv.Atoi(strings.ReplaceAll(string(pidFileBytes), "\n", ""))
	if err != nil {
		return nil, err
	}

	return os.FindProcess(pid)
}

// Stop stops the virtual machine
func (vm *VirtualMachine) Stop() error {
	status, err := vm.Status()
	if err != nil {
		return err
	}

	if status != VirtualMachineStatusRunning {
		return ErrVirtualMachineNotRunning
	}

	fmt.Fprintf(vm.engine.writer, "Stopping virtual machine \"%s\"\n", vm.Name)

	proc, err := vm.findProcess()
	if err != nil {
		return err
	}

	if proc == nil {
		return ErrVirtualMachineNotRunning
	}

	if err := proc.Kill(); err != nil {
		return err
	}

	if err := os.Remove(vm.pidPath()); err != nil {
		return err
	}

	if err := os.Remove(vm.sshPortPath()); err != nil {
		return err
	}

	return nil
}

// Remove stops and remove the virtual machine
func (vm *VirtualMachine) Remove() error {
	status, err := vm.Status()
	if err != nil {
		return err
	}

	if status != VirtualMachineStatusStopped {
		if err := vm.Stop(); err != nil {
			return err
		}
	}

	fmt.Fprintf(vm.engine.writer, "Removing virtual machine \"%s\"\n", vm.Name)

	if err := os.RemoveAll(vm.path); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) SSHSessionWithXterm() error {
	status, err := vm.Status()
	if err != nil {
		return err
	}

	if status != VirtualMachineStatusRunning {
		return ErrVirtualMachineNotRunning
	}

	fmt.Fprintf(vm.engine.writer, "Connecting to virtual machine \"%s\" via SSH\n", vm.Name)

	sshPort, err := vm.sshPort()
	if err != nil {
		return err
	}

	if sshPort == 0 {
		return ErrInvalidSSHPort
	}

	privateKeyBytes, err := os.ReadFile(vm.privateKeyPath())
	if err != nil {
		return err
	}

	sshSigner, err := ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		return err
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("localhost:%d", sshPort), &ssh.ClientConfig{
		User: "ubuntu",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(sshSigner),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		BannerCallback:  ssh.BannerDisplayStderr(),
	})
	if err != nil {
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	session.Stdout = os.Stdout

	terminalWidth, terminalHeight, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err
	}

	if err := session.RequestPty("xterm-256color", terminalHeight, terminalWidth, ssh.TerminalModes{
		// ssh.ECHO:          0,
		// ssh.TTY_OP_ISPEED: 14400,
		// ssh.TTY_OP_OSPEED: 14400,
	}); err != nil {
		return err
	}

	if err := session.Shell(); err != nil {
		return err
	}

	if err := session.Wait(); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) Exec(cmd string) error {
	status, err := vm.Status()
	if err != nil {
		return err
	}

	if status != VirtualMachineStatusRunning {
		return ErrVirtualMachineNotRunning
	}

	sshPort, err := vm.sshPort()
	if err != nil {
		return err
	}

	if sshPort == 0 {
		return ErrInvalidSSHPort
	}

	privateKeyBytes, err := os.ReadFile(vm.privateKeyPath())
	if err != nil {
		return err
	}

	sshSigner, err := ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		return err
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("localhost:%d", sshPort), &ssh.ClientConfig{
		User: "ubuntu",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(sshSigner),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		BannerCallback:  ssh.BannerDisplayStderr(),
	})
	if err != nil {
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	session.Stderr = os.Stderr
	session.Stdout = os.Stdout

	if err := session.Run(cmd); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) SSHConnectionDetails() (*SSHConnectionDetails, error) {
	status, err := vm.Status()
	if err != nil {
		return nil, err
	}

	if status != VirtualMachineStatusRunning {
		return nil, ErrVirtualMachineNotRunning
	}

	sshPort, err := vm.sshPort()
	if err != nil {
		return nil, err
	}

	if sshPort == 0 {
		return nil, ErrInvalidSSHPort
	}

	return &SSHConnectionDetails{
		Host:       "localhost",
		Port:       sshPort,
		Username:   "ubuntu",
		PrivateKey: vm.privateKeyPath(),
	}, nil
}
