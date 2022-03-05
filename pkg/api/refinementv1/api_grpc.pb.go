// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: iamnande/cardmod/refinement/v1/api.proto

package refinementv1

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

// RefinementAPIClient is the client API for RefinementAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RefinementAPIClient interface {
	// Gets a refinement.
	GetRefinement(ctx context.Context, in *GetRefinementRequest, opts ...grpc.CallOption) (*Refinement, error)
	// Lists a collection of refinements.
	ListRefinements(ctx context.Context, in *ListRefinementsRequest, opts ...grpc.CallOption) (*ListRefinementsResponse, error)
}

type refinementAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewRefinementAPIClient(cc grpc.ClientConnInterface) RefinementAPIClient {
	return &refinementAPIClient{cc}
}

func (c *refinementAPIClient) GetRefinement(ctx context.Context, in *GetRefinementRequest, opts ...grpc.CallOption) (*Refinement, error) {
	out := new(Refinement)
	err := c.cc.Invoke(ctx, "/iamnande.cardmod.refinement.v1.RefinementAPI/GetRefinement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *refinementAPIClient) ListRefinements(ctx context.Context, in *ListRefinementsRequest, opts ...grpc.CallOption) (*ListRefinementsResponse, error) {
	out := new(ListRefinementsResponse)
	err := c.cc.Invoke(ctx, "/iamnande.cardmod.refinement.v1.RefinementAPI/ListRefinements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RefinementAPIServer is the server API for RefinementAPI service.
// All implementations must embed UnimplementedRefinementAPIServer
// for forward compatibility
type RefinementAPIServer interface {
	// Gets a refinement.
	GetRefinement(context.Context, *GetRefinementRequest) (*Refinement, error)
	// Lists a collection of refinements.
	ListRefinements(context.Context, *ListRefinementsRequest) (*ListRefinementsResponse, error)
	mustEmbedUnimplementedRefinementAPIServer()
}

// UnimplementedRefinementAPIServer must be embedded to have forward compatible implementations.
type UnimplementedRefinementAPIServer struct {
}

func (UnimplementedRefinementAPIServer) GetRefinement(context.Context, *GetRefinementRequest) (*Refinement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRefinement not implemented")
}
func (UnimplementedRefinementAPIServer) ListRefinements(context.Context, *ListRefinementsRequest) (*ListRefinementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRefinements not implemented")
}
func (UnimplementedRefinementAPIServer) mustEmbedUnimplementedRefinementAPIServer() {}

// UnsafeRefinementAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RefinementAPIServer will
// result in compilation errors.
type UnsafeRefinementAPIServer interface {
	mustEmbedUnimplementedRefinementAPIServer()
}

func RegisterRefinementAPIServer(s grpc.ServiceRegistrar, srv RefinementAPIServer) {
	s.RegisterService(&RefinementAPI_ServiceDesc, srv)
}

func _RefinementAPI_GetRefinement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRefinementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefinementAPIServer).GetRefinement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iamnande.cardmod.refinement.v1.RefinementAPI/GetRefinement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefinementAPIServer).GetRefinement(ctx, req.(*GetRefinementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RefinementAPI_ListRefinements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRefinementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefinementAPIServer).ListRefinements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iamnande.cardmod.refinement.v1.RefinementAPI/ListRefinements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefinementAPIServer).ListRefinements(ctx, req.(*ListRefinementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RefinementAPI_ServiceDesc is the grpc.ServiceDesc for RefinementAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RefinementAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iamnande.cardmod.refinement.v1.RefinementAPI",
	HandlerType: (*RefinementAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRefinement",
			Handler:    _RefinementAPI_GetRefinement_Handler,
		},
		{
			MethodName: "ListRefinements",
			Handler:    _RefinementAPI_ListRefinements_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iamnande/cardmod/refinement/v1/api.proto",
}
