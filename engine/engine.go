package engine

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/adnsio/vmkit/api/v1alpha1"
	"github.com/adnsio/vmkit/util"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type Engine struct {
	// Images          map[string]*Image
	VirtualMachines map[string]*VirtualMachine

	configDir         string
	diskDir           string
	virtualMachineDir string
	// arch              string
}

func (e *Engine) CreateVirtualMachine(apiVM v1alpha1.VirtualMachine) error {
	// TODO: validate apiVM

	vm := &VirtualMachine{
		VirtualMachine: apiVM,
	}

	e.VirtualMachines[vm.Metadata.Name] = vm

	vmYAML, err := yaml.Marshal(vm.VirtualMachine)
	if err != nil {
		return err
	}

	vmYAMLFile := path.Join(e.virtualMachineDir, fmt.Sprintf("%s.yaml", vm.Metadata.Name))
	if err := os.WriteFile(vmYAMLFile, vmYAML, os.ModePerm); err != nil {
		return err
	}

	go func() {
		log.Debug().Msgf("vm %s: creating", vm.Metadata.Name)
		vm.Status = VirtualMachineStatus_CREATING

		if len(vm.Spec.Disks) > 0 {
			log.Debug().Msgf("vm %s: checking disks", vm.Metadata.Name)

			disksOk := 0
			for disksOk != len(vm.Spec.Disks) {
				disksOk = 0

				for _, vmDisk := range vm.Spec.Disks {
					log.Debug().Msgf("vm %s: checking disk %s", vm.Metadata.Name, vmDisk.Name)

					disk, exist := e.Disks[vmDisk.Name]
					if !exist {
						log.Error().Msgf("vm %s: unknow disk %s", vm.Metadata.Name, vmDisk.Name)
						vm.Status = VirtualMachineStatus_ERROR
						return
					}

					if disk.Status == DiskStatus_READY {
						disksOk++
					}
				}

				time.Sleep(2 * time.Second)
			}

			log.Debug().Msgf("vm %s: all disks are ready", vm.Metadata.Name)
		}

		log.Debug().Msgf("vm %s: starting", vm.Metadata.Name)

		qemuArgs := []string{
			"-nodefaults",
			// "-display", "none",
			"-cpu", "host",
			"-machine", "virt",
			"-accel", "hvf",
			"-name", fmt.Sprintf("vmkit-%s", vm.Metadata.Name),
			"-smp", fmt.Sprint(vm.Spec.CPU.CPUs),
			"-m", vm.Spec.Memory.Size,

			"-bios", "edk2-aarch64-code.fd",

			"-device", "qemu-xhci",
			"-device", "virtio-rng-pci",
			"-device", "virtio-gpu-pci",

			// "-device", "virtio-serial-pci,id=chardev0",
			// "-chardev", "stdio,id=chardev0",

			// "-device", "virtio-blk-pci,drive=drive0,bootindex=1",
			// "-drive", "if=none,media=disk,id=drive0,file=debian-11-generic-arm64.qcow2,discard=unmap,detect-zeroes=unmap",
		}

		for i, vmDisk := range vm.Spec.Disks {
			disk, exist := e.Disks[vmDisk.Name]
			if !exist {
				log.Error().Msgf("vm %s: unknow disk %s", vm.Metadata.Name, vmDisk.Name)
				vm.Status = VirtualMachineStatus_ERROR
				return
			}

			qemuArgs = append(qemuArgs, []string{
				"-device", fmt.Sprintf("virtio-blk-pci,drive=drive%d", i),
				"-drive", fmt.Sprintf("if=none,media=disk,id=drive%d,file=%s", i, disk.dir),
			}...)
		}

		qemuCmd := exec.Command("qemu-system-aarch64", qemuArgs...)

		log.Debug().Msgf("vm %s: qemu command: %s", vm.Metadata.Name, qemuCmd.String())

		if err := qemuCmd.Start(); err != nil {
			log.Error().Msgf("vm %s: error: %s", vm.Metadata.Name, err.Error())
			vm.Status = VirtualMachineStatus_ERROR
			return
		}
		defer qemuCmd.Process.Kill()

		vm.Status = VirtualMachineStatus_RUNNING

		if err := qemuCmd.Wait(); err != nil {
			log.Error().Msgf("vm %s: error: %s", vm.Metadata.Name, err.Error())
			vm.Status = VirtualMachineStatus_ERROR
			return
		}

		vm.Status = VirtualMachineStatus_STOPPED
	}()

	return nil
}

func (e *Engine) CreateDisk(apiDisk v1alpha1.Disk) error {
	// TODO: validate apiDisk

	disk := &Disk{
		Disk: apiDisk,
		dir:  path.Join(e.diskDir, fmt.Sprintf("%s.img", apiDisk.Metadata.Name)),
	}

	e.Disks[disk.Metadata.Name] = disk

	diskYAML, err := yaml.Marshal(disk.Disk)
	if err != nil {
		return err
	}

	diskYAMLFile := path.Join(e.diskDir, fmt.Sprintf("%s.yaml", disk.Metadata.Name))
	if err := os.WriteFile(diskYAMLFile, diskYAML, os.ModePerm); err != nil {
		return err
	}

	go func() {
		log.Debug().Msgf("disk %s: creating", disk.Metadata.Name)
		disk.Status = DiskStatus_CREATING

		// resp, err := http.Get(disk.Spec.Source.URL)
		// if err != nil {
		// 	log.Error().Msgf("disk %s: error: %s", disk.Metadata.Name, err.Error())
		// 	return
		// }
		// defer resp.Body.Close()

		// tmpFile := fmt.Sprintf("%s.tmp", disk.dir)
		// diskFile, err := os.OpenFile(tmpFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
		// if err != nil {
		// 	log.Error().Msgf("disk %s: error: %s", disk.Metadata.Name, err.Error())
		// 	return
		// }
		// defer diskFile.Close()

		// log.Debug().Msgf("disk %s: downloading from %s", disk.Metadata.Name, disk.Spec.Source.URL)

		// if _, err := io.Copy(diskFile, resp.Body); err != nil {
		// 	log.Error().Msgf("disk %s: error: %s", disk.Metadata.Name, err.Error())

		// 	if err := os.Remove(tmpFile); err != nil {
		// 		log.Error().Msgf("disk %s: error: %s", disk.Metadata.Name, err.Error())
		// 		return
		// 	}

		// 	return
		// }

		// if err := os.Rename(tmpFile, disk.dir); err != nil {
		// 	log.Error().Msgf("disk %s: error: %s", disk.Metadata.Name, err.Error())
		// 	return
		// }

		disk.Status = DiskStatus_READY
	}()

	return nil
}

func New() (*Engine, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := path.Join(userHomeDir, ".vmkit")

	diskDir := path.Join(configDir, "disk")
	if err := util.MkdirAllIfNotExist(diskDir); err != nil {
		return nil, err
	}

	virtualMachineDir := path.Join(configDir, "virtual-machine")
	if err := util.MkdirAllIfNotExist(virtualMachineDir); err != nil {
		return nil, err
	}

	eng := &Engine{
		configDir:         configDir,
		diskDir:           diskDir,
		virtualMachineDir: virtualMachineDir,
		// arch:              runtime.GOARCH,
	}

	// if err := eng.reloadImages(); err != nil {
	// 	return nil, err
	// }

	// if err := eng.reloadVirtualMachines(); err != nil {
	// 	return nil, err
	// }

	return eng, nil
}
