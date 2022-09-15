package server

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/adnsio/vmkit/api"
	"github.com/adnsio/vmkit/api/v1alpha1"
	"github.com/adnsio/vmkit/engine"
	"github.com/adnsio/vmkit/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/yaml.v2"
)

type protoServer struct {
	proto.UnimplementedVMKitServer
	engine *engine.Engine
}

// func (s *protoServer) Create(req *proto.CreateRequest, str proto.VMKit_CreateServer) error {
// 	stepChan := make(chan engine.CreateVirtualMachineStep, 1)

// 	go func() {
// 		for {
// 			step, ok := <-stepChan
// 			if !ok {
// 				log.Debug().Msg("channel closed")
// 				break
// 			}

// 			switch step {
// 			case engine.CreateVirtualMachineStep_DOWNLOADING_IMAGE:
// 				str.Send(&proto.CreateResponse{
// 					Step: proto.CreateResponse_DOWNLOADING_IMAGE,
// 				})
// 			case engine.CreateVirtualMachineStep_RESIZING_DISK:
// 				str.Send(&proto.CreateResponse{
// 					Step: proto.CreateResponse_RESIZING_DISK,
// 				})
// 			case engine.CreateVirtualMachineStep_STARTING:
// 				str.Send(&proto.CreateResponse{
// 					Step: proto.CreateResponse_STARTING,
// 				})
// 			}
// 		}
// 	}()

// 	if err := s.engine.CreateVirtualMachine(stepChan, req.Name, req.Image); err != nil {
// 		return status.Error(codes.Internal, err.Error())
// 	}

// 	return nil
// }

// func (s *protoServer) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
// 	vms := make([]*proto.ListResponse_VirtualMachine, 0, len(s.engine.VirtualMachines))

// 	for _, vm := range s.engine.VirtualMachines {
// 		vms = append(vms, &proto.ListResponse_VirtualMachine{
// 			Name:   vm.Name,
// 			Status: vm.Status.String(),
// 		})
// 	}

// 	return &proto.ListResponse{
// 		VirtualMachines: vms,
// 	}, nil
// }

// func (s *protoServer) ListImages(ctx context.Context, req *proto.ListImagesRequest) (*proto.ListImagesResponse, error) {
// 	imgs := make([]*proto.ListImagesResponse_Image, 0, len(s.engine.Images))

// 	for _, img := range s.engine.Images {
// 		archs := []string{}
// 		for _, src := range img.Sources {
// 			archs = append(archs, src.Arch)
// 		}

// 		imgs = append(imgs, &proto.ListImagesResponse_Image{
// 			Name:        img.Name,
// 			Description: img.Description,
// 			Archs:       archs,
// 		})
// 	}

// 	return &proto.ListImagesResponse{
// 		Images: imgs,
// 	}, nil
// }

func (s *protoServer) Apply(req *proto.ApplyRequest, str proto.VMKit_ApplyServer) error {
	log.Debug().Msgf("rpc: apply")

	dataParts := bytes.Split(req.Data, []byte("---\n"))

	for _, dataPart := range dataParts {
		var parsedResource api.APIVersionAndKind
		if err := yaml.Unmarshal(dataPart, &parsedResource); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		log.Debug().Msgf("parsed resource: %+v", parsedResource)

		nAPIVersion := api.NormalizeAPIVersion(parsedResource.APIVersion)
		nKind := api.NormalizeKind(parsedResource.Kind)

		switch nKind {
		case api.KindDisk:
			var parsedDisk v1alpha1.Disk
			if err := yaml.Unmarshal(dataPart, &parsedDisk); err != nil {
				return status.Error(codes.Internal, err.Error())
			}

			log.Debug().Msgf("parsed %s: %+v", parsedDisk.Kind, parsedDisk)

			if err := s.engine.CreateDisk(parsedDisk); err != nil {
				return status.Error(codes.Internal, err.Error())
			}

			if err := str.Send(&proto.ApplyResponse{
				Message: strings.ToLower(fmt.Sprintf("%s/%s scheduled for creation", nKind, parsedDisk.Metadata.Name)),
			}); err != nil {
				return status.Error(codes.Internal, err.Error())
			}
		case api.KindVirtualMachine:
			var parsedVirtualMachine v1alpha1.VirtualMachine
			if err := yaml.Unmarshal(dataPart, &parsedVirtualMachine); err != nil {
				return status.Error(codes.Internal, err.Error())
			}

			log.Debug().Msgf("parsed %s: %+v", parsedVirtualMachine.Kind, parsedVirtualMachine)

			if err := s.engine.CreateVirtualMachine(parsedVirtualMachine); err != nil {
				return status.Error(codes.Internal, err.Error())
			}

			if err := str.Send(&proto.ApplyResponse{
				Message: strings.ToLower(fmt.Sprintf("%s/%s scheduled for creation", nKind, parsedVirtualMachine.Metadata.Name)),
			}); err != nil {
				return status.Error(codes.Internal, err.Error())
			}
		default:
			if err := str.Send(&proto.ApplyResponse{
				Message: strings.ToLower(fmt.Sprintf("unknow %s kind version %s", nKind, nAPIVersion)),
			}); err != nil {
				return status.Error(codes.Internal, err.Error())
			}
		}
	}

	return nil
}

// func (s *protoServer) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
// 	nKind := api.NormalizeKind(req.Kind)

// 	switch nKind {
// 	case api.KindVirtualMachine:
// 		res := &proto.GetResponse{
// 			Headers: []string{
// 				"Name",
// 				"Status",
// 			},
// 		}

// 		for _, vm := range s.engine.VirtualMachines {
// 			res.Rows = append(res.Rows, strings.Join([]string{
// 				vm.Metadata.Name,
// 				vm.Status.String(),
// 			}, ","))
// 		}

// 		return res, nil
// 	case api.KindDisk:
// 		res := &proto.GetResponse{
// 			Headers: []string{
// 				"Name",
// 				"Status",
// 			},
// 		}

// 		for _, disk := range s.engine.Disks {
// 			res.Rows = append(res.Rows, strings.Join([]string{
// 				disk.Metadata.Name,
// 				disk.Status.String(),
// 			}, ","))
// 		}

// 		return res, nil
// 	default:
// 		return nil, status.Errorf(codes.NotFound, "unknow kind %s", nKind)
// 	}
// }
