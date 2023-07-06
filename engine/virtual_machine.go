package engine

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"syscall"

	"github.com/alessiodionisi/vmkit/cloudinit"
	"github.com/alessiodionisi/vmkit/qemu"
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

type VirtualMachineConfig struct {
	CPU          int
	Memory       int
	DiskSize     int
	Image        string
	SSHUser      string
	MacAddress   string
	PortForwards map[string]string
}

type VirtualMachine struct {
	Name   string
	Config VirtualMachineConfig

	engine *Engine
	path   string
}

func (v *VirtualMachine) makePath() error {
	return os.MkdirAll(v.path, 0755)
}

func (v *VirtualMachine) writeFile(name string, bytes []byte) error {
	if err := v.makePath(); err != nil {
		return err
	}

	return os.WriteFile(name, bytes, 0644)
}

func (v *VirtualMachine) writeFilePerm(name string, bytes []byte, perm os.FileMode) error {
	if err := v.makePath(); err != nil {
		return err
	}

	return os.WriteFile(name, bytes, perm)
}

func (v *VirtualMachine) pidPath() string {
	return path.Join(v.path, "pid")
}

func (v *VirtualMachine) diskPath() string {
	return path.Join(v.path, "disk.qcow2")
}

func (v *VirtualMachine) cloudInitPath() string {
	return path.Join(v.path, "cloud-init.iso")
}

func (v *VirtualMachine) configPath() string {
	return path.Join(v.path, "config.json")
}

func (v *VirtualMachine) privateKeyPath() string {
	return path.Join(v.path, "key.pem")
}

func (v *VirtualMachine) publicKeyPath() string {
	return path.Join(v.path, "key.pub")
}

// Status returns the status of the virtual machine
func (v *VirtualMachine) Status() (VirtualMachineStatus, error) {
	proc, err := v.findProcess()
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
func (v *VirtualMachine) Start() error {
	v.engine.Printf("Starting virtual machine \"%s\"\n", v.Name)

	// get the status
	status, err := v.Status()
	if err != nil {
		return err
	}

	// check the status
	if status == VirtualMachineStatusRunning {
		return ErrVirtualMachineAlreadyRunning
	}

	// use the driver to create the start command
	cmd, err := v.engine.qemu.Command(qemu.CommandOptions{
		CPU:        v.Config.CPU,
		Memory:     v.Config.Memory,
		MACAddress: v.Config.MacAddress,
		Disks: []qemu.CommandOptionsDisk{
			{
				Path: v.diskPath(),
			},
			{
				Path:     v.cloudInitPath(),
				ReadOnly: true,
			},
		},
		PortForwards: v.Config.PortForwards,
	})
	if err != nil {
		return err
	}

	v.engine.Printf("Running command: %s\n", strings.Join(cmd.Args, " "))

	// start the command
	if err := cmd.Start(); err != nil {
		return err
	}

	// write the process id to disk
	if err := v.writeFile(v.pidPath(), []byte(strconv.Itoa(cmd.Process.Pid))); err != nil {
		return err
	}

	return nil
}

// writeConfigFile writes the config file
func (v *VirtualMachine) writeConfigFile() error {
	configBytes, err := json.Marshal(v.Config)
	if err != nil {
		return err
	}

	return v.writeFile(v.configPath(), configBytes)
}

// loadConfigFile loads the config file
func (v *VirtualMachine) loadConfigFile() error {
	configBytes, err := os.ReadFile(v.configPath())
	if err != nil {
		return err
	}

	return json.Unmarshal(configBytes, &v.Config)
}

// findProcess returns the running process if exist
func (v *VirtualMachine) findProcess() (*os.Process, error) {
	pidFileBytes, err := os.ReadFile(v.pidPath())
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
func (v *VirtualMachine) Stop() error {
	status, err := v.Status()
	if err != nil {
		return err
	}

	if status != VirtualMachineStatusRunning {
		return ErrVirtualMachineNotRunning
	}

	v.engine.Printf("Stopping virtual machine \"%s\"\n", v.Name)

	proc, err := v.findProcess()
	if err != nil {
		return err
	}

	if proc == nil {
		return ErrVirtualMachineNotRunning
	}

	if err := proc.Kill(); err != nil {
		return err
	}

	if err := os.Remove(v.pidPath()); err != nil {
		return err
	}

	return nil
}

// Remove stops and remove the virtual machine
func (v *VirtualMachine) Remove() error {
	status, err := v.Status()
	if err != nil {
		return err
	}

	if status != VirtualMachineStatusStopped {
		if err := v.Stop(); err != nil {
			return err
		}
	}

	v.engine.Printf("Removing virtual machine \"%s\"\n", v.Name)

	if err := os.RemoveAll(v.path); err != nil {
		return err
	}

	return nil
}

func (v *VirtualMachine) SSHSessionWithXterm() error {
	status, err := v.Status()
	if err != nil {
		return err
	}

	if status != VirtualMachineStatusRunning {
		return ErrVirtualMachineNotRunning
	}

	v.engine.Printf("Connecting to virtual machine \"%s\" via SSH\n", v.Name)

	sshPort, exist := v.Config.PortForwards["22"]
	if !exist || sshPort == "" {
		return ErrInvalidSSHPort
	}

	privateKeyBytes, err := os.ReadFile(v.privateKeyPath())
	if err != nil {
		return err
	}

	sshSigner, err := ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		return err
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("localhost:%s", sshPort), &ssh.ClientConfig{
		User: v.Config.SSHUser,
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

func (v *VirtualMachine) Exec(cmd string) error {
	status, err := v.Status()
	if err != nil {
		return err
	}

	if status != VirtualMachineStatusRunning {
		return ErrVirtualMachineNotRunning
	}

	sshPort, exist := v.Config.PortForwards["22"]
	if !exist || sshPort == "" {
		return ErrInvalidSSHPort
	}

	privateKeyBytes, err := os.ReadFile(v.privateKeyPath())
	if err != nil {
		return err
	}

	sshSigner, err := ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		return err
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("localhost:%s", sshPort), &ssh.ClientConfig{
		User: v.Config.SSHUser,
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

func (v *VirtualMachine) SSHConnectionDetails() (*SSHConnectionDetails, error) {
	status, err := v.Status()
	if err != nil {
		return nil, err
	}

	if status != VirtualMachineStatusRunning {
		return nil, ErrVirtualMachineNotRunning
	}

	sshPort, exist := v.Config.PortForwards["22"]
	if !exist || sshPort == "" {
		return nil, ErrInvalidSSHPort
	}

	sshPortI, err := strconv.Atoi(sshPort)
	if err != nil {
		return nil, ErrInvalidSSHPort
	}

	return &SSHConnectionDetails{
		Host:       "localhost",
		Port:       sshPortI,
		Username:   v.Config.SSHUser,
		PrivateKey: v.privateKeyPath(),
	}, nil
}

// FindVirtualMachine returns the virtual machine if exist
func (e *Engine) FindVirtualMachine(name string) *VirtualMachine {
	virtualMachine, exist := e.virtualMachines[name]
	if !exist {
		return nil
	}

	return virtualMachine
}

// ListVirtualMachines returns a slice of loaded virtual machines
func (e *Engine) ListVirtualMachines() []*VirtualMachine {
	virtualMachines := make([]*VirtualMachine, 0, len(e.virtualMachines))
	for _, vm := range e.virtualMachines {
		virtualMachines = append(virtualMachines, vm)
	}

	return virtualMachines
}

// CreateVirtualMachine creates and start a new virtual machine
func (e *Engine) CreateVirtualMachine(opts CreateVirtualMachineOptions) (*VirtualMachine, error) {
	// try to find a virtual machine with the same name
	virtualMachineCheck := e.FindVirtualMachine(opts.Name)
	if virtualMachineCheck != nil {
		return nil, ErrVirtualMachineAlreadyExist
	}

	// try to find the image
	image := e.FindImage(opts.Image)
	if image == nil {
		return nil, ErrImageNotFound
	}

	// check that the image is pulled
	imagePulled, err := image.Pulled()
	if err != nil {
		return nil, err
	}

	if !imagePulled {
		e.Printf("Unable to find image \"%s\" locally\n", opts.Image)

		// pull the image
		if err := image.Pull(); err != nil {
			return nil, err
		}
	}

	e.Printf("Creating virtual machine \"%s\" with image \"%s\"\n", opts.Name, opts.Image)

	macAddress, err := e.RandomLocallyAdministeredMacAddress()
	if err != nil {
		return nil, err
	}

	e.Printf("Using %s as MAC address\n", macAddress)

	sshPort, exist := opts.PortForwards["22"]
	if !exist || sshPort == "" {
		// start a tcp listener to find an unused port
		listener, err := net.Listen("tcp", ":0")
		if err != nil {
			return nil, err
		}

		// get the assigned port
		sshPortI := listener.Addr().(*net.TCPAddr).Port

		// stop the tcp lister
		if err := listener.Close(); err != nil {
			return nil, err
		}

		opts.PortForwards["22"] = strconv.Itoa(sshPortI)
	}

	e.Printf("Using %s as SSH port forward\n", opts.PortForwards["22"])

	// get the virtual machine path
	virtualMachinePath := e.virtualMachinePath(opts.Name)

	// create the virtual machine struct
	virtualMachine := &VirtualMachine{
		Name:   opts.Name,
		engine: e,
		path:   virtualMachinePath,
		Config: VirtualMachineConfig{
			CPU:          opts.CPU,
			Memory:       opts.Memory,
			Image:        opts.Image,
			SSHUser:      image.sshUser,
			DiskSize:     opts.DiskSize,
			MacAddress:   macAddress,
			PortForwards: opts.PortForwards,
		},
	}

	e.Printf("Generating a new SSH key\n")

	// generate a new rsa key for ssh
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, err
	}

	// encode the key to pem format
	privateKeyPEMBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	// write the private key
	if err := virtualMachine.writeFilePerm(virtualMachine.privateKeyPath(), privateKeyPEMBytes, 0600); err != nil {
		return nil, err
	}

	// create the public key for ssh
	sshPublicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, err
	}

	// marshal the data for ssh authorized keys
	publicKeyBytes := ssh.MarshalAuthorizedKey(sshPublicKey)

	// write the public key
	if err := virtualMachine.writeFile(virtualMachine.publicKeyPath(), publicKeyBytes); err != nil {
		return nil, err
	}

	e.Printf("Creating cloud-init ISO\n")

	// create cloud init data
	cloudInitUserData := fmt.Sprintf("#cloud-config\n\npassword: password\nchpasswd: { expire: False }\nssh_pwauth: True\nssh_authorized_keys:\n  - %s", string(publicKeyBytes))
	cloudInitNetworkConfig := `version: 2
ethernets:
  interface0:
    match:
      name: enp*
    dhcp4: true
    dhcp6: true
`
	cloudInitMetaData := fmt.Sprintf("instance-id: %s\nlocal-hostname: %s\n", virtualMachine.Name, virtualMachine.Name)

	// create cloud init iso
	if err := cloudinit.New(cloudinit.NewOptions{
		MetaData:      cloudInitMetaData,
		Name:          virtualMachine.cloudInitPath(),
		NetworkConfig: cloudInitNetworkConfig,
		UserData:      cloudInitUserData,
	}); err != nil {
		return nil, err
	}

	e.Printf("Copying disk from image\n")

	// copy the disk from the image
	imageDisk, err := os.Open(image.diskPath())
	if err != nil {
		return nil, err
	}
	defer imageDisk.Close()

	virtualMachineDisk, err := os.Create(virtualMachine.diskPath())
	if err != nil {
		return nil, err
	}
	defer virtualMachineDisk.Close()

	if _, err := io.Copy(virtualMachineDisk, imageDisk); err != nil {
		return nil, err
	}

	e.Printf("Resizing disk\n")

	// resize the disk
	resizeCmd := exec.Command("qemu-img", "resize", virtualMachine.diskPath(), fmt.Sprintf("%dG", virtualMachine.Config.DiskSize))

	if err := resizeCmd.Run(); err != nil {
		return nil, err
	}

	// write the config file
	if err := virtualMachine.writeConfigFile(); err != nil {
		return nil, err
	}

	return virtualMachine, nil
}
