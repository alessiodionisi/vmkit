package v1alpha1

import "github.com/adnsio/vmkit/api"

type VirtualMachineSpec struct {
}

type VirtualMachine struct {
	api.APIVersionAndKind
	api.Metadata
	Spec *VirtualMachineSpec
}
