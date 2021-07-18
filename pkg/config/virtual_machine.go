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

package config

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
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
	CPU        uint                                  `yaml:"cpu"`
	Memory     string                                `yaml:"memory"`
	BootLoader *VirtualMachineV1Alpha1SpecBootLoader `yaml:"bootLoader"`
	Disks      []*VirtualMachineV1Alpha1SpecDisk     `yaml:"disks"`
	Networks   []*VirtualMachineV1Alpha1SpecNetwork  `yaml:"networks"`
	CloudInit  *VirtualMachineV1Alpha1SpecCloudInit  `yaml:"cloudInit"`
}

type VirtualMachineV1Alpha1SpecDisk struct {
	Path string `yaml:"path"`
}

type VirtualMachineV1Alpha1SpecNetwork struct {
	MACAddress string `yaml:"macAddress"`
}

type VirtualMachineV1Alpha1SpecBootLoader struct {
	Linux *VirtualMachineV1Alpha1SpecBootLoaderLinux `yaml:"linux"`
	EFI   *VirtualMachineV1Alpha1SpecBootLoaderEFI   `yaml:"efi"`
}

type VirtualMachineV1Alpha1SpecBootLoaderEFI struct {
	Path string `yaml:"path"`
}

type VirtualMachineV1Alpha1SpecBootLoaderLinux struct {
	Kernel         string `yaml:"kernel"`
	InitialRamdisk string `yaml:"initialRamdisk"`
	CommandLine    string `yaml:"commandLine"`
}

type VirtualMachineV1Alpha1SpecCloudInit struct {
	Enabled              bool   `yaml:"enabled"`
	UserData             string `yaml:"userData"`
	NetworkConfiguration string `yaml:"networkConfiguration"`
}

func (c *VirtualMachineV1Alpha1) Validate() error {
	// VALIDATE METADATA
	if c.Metadata == nil {
		return fmt.Errorf("%w, Metadata is required", ErrInvalidMetadataConfiguration)
	}

	if c.Metadata != nil && c.Metadata.Name == "" {
		return fmt.Errorf("%w, Metadata.Name is required", ErrInvalidMetadataConfiguration)
	}

	// VALIDATE SPEC
	if c.Spec == nil {
		return fmt.Errorf("%w, Spec is required", ErrInvalidSpecConfiguration)
	}

	if c.Spec.CPU == 0 {
		return fmt.Errorf("%w, Spec.CPU is required", ErrInvalidSpecConfiguration)
	}

	if c.Spec.Memory == "" {
		return fmt.Errorf("%w, Spec.Memory is required", ErrInvalidSpecConfiguration)
	}

	// VALIDATE SPEC BOOTLOADER
	if c.Spec.BootLoader == nil &&
		c.Spec.BootLoader.EFI == nil &&
		c.Spec.BootLoader.Linux == nil {
		return fmt.Errorf("%w, one of Spec.BootLoader.EFI or Spec.BootLoader.Linux is required", ErrInvalidBootLoaderConfiguration)
	}

	if c.Spec.BootLoader.EFI != nil && c.Spec.BootLoader.EFI.Path == "" {
		return fmt.Errorf("%w, Spec.BootLoader.EFI.Path is required", ErrInvalidBootLoaderConfiguration)
	}

	return nil
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
