package server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/adnsio/vmkit/assets"
	"github.com/adnsio/vmkit/proto"
)

type imageSource struct {
	URL      string
	Checksum string
	Arch     string
}

type imageSSH struct {
	Username string
}

type imageFeatures struct {
	CloudInit bool
}

type image struct {
	Name        string
	Description string
	Sources     []imageSource
	SSH         imageSSH
	Features    imageFeatures
}

type imagesJSON struct {
	Images []image
}

type protoServer struct {
	proto.UnimplementedVMKitServer
}

func (s *protoServer) Create(_ context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	return &proto.CreateResponse{}, nil
}

func (s *protoServer) ListImages(_ context.Context, _ *proto.ListImagesRequest) (*proto.ListImagesResponse, error) {
	var imgsJSON imagesJSON
	if err := json.Unmarshal(assets.ImagesJSON, &imgsJSON); err != nil {
		return nil, fmt.Errorf("json: %w", err)
	}

	res := &proto.ListImagesResponse{
		Images: make([]*proto.ListImagesResponse_Image, len(imgsJSON.Images)),
	}

	for imgI, img := range imgsJSON.Images {
		archs := make([]string, len(img.Sources))
		for srcI, src := range img.Sources {
			archs[srcI] = src.Arch
		}

		res.Images[imgI] = &proto.ListImagesResponse_Image{
			Name:        img.Name,
			Description: img.Description,
			Archs:       archs,
		}
	}

	return res, nil
}
