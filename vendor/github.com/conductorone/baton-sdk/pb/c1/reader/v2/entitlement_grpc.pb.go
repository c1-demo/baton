// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: c1/reader/v2/entitlement.proto

package v2

import (
	context "context"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EntitlementsReaderServiceClient is the client API for EntitlementsReaderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntitlementsReaderServiceClient interface {
	GetEntitlement(ctx context.Context, in *EntitlementsReaderServiceGetEntitlementRequest, opts ...grpc.CallOption) (*v2.Entitlement, error)
}

type entitlementsReaderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEntitlementsReaderServiceClient(cc grpc.ClientConnInterface) EntitlementsReaderServiceClient {
	return &entitlementsReaderServiceClient{cc}
}

func (c *entitlementsReaderServiceClient) GetEntitlement(ctx context.Context, in *EntitlementsReaderServiceGetEntitlementRequest, opts ...grpc.CallOption) (*v2.Entitlement, error) {
	out := new(v2.Entitlement)
	err := c.cc.Invoke(ctx, "/c1.reader.v2.EntitlementsReaderService/GetEntitlement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntitlementsReaderServiceServer is the server API for EntitlementsReaderService service.
// All implementations should embed UnimplementedEntitlementsReaderServiceServer
// for forward compatibility
type EntitlementsReaderServiceServer interface {
	GetEntitlement(context.Context, *EntitlementsReaderServiceGetEntitlementRequest) (*v2.Entitlement, error)
}

// UnimplementedEntitlementsReaderServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEntitlementsReaderServiceServer struct {
}

func (UnimplementedEntitlementsReaderServiceServer) GetEntitlement(context.Context, *EntitlementsReaderServiceGetEntitlementRequest) (*v2.Entitlement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntitlement not implemented")
}

// UnsafeEntitlementsReaderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntitlementsReaderServiceServer will
// result in compilation errors.
type UnsafeEntitlementsReaderServiceServer interface {
	mustEmbedUnimplementedEntitlementsReaderServiceServer()
}

func RegisterEntitlementsReaderServiceServer(s grpc.ServiceRegistrar, srv EntitlementsReaderServiceServer) {
	s.RegisterService(&EntitlementsReaderService_ServiceDesc, srv)
}

func _EntitlementsReaderService_GetEntitlement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntitlementsReaderServiceGetEntitlementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntitlementsReaderServiceServer).GetEntitlement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/c1.reader.v2.EntitlementsReaderService/GetEntitlement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntitlementsReaderServiceServer).GetEntitlement(ctx, req.(*EntitlementsReaderServiceGetEntitlementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EntitlementsReaderService_ServiceDesc is the grpc.ServiceDesc for EntitlementsReaderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntitlementsReaderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "c1.reader.v2.EntitlementsReaderService",
	HandlerType: (*EntitlementsReaderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntitlement",
			Handler:    _EntitlementsReaderService_GetEntitlement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "c1/reader/v2/entitlement.proto",
}
