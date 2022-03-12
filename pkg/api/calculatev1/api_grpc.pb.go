// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: iamnande/cardmod/calculate/v1/api.proto

package calculatev1

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

// CalculateAPIClient is the client API for CalculateAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculateAPIClient interface {
	// Calculates quantities and sources of a given refinement.
	Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error)
}

type calculateAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculateAPIClient(cc grpc.ClientConnInterface) CalculateAPIClient {
	return &calculateAPIClient{cc}
}

func (c *calculateAPIClient) Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error) {
	out := new(CalculateResponse)
	err := c.cc.Invoke(ctx, "/iamnande.cardmod.calculate.v1.CalculateAPI/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculateAPIServer is the server API for CalculateAPI service.
// All implementations must embed UnimplementedCalculateAPIServer
// for forward compatibility
type CalculateAPIServer interface {
	// Calculates quantities and sources of a given refinement.
	Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error)
	mustEmbedUnimplementedCalculateAPIServer()
}

// UnimplementedCalculateAPIServer must be embedded to have forward compatible implementations.
type UnimplementedCalculateAPIServer struct {
}

func (UnimplementedCalculateAPIServer) Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedCalculateAPIServer) mustEmbedUnimplementedCalculateAPIServer() {}

// UnsafeCalculateAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculateAPIServer will
// result in compilation errors.
type UnsafeCalculateAPIServer interface {
	mustEmbedUnimplementedCalculateAPIServer()
}

func RegisterCalculateAPIServer(s grpc.ServiceRegistrar, srv CalculateAPIServer) {
	s.RegisterService(&CalculateAPI_ServiceDesc, srv)
}

func _CalculateAPI_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateAPIServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iamnande.cardmod.calculate.v1.CalculateAPI/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateAPIServer).Calculate(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalculateAPI_ServiceDesc is the grpc.ServiceDesc for CalculateAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculateAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iamnande.cardmod.calculate.v1.CalculateAPI",
	HandlerType: (*CalculateAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _CalculateAPI_Calculate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iamnande/cardmod/calculate/v1/api.proto",
}
