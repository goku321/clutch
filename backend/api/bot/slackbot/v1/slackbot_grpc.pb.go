// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package slackbotv1

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

// SlackBotAPIClient is the client API for SlackBotAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SlackBotAPIClient interface {
	Event(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
}

type slackBotAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewSlackBotAPIClient(cc grpc.ClientConnInterface) SlackBotAPIClient {
	return &slackBotAPIClient{cc}
}

func (c *slackBotAPIClient) Event(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/clutch.bot.slackbot.v1.SlackBotAPI/Event", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SlackBotAPIServer is the server API for SlackBotAPI service.
// All implementations should embed UnimplementedSlackBotAPIServer
// for forward compatibility
type SlackBotAPIServer interface {
	Event(context.Context, *EventRequest) (*EventResponse, error)
}

// UnimplementedSlackBotAPIServer should be embedded to have forward compatible implementations.
type UnimplementedSlackBotAPIServer struct {
}

func (UnimplementedSlackBotAPIServer) Event(context.Context, *EventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Event not implemented")
}

// UnsafeSlackBotAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SlackBotAPIServer will
// result in compilation errors.
type UnsafeSlackBotAPIServer interface {
	mustEmbedUnimplementedSlackBotAPIServer()
}

func RegisterSlackBotAPIServer(s grpc.ServiceRegistrar, srv SlackBotAPIServer) {
	s.RegisterService(&SlackBotAPI_ServiceDesc, srv)
}

func _SlackBotAPI_Event_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlackBotAPIServer).Event(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.bot.slackbot.v1.SlackBotAPI/Event",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlackBotAPIServer).Event(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SlackBotAPI_ServiceDesc is the grpc.ServiceDesc for SlackBotAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SlackBotAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.bot.slackbot.v1.SlackBotAPI",
	HandlerType: (*SlackBotAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Event",
			Handler:    _SlackBotAPI_Event_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bot/slackbot/v1/slackbot.proto",
}