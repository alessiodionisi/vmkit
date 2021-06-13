package rpc

import (
	"github.com/adnsio/vmkit/pkg/config"
	"github.com/adnsio/vmkit/pkg/engine"
)

type StartVirtualMachineOptions struct {
	Name string
}

type VirtualMachineLogsOptions struct {
	Name string
}

type VirtualMachine struct {
	Config *config.VirtualMachineV1Alpha1
	Status engine.VirtualMachinesStatus
}

type VirtualMachineReceiver struct {
	engine *engine.Engine
}

func (r *VirtualMachineReceiver) Start(opts *StartVirtualMachineOptions, _ *interface{}) error {
	return r.engine.StartVirtualMachine(&engine.StartVirtualMachineOptions{
		Name: opts.Name,
	})
}

func (r *VirtualMachineReceiver) List(_ *interface{}, reply *[]*VirtualMachine) error {
	virtualMachines := r.engine.VirtualMachinesList()

	var res = make([]*VirtualMachine, len(virtualMachines))
	for i, vm := range virtualMachines {
		res[i] = &VirtualMachine{
			Config: vm.Config,
			Status: vm.Status(),
		}
	}

	*reply = res

	return nil
}

func (r *VirtualMachineReceiver) Logs(opts *VirtualMachineLogsOptions, reply *string) error {
	virtualMachine, err := r.engine.FindVirtualMachine(opts.Name)
	if err != nil {
		return err
	}

	logs, err := virtualMachine.Logs()
	if err != nil {
		return err
	}

	*reply = logs

	return nil
}
