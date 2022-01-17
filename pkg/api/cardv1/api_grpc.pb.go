// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: iamnande/cardmod/card/v1/api.proto

package cardv1

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

// CardAPIClient is the client API for CardAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardAPIClient interface {
	// ListCards lists all available card resources.
	ListCards(ctx context.Context, in *ListCardsRequest, opts ...grpc.CallOption) (*ListCardsResponse, error)
	// GetCard gets a single card.
	GetCard(ctx context.Context, in *GetCardRequest, opts ...grpc.CallOption) (*Card, error)
}

type cardAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewCardAPIClient(cc grpc.ClientConnInterface) CardAPIClient {
	return &cardAPIClient{cc}
}

func (c *cardAPIClient) ListCards(ctx context.Context, in *ListCardsRequest, opts ...grpc.CallOption) (*ListCardsResponse, error) {
	out := new(ListCardsResponse)
	err := c.cc.Invoke(ctx, "/iamnande.cardmod.card.v1.CardAPI/ListCards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardAPIClient) GetCard(ctx context.Context, in *GetCardRequest, opts ...grpc.CallOption) (*Card, error) {
	out := new(Card)
	err := c.cc.Invoke(ctx, "/iamnande.cardmod.card.v1.CardAPI/GetCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardAPIServer is the server API for CardAPI service.
// All implementations must embed UnimplementedCardAPIServer
// for forward compatibility
type CardAPIServer interface {
	// ListCards lists all available card resources.
	ListCards(context.Context, *ListCardsRequest) (*ListCardsResponse, error)
	// GetCard gets a single card.
	GetCard(context.Context, *GetCardRequest) (*Card, error)
	mustEmbedUnimplementedCardAPIServer()
}

// UnimplementedCardAPIServer must be embedded to have forward compatible implementations.
type UnimplementedCardAPIServer struct {
}

func (UnimplementedCardAPIServer) ListCards(context.Context, *ListCardsRequest) (*ListCardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCards not implemented")
}
func (UnimplementedCardAPIServer) GetCard(context.Context, *GetCardRequest) (*Card, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCard not implemented")
}
func (UnimplementedCardAPIServer) mustEmbedUnimplementedCardAPIServer() {}

// UnsafeCardAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardAPIServer will
// result in compilation errors.
type UnsafeCardAPIServer interface {
	mustEmbedUnimplementedCardAPIServer()
}

func RegisterCardAPIServer(s grpc.ServiceRegistrar, srv CardAPIServer) {
	s.RegisterService(&CardAPI_ServiceDesc, srv)
}

func _CardAPI_ListCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardAPIServer).ListCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iamnande.cardmod.card.v1.CardAPI/ListCards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardAPIServer).ListCards(ctx, req.(*ListCardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardAPI_GetCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardAPIServer).GetCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iamnande.cardmod.card.v1.CardAPI/GetCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardAPIServer).GetCard(ctx, req.(*GetCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CardAPI_ServiceDesc is the grpc.ServiceDesc for CardAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CardAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iamnande.cardmod.card.v1.CardAPI",
	HandlerType: (*CardAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCards",
			Handler:    _CardAPI_ListCards_Handler,
		},
		{
			MethodName: "GetCard",
			Handler:    _CardAPI_GetCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iamnande/cardmod/card/v1/api.proto",
}
