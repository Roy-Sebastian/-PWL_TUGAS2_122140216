// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: proto/shipping.proto

package shippingpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ShippingService_StartShipping_FullMethodName  = "/shipping.ShippingService/StartShipping"
	ShippingService_CancelShipping_FullMethodName = "/shipping.ShippingService/CancelShipping"
)

// ShippingServiceClient is the client API for ShippingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShippingServiceClient interface {
	StartShipping(ctx context.Context, in *StartShippingRequest, opts ...grpc.CallOption) (*StartShippingResponse, error)
	CancelShipping(ctx context.Context, in *CancelShippingRequest, opts ...grpc.CallOption) (*CancelShippingResponse, error)
}

type shippingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShippingServiceClient(cc grpc.ClientConnInterface) ShippingServiceClient {
	return &shippingServiceClient{cc}
}

func (c *shippingServiceClient) StartShipping(ctx context.Context, in *StartShippingRequest, opts ...grpc.CallOption) (*StartShippingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartShippingResponse)
	err := c.cc.Invoke(ctx, ShippingService_StartShipping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) CancelShipping(ctx context.Context, in *CancelShippingRequest, opts ...grpc.CallOption) (*CancelShippingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelShippingResponse)
	err := c.cc.Invoke(ctx, ShippingService_CancelShipping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServiceServer is the server API for ShippingService service.
// All implementations must embed UnimplementedShippingServiceServer
// for forward compatibility.
type ShippingServiceServer interface {
	StartShipping(context.Context, *StartShippingRequest) (*StartShippingResponse, error)
	CancelShipping(context.Context, *CancelShippingRequest) (*CancelShippingResponse, error)
	mustEmbedUnimplementedShippingServiceServer()
}

// UnimplementedShippingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedShippingServiceServer struct{}

func (UnimplementedShippingServiceServer) StartShipping(context.Context, *StartShippingRequest) (*StartShippingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartShipping not implemented")
}
func (UnimplementedShippingServiceServer) CancelShipping(context.Context, *CancelShippingRequest) (*CancelShippingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelShipping not implemented")
}
func (UnimplementedShippingServiceServer) mustEmbedUnimplementedShippingServiceServer() {}
func (UnimplementedShippingServiceServer) testEmbeddedByValue()                         {}

// UnsafeShippingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShippingServiceServer will
// result in compilation errors.
type UnsafeShippingServiceServer interface {
	mustEmbedUnimplementedShippingServiceServer()
}

func RegisterShippingServiceServer(s grpc.ServiceRegistrar, srv ShippingServiceServer) {
	// If the following call pancis, it indicates UnimplementedShippingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ShippingService_ServiceDesc, srv)
}

func _ShippingService_StartShipping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartShippingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).StartShipping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingService_StartShipping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).StartShipping(ctx, req.(*StartShippingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_CancelShipping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelShippingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).CancelShipping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingService_CancelShipping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).CancelShipping(ctx, req.(*CancelShippingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShippingService_ServiceDesc is the grpc.ServiceDesc for ShippingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShippingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shipping.ShippingService",
	HandlerType: (*ShippingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartShipping",
			Handler:    _ShippingService_StartShipping_Handler,
		},
		{
			MethodName: "CancelShipping",
			Handler:    _ShippingService_CancelShipping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/shipping.proto",
}
