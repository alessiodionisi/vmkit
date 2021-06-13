package driver

import (
	"encoding/base64"
	"fmt"
	"os/exec"
	"strings"

	"github.com/adnsio/vmkit/pkg/config"
	"gopkg.in/yaml.v2"
)

const (
	AVFVMExecutableName = "avfvm"
)

type AVFVM struct {
	executableName string
}

func (d *AVFVM) Command(config *config.VirtualMachineV1Alpha1) (*exec.Cmd, error) {
	linuxCommandLine := []string{
		"console=hvc0",
	}

	if config.Spec.BootLoader.Linux.CommandLine != "" {
		linuxCommandLine = append(linuxCommandLine, config.Spec.BootLoader.Linux.CommandLine)
	}

	if config.Spec.CloudInit != nil {
		// CLOUD INIT - DATA SOURCE
		linuxCommandLine = append(linuxCommandLine, fmt.Sprintf("ds=nocloud;h=%s;i=%s", config.Metadata.Name, config.Metadata.Name))

		// CLOUD INIT - NETWORK CONFIGURATION
		if config.Spec.CloudInit.NetworkConfiguration != nil {
			cloudInitNetworkConfig, err := yaml.Marshal(config.Spec.CloudInit.NetworkConfiguration)
			if err != nil {
				return nil, err
			}

			linuxCommandLine = append(linuxCommandLine, fmt.Sprintf("network-config=%s", base64.StdEncoding.EncodeToString(cloudInitNetworkConfig)))
		}

		// CLOUD INIT - USER DATA
		if config.Spec.CloudInit.UserData != nil {
			cloudInitUserData, err := yaml.Marshal(config.Spec.CloudInit.UserData)
			if err != nil {
				return nil, err
			}

			linuxCommandLine = append(linuxCommandLine, fmt.Sprintf("cc: %s end_cc", string(cloudInitUserData)))
		}
	}

	return exec.Command(
		d.executableName,
		"--linux-kernel", config.Spec.BootLoader.Linux.Kernel,
		"--linux-initial-ramdisk", config.Spec.BootLoader.Linux.InitialRamdisk,
		"--linux-command-line", strings.Join(linuxCommandLine, " "),
		"--disk-image", config.Spec.Disks[0].Path,
		"--network", fmt.Sprintf("nat,macAddress=%s", config.Spec.Networks[0].MACAddress),
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
