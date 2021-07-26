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
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/adnsio/vmkit/pkg/cloudinit"
	"golang.org/x/crypto/ssh"
)

// FindVirtualMachine returns the virtual machine if exist
func (eng *Engine) FindVirtualMachine(name string) *VirtualMachine {
	virtualMachine, exist := eng.virtualMachines[name]
	if !exist {
		return nil
	}

	return virtualMachine
}

// ListVirtualMachines returns a slice of loaded virtual machines
func (eng *Engine) ListVirtualMachines() []*VirtualMachine {
	virtualMachines := make([]*VirtualMachine, 0, len(eng.virtualMachines))
	for _, vm := range eng.virtualMachines {
		virtualMachines = append(virtualMachines, vm)
	}

	return virtualMachines
}

// CreateVirtualMachine creates and start a new virtual machine
func (eng *Engine) CreateVirtualMachine(opts *CreateVirtualMachineOptions) (*VirtualMachine, error) {
	// try to find a virtual machine with the same name
	virtualMachineCheck := eng.FindVirtualMachine(opts.Name)
	if virtualMachineCheck != nil {
		return nil, ErrVirtualMachineAlreadyExist
	}

	fmt.Fprintf(eng.writer, "Creating virtual machine \"%s\" with image \"%s\"\n", opts.Name, opts.Image)

	// get the virtual machine path
	virtualMachinePath := eng.virtualMachinePath(opts.Name)

	// create the virtual machine struct
	virtualMachine := &VirtualMachine{
		engine: eng,
		path:   virtualMachinePath,
		config: &vmConfig{
			CPU:    opts.CPU,
			Memory: opts.Memory,
		},

		Name: opts.Name,
	}

	// try to find the image
	image := eng.FindImage(opts.Image)
	if image == nil {
		return nil, ErrImageNotFound
	}

	// check that the image is pulled
	imagePulled, err := image.Pulled()
	if err != nil {
		return nil, err
	}

	if !imagePulled {
		fmt.Fprintf(eng.writer, "Unable to find image \"%s\" locally\n", opts.Image)

		// pull the image
		if err := image.Pull(); err != nil {
			return nil, err
		}
	}

	fmt.Fprint(eng.writer, "Generating a new SSH key\n")

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

	// write the private key to disk
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

	// write the public key to disk
	if err := virtualMachine.writeFile(virtualMachine.publicKeyPath(), publicKeyBytes); err != nil {
		return nil, err
	}

	fmt.Fprint(eng.writer, "Creating cloud-init ISO\n")

	// create cloud init data
	cloudInitUserData := fmt.Sprintf("#cloud-config\nssh_authorized_keys:\n  - %s", string(publicKeyBytes))
	cloudInitNetworkConfig := `version: 2
ethernets:
  interface0:
    match:
      name: enp*
    dhcp4: true
    dhcp6: true
`
	cloudInitMetaData := fmt.Sprintf("instance-id: %s\nlocal-hostname: %s\n", opts.Name, opts.Name)

	// create cloud init iso
	if err := cloudinit.NewCloudInitISO(&cloudinit.NewCloudInitISOOptions{
		MetaData:      cloudInitMetaData,
		Name:          virtualMachine.cloudInitPath(),
		NetworkConfig: cloudInitNetworkConfig,
		UserData:      cloudInitUserData,
	}); err != nil {
		return nil, err
	}

	fmt.Fprint(eng.writer, "Copying disk from image\n")

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

	fmt.Fprint(eng.writer, "Resizing disk\n")

	// TODO: get disk size from opts
	resizeCmd := exec.Command("qemu-img", "resize", virtualMachine.diskPath(), "10G")

	if err := resizeCmd.Run(); err != nil {
		return nil, err
	}

	// write the config file to disk
	if err := virtualMachine.writeConfigFile(); err != nil {
		return nil, err
	}

	return virtualMachine, nil
}
