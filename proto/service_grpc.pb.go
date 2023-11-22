// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	VMKit_DiskCreate_FullMethodName           = "/VMKit/DiskCreate"
	VMKit_DiskDelete_FullMethodName           = "/VMKit/DiskDelete"
	VMKit_DiskResize_FullMethodName           = "/VMKit/DiskResize"
	VMKit_DiskList_FullMethodName             = "/VMKit/DiskList"
	VMKit_DiskAttach_FullMethodName           = "/VMKit/DiskAttach"
	VMKit_VirtualMachineCreate_FullMethodName = "/VMKit/VirtualMachineCreate"
	VMKit_VirtualMachineDelete_FullMethodName = "/VMKit/VirtualMachineDelete"
	VMKit_VirtualMachineStart_FullMethodName  = "/VMKit/VirtualMachineStart"
	VMKit_VirtualMachineStop_FullMethodName   = "/VMKit/VirtualMachineStop"
	VMKit_VirtualMachineList_FullMethodName   = "/VMKit/VirtualMachineList"
	VMKit_ImagePull_FullMethodName            = "/VMKit/ImagePull"
	VMKit_ImageCreate_FullMethodName          = "/VMKit/ImageCreate"
	VMKit_ImageDelete_FullMethodName          = "/VMKit/ImageDelete"
	VMKit_ImageList_FullMethodName            = "/VMKit/ImageList"
)

// VMKitClient is the client API for VMKit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VMKitClient interface {
	// DiskCreate creates a new disk.
	// The disk can be created from an image or from scratch.
	DiskCreate(ctx context.Context, in *DiskCreateRequest, opts ...grpc.CallOption) (*DiskCreateResponse, error)
	// DiskDelete deletes a disk.
	// The disk must be detached from any virtual machine.
	DiskDelete(ctx context.Context, in *DiskDeleteRequest, opts ...grpc.CallOption) (*DiskDeleteResponse, error)
	// DiskResize resizes a disk.
	// The disk must be detached from any virtual machine. The size must be greater than the current size.
	DiskResize(ctx context.Context, in *DiskResizeRequest, opts ...grpc.CallOption) (*DiskResizeResponse, error)
	// DiskList lists all disks.
	DiskList(ctx context.Context, in *DiskListRequest, opts ...grpc.CallOption) (*DiskListResponse, error)
	// DiskAttach attaches a disk to a virtual machine.
	// The virtual machine must be stopped.
	DiskAttach(ctx context.Context, in *DiskAttachRequest, opts ...grpc.CallOption) (*DiskAttachResponse, error)
	// VirtualMachineCreate creates a new virtual machine.
	VirtualMachineCreate(ctx context.Context, in *VirtualMachineCreateRequest, opts ...grpc.CallOption) (*VirtualMachineCreateResponse, error)
	// VirtualMachineDelete deletes a virtual machine.
	// The virtual machine must be stopped.
	VirtualMachineDelete(ctx context.Context, in *VirtualMachineDeleteRequest, opts ...grpc.CallOption) (*VirtualMachineDeleteResponse, error)
	// VirtualMachineStart starts a virtual machine.
	// The virtual machine must be stopped.
	VirtualMachineStart(ctx context.Context, in *VirtualMachineStartRequest, opts ...grpc.CallOption) (*VirtualMachineStartResponse, error)
	// VirtualMachineStop stops a virtual machine.
	// The virtual machine must be started.
	VirtualMachineStop(ctx context.Context, in *VirtualMachineStopRequest, opts ...grpc.CallOption) (*VirtualMachineStopResponse, error)
	// VirtualMachineList lists all virtual machines.
	VirtualMachineList(ctx context.Context, in *VirtualMachineListRequest, opts ...grpc.CallOption) (*VirtualMachineListResponse, error)
	// ImagePull pulls an image in the local storage.
	ImagePull(ctx context.Context, in *ImagePullRequest, opts ...grpc.CallOption) (*ImagePullResponse, error)
	// ImageCreate creates a new image.
	// The image can be created from a disk or from a remote resource.
	ImageCreate(ctx context.Context, in *ImageCreateRequest, opts ...grpc.CallOption) (*ImageCreateResponse, error)
	// ImageDelete deletes an image.
	ImageDelete(ctx context.Context, in *ImageDeleteRequest, opts ...grpc.CallOption) (*ImageDeleteResponse, error)
	// ImageList lists all images.
	ImageList(ctx context.Context, in *ImageListRequest, opts ...grpc.CallOption) (*ImageListResponse, error)
}

type vMKitClient struct {
	cc grpc.ClientConnInterface
}

func NewVMKitClient(cc grpc.ClientConnInterface) VMKitClient {
	return &vMKitClient{cc}
}

func (c *vMKitClient) DiskCreate(ctx context.Context, in *DiskCreateRequest, opts ...grpc.CallOption) (*DiskCreateResponse, error) {
	out := new(DiskCreateResponse)
	err := c.cc.Invoke(ctx, VMKit_DiskCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMKitClient) DiskDelete(ctx context.Context, in *DiskDeleteRequest, opts ...grpc.CallOption) (*DiskDeleteResponse, error) {
	out := new(DiskDeleteResponse)
	err := c.cc.Invoke(ctx, VMKit_DiskDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMKitClient) DiskResize(ctx context.Context, in *DiskResizeRequest, opts ...grpc.CallOption) (*DiskResizeResponse, error) {
	out := new(DiskResizeResponse)
	err := c.cc.Invoke(ctx, VMKit_DiskResize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMKitClient) DiskList(ctx context.Context, in *DiskListRequest, opts ...grpc.CallOption) (*DiskListResponse, error) {
	out := new(DiskListResponse)
	err := c.cc.Invoke(ctx, VMKit_DiskList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMKitClient) DiskAttach(ctx context.Context, in *DiskAttachRequest, opts ...grpc.CallOption) (*DiskAttachResponse, error) {
	out := new(DiskAttachResponse)
	err := c.cc.Invoke(ctx, VMKit_DiskAttach_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMKitClient) VirtualMachineCreate(ctx context.Context, in *VirtualMachineCreateRequest, opts ...grpc.CallOption) (*VirtualMachineCreateResponse, error) {
	out := new(VirtualMachineCreateResponse)
	err := c.cc.Invoke(ctx, VMKit_VirtualMachineCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMKitClient) VirtualMachineDelete(ctx context.Context, in *VirtualMachineDeleteRequest, opts ...grpc.CallOption) (*VirtualMachineDeleteResponse, error) {
	out := new(VirtualMachineDeleteResponse)
	err := c.cc.Invoke(ctx, VMKit_VirtualMachineDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMKitClient) VirtualMachineStart(ctx context.Context, in *VirtualMachineStartRequest, opts ...grpc.CallOption) (*VirtualMachineStartResponse, error) {
	out := new(VirtualMachineStartResponse)
	err := c.cc.Invoke(ctx, VMKit_VirtualMachineStart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMKitClient) VirtualMachineStop(ctx context.Context, in *VirtualMachineStopRequest, opts ...grpc.CallOption) (*VirtualMachineStopResponse, error) {
	out := new(VirtualMachineStopResponse)
	err := c.cc.Invoke(ctx, VMKit_VirtualMachineStop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMKitClient) VirtualMachineList(ctx context.Context, in *VirtualMachineListRequest, opts ...grpc.CallOption) (*VirtualMachineListResponse, error) {
	out := new(VirtualMachineListResponse)
	err := c.cc.Invoke(ctx, VMKit_VirtualMachineList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMKitClient) ImagePull(ctx context.Context, in *ImagePullRequest, opts ...grpc.CallOption) (*ImagePullResponse, error) {
	out := new(ImagePullResponse)
	err := c.cc.Invoke(ctx, VMKit_ImagePull_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMKitClient) ImageCreate(ctx context.Context, in *ImageCreateRequest, opts ...grpc.CallOption) (*ImageCreateResponse, error) {
	out := new(ImageCreateResponse)
	err := c.cc.Invoke(ctx, VMKit_ImageCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMKitClient) ImageDelete(ctx context.Context, in *ImageDeleteRequest, opts ...grpc.CallOption) (*ImageDeleteResponse, error) {
	out := new(ImageDeleteResponse)
	err := c.cc.Invoke(ctx, VMKit_ImageDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMKitClient) ImageList(ctx context.Context, in *ImageListRequest, opts ...grpc.CallOption) (*ImageListResponse, error) {
	out := new(ImageListResponse)
	err := c.cc.Invoke(ctx, VMKit_ImageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VMKitServer is the server API for VMKit service.
// All implementations must embed UnimplementedVMKitServer
// for forward compatibility
type VMKitServer interface {
	// DiskCreate creates a new disk.
	// The disk can be created from an image or from scratch.
	DiskCreate(context.Context, *DiskCreateRequest) (*DiskCreateResponse, error)
	// DiskDelete deletes a disk.
	// The disk must be detached from any virtual machine.
	DiskDelete(context.Context, *DiskDeleteRequest) (*DiskDeleteResponse, error)
	// DiskResize resizes a disk.
	// The disk must be detached from any virtual machine. The size must be greater than the current size.
	DiskResize(context.Context, *DiskResizeRequest) (*DiskResizeResponse, error)
	// DiskList lists all disks.
	DiskList(context.Context, *DiskListRequest) (*DiskListResponse, error)
	// DiskAttach attaches a disk to a virtual machine.
	// The virtual machine must be stopped.
	DiskAttach(context.Context, *DiskAttachRequest) (*DiskAttachResponse, error)
	// VirtualMachineCreate creates a new virtual machine.
	VirtualMachineCreate(context.Context, *VirtualMachineCreateRequest) (*VirtualMachineCreateResponse, error)
	// VirtualMachineDelete deletes a virtual machine.
	// The virtual machine must be stopped.
	VirtualMachineDelete(context.Context, *VirtualMachineDeleteRequest) (*VirtualMachineDeleteResponse, error)
	// VirtualMachineStart starts a virtual machine.
	// The virtual machine must be stopped.
	VirtualMachineStart(context.Context, *VirtualMachineStartRequest) (*VirtualMachineStartResponse, error)
	// VirtualMachineStop stops a virtual machine.
	// The virtual machine must be started.
	VirtualMachineStop(context.Context, *VirtualMachineStopRequest) (*VirtualMachineStopResponse, error)
	// VirtualMachineList lists all virtual machines.
	VirtualMachineList(context.Context, *VirtualMachineListRequest) (*VirtualMachineListResponse, error)
	// ImagePull pulls an image in the local storage.
	ImagePull(context.Context, *ImagePullRequest) (*ImagePullResponse, error)
	// ImageCreate creates a new image.
	// The image can be created from a disk or from a remote resource.
	ImageCreate(context.Context, *ImageCreateRequest) (*ImageCreateResponse, error)
	// ImageDelete deletes an image.
	ImageDelete(context.Context, *ImageDeleteRequest) (*ImageDeleteResponse, error)
	// ImageList lists all images.
	ImageList(context.Context, *ImageListRequest) (*ImageListResponse, error)
	mustEmbedUnimplementedVMKitServer()
}

// UnimplementedVMKitServer must be embedded to have forward compatible implementations.
type UnimplementedVMKitServer struct {
}

func (UnimplementedVMKitServer) DiskCreate(context.Context, *DiskCreateRequest) (*DiskCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiskCreate not implemented")
}
func (UnimplementedVMKitServer) DiskDelete(context.Context, *DiskDeleteRequest) (*DiskDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiskDelete not implemented")
}
func (UnimplementedVMKitServer) DiskResize(context.Context, *DiskResizeRequest) (*DiskResizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiskResize not implemented")
}
func (UnimplementedVMKitServer) DiskList(context.Context, *DiskListRequest) (*DiskListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiskList not implemented")
}
func (UnimplementedVMKitServer) DiskAttach(context.Context, *DiskAttachRequest) (*DiskAttachResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiskAttach not implemented")
}
func (UnimplementedVMKitServer) VirtualMachineCreate(context.Context, *VirtualMachineCreateRequest) (*VirtualMachineCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VirtualMachineCreate not implemented")
}
func (UnimplementedVMKitServer) VirtualMachineDelete(context.Context, *VirtualMachineDeleteRequest) (*VirtualMachineDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VirtualMachineDelete not implemented")
}
func (UnimplementedVMKitServer) VirtualMachineStart(context.Context, *VirtualMachineStartRequest) (*VirtualMachineStartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VirtualMachineStart not implemented")
}
func (UnimplementedVMKitServer) VirtualMachineStop(context.Context, *VirtualMachineStopRequest) (*VirtualMachineStopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VirtualMachineStop not implemented")
}
func (UnimplementedVMKitServer) VirtualMachineList(context.Context, *VirtualMachineListRequest) (*VirtualMachineListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VirtualMachineList not implemented")
}
func (UnimplementedVMKitServer) ImagePull(context.Context, *ImagePullRequest) (*ImagePullResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImagePull not implemented")
}
func (UnimplementedVMKitServer) ImageCreate(context.Context, *ImageCreateRequest) (*ImageCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImageCreate not implemented")
}
func (UnimplementedVMKitServer) ImageDelete(context.Context, *ImageDeleteRequest) (*ImageDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImageDelete not implemented")
}
func (UnimplementedVMKitServer) ImageList(context.Context, *ImageListRequest) (*ImageListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImageList not implemented")
}
func (UnimplementedVMKitServer) mustEmbedUnimplementedVMKitServer() {}

// UnsafeVMKitServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VMKitServer will
// result in compilation errors.
type UnsafeVMKitServer interface {
	mustEmbedUnimplementedVMKitServer()
}

func RegisterVMKitServer(s grpc.ServiceRegistrar, srv VMKitServer) {
	s.RegisterService(&VMKit_ServiceDesc, srv)
}

func _VMKit_DiskCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiskCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMKitServer).DiskCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMKit_DiskCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMKitServer).DiskCreate(ctx, req.(*DiskCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMKit_DiskDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiskDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMKitServer).DiskDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMKit_DiskDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMKitServer).DiskDelete(ctx, req.(*DiskDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMKit_DiskResize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiskResizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMKitServer).DiskResize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMKit_DiskResize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMKitServer).DiskResize(ctx, req.(*DiskResizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMKit_DiskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMKitServer).DiskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMKit_DiskList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMKitServer).DiskList(ctx, req.(*DiskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMKit_DiskAttach_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiskAttachRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMKitServer).DiskAttach(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMKit_DiskAttach_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMKitServer).DiskAttach(ctx, req.(*DiskAttachRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMKit_VirtualMachineCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualMachineCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMKitServer).VirtualMachineCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMKit_VirtualMachineCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMKitServer).VirtualMachineCreate(ctx, req.(*VirtualMachineCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMKit_VirtualMachineDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualMachineDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMKitServer).VirtualMachineDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMKit_VirtualMachineDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMKitServer).VirtualMachineDelete(ctx, req.(*VirtualMachineDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMKit_VirtualMachineStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualMachineStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMKitServer).VirtualMachineStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMKit_VirtualMachineStart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMKitServer).VirtualMachineStart(ctx, req.(*VirtualMachineStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMKit_VirtualMachineStop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualMachineStopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMKitServer).VirtualMachineStop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMKit_VirtualMachineStop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMKitServer).VirtualMachineStop(ctx, req.(*VirtualMachineStopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMKit_VirtualMachineList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualMachineListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMKitServer).VirtualMachineList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMKit_VirtualMachineList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMKitServer).VirtualMachineList(ctx, req.(*VirtualMachineListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMKit_ImagePull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImagePullRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMKitServer).ImagePull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMKit_ImagePull_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMKitServer).ImagePull(ctx, req.(*ImagePullRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMKit_ImageCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMKitServer).ImageCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMKit_ImageCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMKitServer).ImageCreate(ctx, req.(*ImageCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMKit_ImageDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMKitServer).ImageDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMKit_ImageDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMKitServer).ImageDelete(ctx, req.(*ImageDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VMKit_ImageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMKitServer).ImageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VMKit_ImageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMKitServer).ImageList(ctx, req.(*ImageListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VMKit_ServiceDesc is the grpc.ServiceDesc for VMKit service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VMKit_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "VMKit",
	HandlerType: (*VMKitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DiskCreate",
			Handler:    _VMKit_DiskCreate_Handler,
		},
		{
			MethodName: "DiskDelete",
			Handler:    _VMKit_DiskDelete_Handler,
		},
		{
			MethodName: "DiskResize",
			Handler:    _VMKit_DiskResize_Handler,
		},
		{
			MethodName: "DiskList",
			Handler:    _VMKit_DiskList_Handler,
		},
		{
			MethodName: "DiskAttach",
			Handler:    _VMKit_DiskAttach_Handler,
		},
		{
			MethodName: "VirtualMachineCreate",
			Handler:    _VMKit_VirtualMachineCreate_Handler,
		},
		{
			MethodName: "VirtualMachineDelete",
			Handler:    _VMKit_VirtualMachineDelete_Handler,
		},
		{
			MethodName: "VirtualMachineStart",
			Handler:    _VMKit_VirtualMachineStart_Handler,
		},
		{
			MethodName: "VirtualMachineStop",
			Handler:    _VMKit_VirtualMachineStop_Handler,
		},
		{
			MethodName: "VirtualMachineList",
			Handler:    _VMKit_VirtualMachineList_Handler,
		},
		{
			MethodName: "ImagePull",
			Handler:    _VMKit_ImagePull_Handler,
		},
		{
			MethodName: "ImageCreate",
			Handler:    _VMKit_ImageCreate_Handler,
		},
		{
			MethodName: "ImageDelete",
			Handler:    _VMKit_ImageDelete_Handler,
		},
		{
			MethodName: "ImageList",
			Handler:    _VMKit_ImageList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
