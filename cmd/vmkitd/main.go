package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// type vmKitServer struct {
// 	proto.UnimplementedVMKitServer
// 	configDir string
// 	diskDir   string
// }

// func (s *vmKitServer) Apply(req *proto.ApplyRequest, stream proto.VMKit_ApplyServer) error {
// 	dataParts := bytes.Split(req.Data, []byte("---\n"))

// 	for _, dataPart := range dataParts {
// 		var parsedResource api.APIVersionAndKind
// 		if err := yaml.Unmarshal(dataPart, &parsedResource); err != nil {
// 			return err
// 		}

// 		log.Debug().Msgf("parsed resource: %+v", parsedResource)

// 		switch parsedResource.Kind {

// 		case api.KindDisk:
// 			var parsedDisk v1beta1.Disk
// 			if err := yaml.Unmarshal(dataPart, &parsedDisk); err != nil {
// 				return err
// 			}

// 			log.Debug().Msgf("parsed %s: %+v", api.KindDisk, parsedDisk)

// 			fullDiskDir := path.Join(s.diskDir, fmt.Sprintf("%s.%s", parsedDisk.Metadata.Name, parsedDisk.Spec.Format))

// 			qemuImgArgs := []string{
// 				"create",
// 			}

// 			qemuImgArgs = append(qemuImgArgs, "-f")
// 			qemuImgArgs = append(qemuImgArgs, parsedDisk.Spec.Format)
// 			qemuImgArgs = append(qemuImgArgs, fullDiskDir)
// 			qemuImgArgs = append(qemuImgArgs, parsedDisk.Spec.Size)

// 			qemuImgCmd := exec.Command("qemu-img", qemuImgArgs...)

// 			log.Debug().Msgf("running command: %s", qemuImgCmd.String())

// 			if err := qemuImgCmd.Run(); err != nil {
// 				return err
// 			}

// 			if err := stream.Send(&proto.ApplyReply{
// 				Message: strings.ToLower(fmt.Sprintf("%s/%s created", parsedDisk.Kind, parsedDisk.Metadata.Name)),
// 			}); err != nil {
// 				return err
// 			}

// 		case api.KindVirtualMachine:
// 			var parsedVirtualMachine v1beta1.VirtualMachine
// 			if err := yaml.Unmarshal(dataPart, &parsedVirtualMachine); err != nil {
// 				return err
// 			}

// 			log.Debug().Msgf("parsed %s: %+v", api.KindVirtualMachine, parsedVirtualMachine)

// 			qemuArgs := []string{
// 				"-nodefaults",
// 				"-vga", "none",
// 				"-nographic",
// 				"-cpu", "host",
// 				"-machine", "virt",
// 				"-accel", "hvf",
// 				"-name", fmt.Sprintf("vmkit-%s", parsedVirtualMachine.Metadata.Name),
// 				"-smp", fmt.Sprint(parsedVirtualMachine.Spec.CPU),
// 				"-m", parsedVirtualMachine.Spec.Memory,

// 				"-device", "virtio-serial-pci,id=chardev0",
// 				"-chardev", "stdio,id=chardev0",

// 				"-device", "qemu-xhci",

// 				"-device", "virtio-blk-pci,drive=drive0,bootindex=1",
// 				"-drive", "if=none,media=disk,id=drive0,file=debian-11-generic-arm64.qcow2,discard=unmap,detect-zeroes=unmap",

// 				"-device", "virtio-rng-pci",
// 			}

// 			qemuImgCmd := exec.Command("qemu-system-aarch64", qemuArgs...)

// 			log.Debug().Msgf("running command: %s", qemuImgCmd.String())

// 			if err := stream.Send(&proto.ApplyReply{
// 				Message: strings.ToLower(fmt.Sprintf("%s/%s created", parsedVirtualMachine.Kind, parsedVirtualMachine.Metadata.Name)),
// 			}); err != nil {
// 				return err
// 			}
// 		}

// 		time.Sleep(1 * time.Second)
// 	}

// 	return nil
// }

// func MkdirIfNotExist(name string) error {
// 	_, err := os.Stat(name)
// 	if err != nil {
// 		if !errors.Is(err, os.ErrNotExist) {
// 			return err
// 		}

// 		if err := os.Mkdir(name, os.ModePerm); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

func handleCmdError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	rootCmd := newRootCmd()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	// if err := runMain(); err != nil {
	// 	panic(err)
	// }
}

// func runMain() error {
// 	userHomeDir, err := os.UserHomeDir()
// 	if err != nil {
// 		return err
// 	}

// 	configDir := path.Join(userHomeDir, ".vmkit")
// 	diskDir := path.Join(configDir, "disk")

// 	lis, err := net.Listen("tcp", "[::1]:8000")
// 	if err != nil {
// 		return err
// 	}
// 	defer lis.Close()

// 	if err = MkdirIfNotExist(configDir); err != nil {
// 		return err
// 	}

// 	if err = MkdirIfNotExist(diskDir); err != nil {
// 		return err
// 	}

// 	srv := &vmKitServer{
// 		configDir: configDir,
// 		diskDir:   diskDir,
// 	}

// 	grpcServer := grpc.NewServer()
// 	defer grpcServer.GracefulStop()

// 	proto.RegisterVMKitServer(grpcServer, srv)

// 	return grpcServer.Serve(lis)
// }
