// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: net/redeemer.proto

package net

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

// TicketRedeemerClient is the client API for TicketRedeemer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketRedeemerClient interface {
	QueueTicket(ctx context.Context, in *Ticket, opts ...grpc.CallOption) (*QueueTicketRes, error)
	MaxFloat(ctx context.Context, in *MaxFloatReq, opts ...grpc.CallOption) (*MaxFloatUpdate, error)
	MonitorMaxFloat(ctx context.Context, in *MaxFloatReq, opts ...grpc.CallOption) (TicketRedeemer_MonitorMaxFloatClient, error)
}

type ticketRedeemerClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketRedeemerClient(cc grpc.ClientConnInterface) TicketRedeemerClient {
	return &ticketRedeemerClient{cc}
}

func (c *ticketRedeemerClient) QueueTicket(ctx context.Context, in *Ticket, opts ...grpc.CallOption) (*QueueTicketRes, error) {
	out := new(QueueTicketRes)
	err := c.cc.Invoke(ctx, "/net.TicketRedeemer/QueueTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketRedeemerClient) MaxFloat(ctx context.Context, in *MaxFloatReq, opts ...grpc.CallOption) (*MaxFloatUpdate, error) {
	out := new(MaxFloatUpdate)
	err := c.cc.Invoke(ctx, "/net.TicketRedeemer/MaxFloat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketRedeemerClient) MonitorMaxFloat(ctx context.Context, in *MaxFloatReq, opts ...grpc.CallOption) (TicketRedeemer_MonitorMaxFloatClient, error) {
	stream, err := c.cc.NewStream(ctx, &TicketRedeemer_ServiceDesc.Streams[0], "/net.TicketRedeemer/MonitorMaxFloat", opts...)
	if err != nil {
		return nil, err
	}
	x := &ticketRedeemerMonitorMaxFloatClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TicketRedeemer_MonitorMaxFloatClient interface {
	Recv() (*MaxFloatUpdate, error)
	grpc.ClientStream
}

type ticketRedeemerMonitorMaxFloatClient struct {
	grpc.ClientStream
}

func (x *ticketRedeemerMonitorMaxFloatClient) Recv() (*MaxFloatUpdate, error) {
	m := new(MaxFloatUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TicketRedeemerServer is the server API for TicketRedeemer service.
// All implementations must embed UnimplementedTicketRedeemerServer
// for forward compatibility
type TicketRedeemerServer interface {
	QueueTicket(context.Context, *Ticket) (*QueueTicketRes, error)
	MaxFloat(context.Context, *MaxFloatReq) (*MaxFloatUpdate, error)
	MonitorMaxFloat(*MaxFloatReq, TicketRedeemer_MonitorMaxFloatServer) error
	mustEmbedUnimplementedTicketRedeemerServer()
}

// UnimplementedTicketRedeemerServer must be embedded to have forward compatible implementations.
type UnimplementedTicketRedeemerServer struct {
}

func (UnimplementedTicketRedeemerServer) QueueTicket(context.Context, *Ticket) (*QueueTicketRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueueTicket not implemented")
}
func (UnimplementedTicketRedeemerServer) MaxFloat(context.Context, *MaxFloatReq) (*MaxFloatUpdate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MaxFloat not implemented")
}
func (UnimplementedTicketRedeemerServer) MonitorMaxFloat(*MaxFloatReq, TicketRedeemer_MonitorMaxFloatServer) error {
	return status.Errorf(codes.Unimplemented, "method MonitorMaxFloat not implemented")
}
func (UnimplementedTicketRedeemerServer) mustEmbedUnimplementedTicketRedeemerServer() {}

// UnsafeTicketRedeemerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketRedeemerServer will
// result in compilation errors.
type UnsafeTicketRedeemerServer interface {
	mustEmbedUnimplementedTicketRedeemerServer()
}

func RegisterTicketRedeemerServer(s grpc.ServiceRegistrar, srv TicketRedeemerServer) {
	s.RegisterService(&TicketRedeemer_ServiceDesc, srv)
}

func _TicketRedeemer_QueueTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ticket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketRedeemerServer).QueueTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/net.TicketRedeemer/QueueTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketRedeemerServer).QueueTicket(ctx, req.(*Ticket))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketRedeemer_MaxFloat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MaxFloatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketRedeemerServer).MaxFloat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/net.TicketRedeemer/MaxFloat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketRedeemerServer).MaxFloat(ctx, req.(*MaxFloatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketRedeemer_MonitorMaxFloat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MaxFloatReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TicketRedeemerServer).MonitorMaxFloat(m, &ticketRedeemerMonitorMaxFloatServer{stream})
}

type TicketRedeemer_MonitorMaxFloatServer interface {
	Send(*MaxFloatUpdate) error
	grpc.ServerStream
}

type ticketRedeemerMonitorMaxFloatServer struct {
	grpc.ServerStream
}

func (x *ticketRedeemerMonitorMaxFloatServer) Send(m *MaxFloatUpdate) error {
	return x.ServerStream.SendMsg(m)
}

// TicketRedeemer_ServiceDesc is the grpc.ServiceDesc for TicketRedeemer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketRedeemer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "net.TicketRedeemer",
	HandlerType: (*TicketRedeemerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueueTicket",
			Handler:    _TicketRedeemer_QueueTicket_Handler,
		},
		{
			MethodName: "MaxFloat",
			Handler:    _TicketRedeemer_MaxFloat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorMaxFloat",
			Handler:       _TicketRedeemer_MonitorMaxFloat_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "net/redeemer.proto",
}
