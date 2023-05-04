// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: sport/sport.proto

package sport

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
	Sport_ListEvents_FullMethodName   = "/sport.Sport/ListEvents"
	Sport_GetEventById_FullMethodName = "/sport.Sport/GetEventById"
)

// SportClient is the client API for Sport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SportClient interface {
	// ListEventys returns a list of all events.
	ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error)
	// GetEventById returns the event found using the id.
	GetEventById(ctx context.Context, in *GetEventByIdRequest, opts ...grpc.CallOption) (*GetEventByIdResponse, error)
}

type sportClient struct {
	cc grpc.ClientConnInterface
}

func NewSportClient(cc grpc.ClientConnInterface) SportClient {
	return &sportClient{cc}
}

func (c *sportClient) ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error) {
	out := new(ListEventsResponse)
	err := c.cc.Invoke(ctx, Sport_ListEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sportClient) GetEventById(ctx context.Context, in *GetEventByIdRequest, opts ...grpc.CallOption) (*GetEventByIdResponse, error) {
	out := new(GetEventByIdResponse)
	err := c.cc.Invoke(ctx, Sport_GetEventById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SportServer is the server API for Sport service.
// All implementations must embed UnimplementedSportServer
// for forward compatibility
type SportServer interface {
	// ListEventys returns a list of all events.
	ListEvents(context.Context, *ListEventsRequest) (*ListEventsResponse, error)
	// GetEventById returns the event found using the id.
	GetEventById(context.Context, *GetEventByIdRequest) (*GetEventByIdResponse, error)
	mustEmbedUnimplementedSportServer()
}

// UnimplementedSportServer must be embedded to have forward compatible implementations.
type UnimplementedSportServer struct {
}

func (UnimplementedSportServer) ListEvents(context.Context, *ListEventsRequest) (*ListEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEvents not implemented")
}
func (UnimplementedSportServer) GetEventById(context.Context, *GetEventByIdRequest) (*GetEventByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventById not implemented")
}
func (UnimplementedSportServer) mustEmbedUnimplementedSportServer() {}

// UnsafeSportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SportServer will
// result in compilation errors.
type UnsafeSportServer interface {
	mustEmbedUnimplementedSportServer()
}

func RegisterSportServer(s grpc.ServiceRegistrar, srv SportServer) {
	s.RegisterService(&Sport_ServiceDesc, srv)
}

func _Sport_ListEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SportServer).ListEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sport_ListEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SportServer).ListEvents(ctx, req.(*ListEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sport_GetEventById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SportServer).GetEventById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sport_GetEventById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SportServer).GetEventById(ctx, req.(*GetEventByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sport_ServiceDesc is the grpc.ServiceDesc for Sport service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sport_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sport.Sport",
	HandlerType: (*SportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEvents",
			Handler:    _Sport_ListEvents_Handler,
		},
		{
			MethodName: "GetEventById",
			Handler:    _Sport_GetEventById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sport/sport.proto",
}
