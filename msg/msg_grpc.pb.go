// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.17.3
// source: msg/msg.proto

package msg

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

// MessengerClient is the client API for Messenger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessengerClient interface {
	SendMsg(ctx context.Context, opts ...grpc.CallOption) (Messenger_SendMsgClient, error)
}

type messengerClient struct {
	cc grpc.ClientConnInterface
}

func NewMessengerClient(cc grpc.ClientConnInterface) MessengerClient {
	return &messengerClient{cc}
}

func (c *messengerClient) SendMsg(ctx context.Context, opts ...grpc.CallOption) (Messenger_SendMsgClient, error) {
	stream, err := c.cc.NewStream(ctx, &Messenger_ServiceDesc.Streams[0], "/msg.Messenger/SendMsg", opts...)
	if err != nil {
		return nil, err
	}
	x := &messengerSendMsgClient{stream}
	return x, nil
}

type Messenger_SendMsgClient interface {
	Send(*Msg) error
	Recv() (*Msg, error)
	grpc.ClientStream
}

type messengerSendMsgClient struct {
	grpc.ClientStream
}

func (x *messengerSendMsgClient) Send(m *Msg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messengerSendMsgClient) Recv() (*Msg, error) {
	m := new(Msg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessengerServer is the server API for Messenger service.
// All implementations must embed UnimplementedMessengerServer
// for forward compatibility
type MessengerServer interface {
	SendMsg(Messenger_SendMsgServer) error
	mustEmbedUnimplementedMessengerServer()
}

// UnimplementedMessengerServer must be embedded to have forward compatible implementations.
type UnimplementedMessengerServer struct {
}

func (UnimplementedMessengerServer) SendMsg(Messenger_SendMsgServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}
func (UnimplementedMessengerServer) mustEmbedUnimplementedMessengerServer() {}

// UnsafeMessengerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessengerServer will
// result in compilation errors.
type UnsafeMessengerServer interface {
	mustEmbedUnimplementedMessengerServer()
}

func RegisterMessengerServer(s grpc.ServiceRegistrar, srv MessengerServer) {
	s.RegisterService(&Messenger_ServiceDesc, srv)
}

func _Messenger_SendMsg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessengerServer).SendMsg(&messengerSendMsgServer{stream})
}

type Messenger_SendMsgServer interface {
	Send(*Msg) error
	Recv() (*Msg, error)
	grpc.ServerStream
}

type messengerSendMsgServer struct {
	grpc.ServerStream
}

func (x *messengerSendMsgServer) Send(m *Msg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messengerSendMsgServer) Recv() (*Msg, error) {
	m := new(Msg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Messenger_ServiceDesc is the grpc.ServiceDesc for Messenger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messenger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msg.Messenger",
	HandlerType: (*MessengerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMsg",
			Handler:       _Messenger_SendMsg_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "msg/msg.proto",
}
