package engine

import "github.com/adnsio/vmkit/api/v1alpha1"

type VirtualMachineStatus int

func (s VirtualMachineStatus) String() string {
	switch s {
	case VirtualMachineStatus_CREATING:
		return "Creating"
	default:
		return "Unknow"
	}
}

const (
	VirtualMachineStatus_CREATING VirtualMachineStatus = iota
	VirtualMachineStatus_STOPPED
	VirtualMachineStatus_RUNNING
	VirtualMachineStatus_ERROR
)

type VirtualMachine struct {
	v1alpha1.VirtualMachine
	Status VirtualMachineStatus
}

// type VirtualMachine struct {
// 	Name   string
// 	Status VirtualMachineStatus

// 	dir string
// }

// type VirtualMachineStatus int

// func (s VirtualMachineStatus) String() string {
// 	switch s {
// 	case VirtualMachineStatus_STOPPED:
// 		return "Stopped"
// 	case VirtualMachineStatus_RUNNING:
// 		return "Running"
// 	case VirtualMachineStatus_ERROR:
// 		return "Error"
// 	default:
// 		return "Unknow"
// 	}
// }

// const (
// 	VirtualMachineStatus_STOPPED VirtualMachineStatus = iota
// 	VirtualMachineStatus_RUNNING
// 	VirtualMachineStatus_ERROR
// )

// type CreateVirtualMachineStep int

// const (
// 	CreateVirtualMachineStep_DOWNLOADING_IMAGE CreateVirtualMachineStep = iota
// 	CreateVirtualMachineStep_RESIZING_DISK
// 	CreateVirtualMachineStep_STARTING
// )

// func (e *Engine) CreateVirtualMachine(stepChan chan CreateVirtualMachineStep, name string, image string) error {
// 	defer close(stepChan)

// 	_, ok := e.VirtualMachines[name]
// 	if ok {
// 		return ErrVirtualMachineAlreadyExist
// 	}

// 	img, ok := e.Images[image]
// 	if !ok {
// 		return ErrImageNotFound
// 	}

// 	if !img.IsDownloaded() {
// 		stepChan <- CreateVirtualMachineStep_DOWNLOADING_IMAGE

// 		progressFunc := func(totalBytes int64, downloadedBytes int64) {
// 		}

// 		if err := img.Download(progressFunc); err != nil {
// 			return err
// 		}
// 	}

// 	vm := &VirtualMachine{
// 		Name: name,
// 	}

// 	e.VirtualMachines[vm.Name] = vm

// 	stepChan <- CreateVirtualMachineStep_RESIZING_DISK

// 	stepChan <- CreateVirtualMachineStep_STARTING

// 	return nil
// }

// func (e *Engine) reloadVirtualMachines() error {
// 	e.VirtualMachines = map[string]*VirtualMachine{}

// 	return nil
// }
