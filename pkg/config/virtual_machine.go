package config

import (
	"errors"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	ErrInvalidVersion = errors.New("invalid version")
	ErrInvalidKind    = errors.New("invalid kind")
)

type VirtualMachineV1Alpha1 struct {
	APIVersion string                          `yaml:"apiVersion"`
	Kind       string                          `yaml:"kind"`
	Metadata   *VirtualMachineV1Alpha1Metadata `yaml:"metadata"`
	Spec       *VirtualMachineV1Alpha1Spec     `yaml:"spec"`
}

type VirtualMachineV1Alpha1Metadata struct {
	Name string `yaml:"name"`
}

type VirtualMachineV1Alpha1Spec struct {
	CPU        int                                   `yaml:"cpu"`
	Memory     string                                `yaml:"memory"`
	BootLoader *VirtualMachineV1Alpha1SpecBootLoader `yaml:"bootLoader"`
	Disks      []*VirtualMachineV1Alpha1SpecDisk     `yaml:"disks"`
	Networks   []*VirtualMachineV1Alpha1SpecNetwork  `yaml:"networks"`
	CloudInit  *VirtualMachineV1Alpha1SpecCloudInit  `yaml:"cloudInit"`
}

type VirtualMachineV1Alpha1SpecDisk struct {
	Path     string `yaml:"path"`
	ReadOnly bool   `yaml:"readOnly"`
}

type VirtualMachineV1Alpha1SpecNetwork struct {
	Type       string `yaml:"type"`
	MACAddress string `yaml:"macAddress"`
}

type VirtualMachineV1Alpha1SpecBootLoader struct {
	Linux *VirtualMachineV1Alpha1SpecBootLoaderLinux `yaml:"linux"`
}

type VirtualMachineV1Alpha1SpecBootLoaderLinux struct {
	Kernel         string `yaml:"kernel"`
	InitialRamdisk string `yaml:"initialRamdisk"`
	CommandLine    string `yaml:"commandLine"`
}

type VirtualMachineV1Alpha1SpecCloudInit struct {
	UserData             interface{} `yaml:"userData"`
	NetworkConfiguration interface{} `yaml:"networkConfiguration"`
}

func Unmarshal(bytes []byte) (*VirtualMachineV1Alpha1, error) {
	var configData VirtualMachineV1Alpha1
	if err := yaml.Unmarshal(bytes, &configData); err != nil {
		return nil, err
	}

	if strings.ToLower(configData.APIVersion) != "v1alpha1" {
		return nil, ErrInvalidVersion
	}

	if strings.ToLower(configData.Kind) != "virtualmachine" {
		return nil, ErrInvalidKind
	}

	return &configData, nil
}

func NewVirtualMachine() *VirtualMachineV1Alpha1 {
	return &VirtualMachineV1Alpha1{
		APIVersion: "v1alpha1",
		Kind:       "VirtualMachine",
	}
}
