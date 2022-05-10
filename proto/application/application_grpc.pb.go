// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: application.proto

package application

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

// ApplicationClient is the client API for Application service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApplicationClient interface {
	GetApplicationById(ctx context.Context, in *GetApplicationByIdRequest, opts ...grpc.CallOption) (*GetApplicationByIdReply, error)
}

type applicationClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationClient(cc grpc.ClientConnInterface) ApplicationClient {
	return &applicationClient{cc}
}

func (c *applicationClient) GetApplicationById(ctx context.Context, in *GetApplicationByIdRequest, opts ...grpc.CallOption) (*GetApplicationByIdReply, error) {
	out := new(GetApplicationByIdReply)
	err := c.cc.Invoke(ctx, "/application.Application/GetApplicationById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationServer is the server API for Application service.
// All implementations must embed UnimplementedApplicationServer
// for forward compatibility
type ApplicationServer interface {
	GetApplicationById(context.Context, *GetApplicationByIdRequest) (*GetApplicationByIdReply, error)
	mustEmbedUnimplementedApplicationServer()
}

// UnimplementedApplicationServer must be embedded to have forward compatible implementations.
type UnimplementedApplicationServer struct {
}

func (UnimplementedApplicationServer) GetApplicationById(context.Context, *GetApplicationByIdRequest) (*GetApplicationByIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplicationById not implemented")
}
func (UnimplementedApplicationServer) mustEmbedUnimplementedApplicationServer() {}

// UnsafeApplicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApplicationServer will
// result in compilation errors.
type UnsafeApplicationServer interface {
	mustEmbedUnimplementedApplicationServer()
}

func RegisterApplicationServer(s grpc.ServiceRegistrar, srv ApplicationServer) {
	s.RegisterService(&Application_ServiceDesc, srv)
}

func _Application_GetApplicationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).GetApplicationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.Application/GetApplicationById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).GetApplicationById(ctx, req.(*GetApplicationByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Application_ServiceDesc is the grpc.ServiceDesc for Application service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Application_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "application.Application",
	HandlerType: (*ApplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApplicationById",
			Handler:    _Application_GetApplicationById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application.proto",
}