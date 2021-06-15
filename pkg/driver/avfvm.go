package driver

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/adnsio/vmkit/pkg/config"
)

var (
	ErrInvalidBootLoaderConfiguration    = errors.New("invalid boot loader configuration")
	ErrInvalidLinuxBootLoaderCommandLine = errors.New("invalid linux boot loader command line")
)

const (
	AVFVMExecutableName = "avfvm"
)

type AVFVM struct {
	executableName string
}

func (d *AVFVM) Command(config *config.VirtualMachineV1Alpha1) (*exec.Cmd, error) {
	cmdArgs := []string{}

	if config.Spec.BootLoader.Linux == nil {
		return nil, fmt.Errorf("%w, linux boot loader is mandatory with avfvm", ErrInvalidBootLoaderConfiguration)
	}

	cmdArgs = append(cmdArgs, "--cpu-count", fmt.Sprint(config.Spec.CPU))
	cmdArgs = append(cmdArgs, "--linux-kernel", config.Spec.BootLoader.Linux.Kernel)
	cmdArgs = append(cmdArgs, "--linux-initial-ramdisk", config.Spec.BootLoader.Linux.InitialRamdisk)

	var linuxCommandLine []string

	if config.Spec.BootLoader.Linux.CommandLine != "" {
		linuxCommandLine = append(linuxCommandLine, config.Spec.BootLoader.Linux.CommandLine)
	}

	if config.Spec.CloudInit != nil {
		// CLOUD INIT - DATA SOURCE
		linuxCommandLine = append(linuxCommandLine, fmt.Sprintf("ds=nocloud;h=%s;i=%s", config.Metadata.Name, config.Metadata.Name))

		// CLOUD INIT - NETWORK CONFIGURATION
		if config.Spec.CloudInit.NetworkConfiguration != nil {
			linuxCommandLine = append(
				linuxCommandLine,
				fmt.Sprintf("network-config=%s",
					base64.StdEncoding.EncodeToString([]byte(*config.Spec.CloudInit.NetworkConfiguration)),
				),
			)
		}

		// CLOUD INIT - USER DATA
		if config.Spec.CloudInit.UserData != nil {
			linuxCommandLine = append(linuxCommandLine, fmt.Sprintf("cc: %s end_cc", *config.Spec.CloudInit.UserData))
		}
	}

	linuxCommandLineString := strings.Join(linuxCommandLine, " ")

	if !strings.Contains(linuxCommandLineString, "console=hvc0") {
		return nil, fmt.Errorf(
			"%w, console=hvc0 is mandatory with avfvm",
			ErrInvalidLinuxBootLoaderCommandLine,
		)
	}

	if config.Spec.CPU > 1 && !strings.Contains(linuxCommandLineString, "irqaffinity=0") {
		return nil, fmt.Errorf(
			"%w, irqaffinity=0 is needed to fix sync problems with more than one cpu with avfvm",
			ErrInvalidLinuxBootLoaderCommandLine,
		)
	}

	cmdArgs = append(cmdArgs, "--linux-command-line", linuxCommandLineString)

	for _, disk := range config.Spec.Disks {
		cmdArgs = append(cmdArgs, "--disk-image", disk.Path)
	}

	for _, network := range config.Spec.Networks {
		cmdArgs = append(cmdArgs, "--network", fmt.Sprintf("nat,macAddress=%s", network.MACAddress))
	}

	return exec.Command(
		d.executableName,
		cmdArgs...,
	), nil
}

func (d *AVFVM) supported() bool {
	path, err := exec.LookPath(d.executableName)
	if err != nil {
		return false
	}

	if path == "" {
		return false
	}

	return true
}

func NewAVFVM(
	executableName string,
) (Driver, error) {
	d := &AVFVM{
		executableName: executableName,
	}

	if !d.supported() {
		return nil, ErrNotSupported
	}

	return d, nil
}
