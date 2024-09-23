// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: streamrpc.proto

package streamrpc

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
	Streamrpc_Ping_FullMethodName   = "/streamrpc.Streamrpc/Ping"
	Streamrpc_List_FullMethodName   = "/streamrpc.Streamrpc/List"
	Streamrpc_Record_FullMethodName = "/streamrpc.Streamrpc/Record"
	Streamrpc_Route_FullMethodName  = "/streamrpc.Streamrpc/Route"
)

// StreamrpcClient is the client API for Streamrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamrpcClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	List(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Streamrpc_ListClient, error)
	Record(ctx context.Context, opts ...grpc.CallOption) (Streamrpc_RecordClient, error)
	Route(ctx context.Context, opts ...grpc.CallOption) (Streamrpc_RouteClient, error)
}

type streamrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamrpcClient(cc grpc.ClientConnInterface) StreamrpcClient {
	return &streamrpcClient{cc}
}

func (c *streamrpcClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Streamrpc_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamrpcClient) List(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Streamrpc_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Streamrpc_ServiceDesc.Streams[0], Streamrpc_List_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &streamrpcListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Streamrpc_ListClient interface {
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamrpcListClient struct {
	grpc.ClientStream
}

func (x *streamrpcListClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamrpcClient) Record(ctx context.Context, opts ...grpc.CallOption) (Streamrpc_RecordClient, error) {
	stream, err := c.cc.NewStream(ctx, &Streamrpc_ServiceDesc.Streams[1], Streamrpc_Record_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &streamrpcRecordClient{stream}
	return x, nil
}

type Streamrpc_RecordClient interface {
	Send(*StreamRequest) error
	CloseAndRecv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamrpcRecordClient struct {
	grpc.ClientStream
}

func (x *streamrpcRecordClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamrpcRecordClient) CloseAndRecv() (*StreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamrpcClient) Route(ctx context.Context, opts ...grpc.CallOption) (Streamrpc_RouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &Streamrpc_ServiceDesc.Streams[2], Streamrpc_Route_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &streamrpcRouteClient{stream}
	return x, nil
}

type Streamrpc_RouteClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamrpcRouteClient struct {
	grpc.ClientStream
}

func (x *streamrpcRouteClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamrpcRouteClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamrpcServer is the server API for Streamrpc service.
// All implementations must embed UnimplementedStreamrpcServer
// for forward compatibility
type StreamrpcServer interface {
	Ping(context.Context, *Request) (*Response, error)
	List(*StreamRequest, Streamrpc_ListServer) error
	Record(Streamrpc_RecordServer) error
	Route(Streamrpc_RouteServer) error
	mustEmbedUnimplementedStreamrpcServer()
}

// UnimplementedStreamrpcServer must be embedded to have forward compatible implementations.
type UnimplementedStreamrpcServer struct {
}

func (UnimplementedStreamrpcServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedStreamrpcServer) List(*StreamRequest, Streamrpc_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedStreamrpcServer) Record(Streamrpc_RecordServer) error {
	return status.Errorf(codes.Unimplemented, "method Record not implemented")
}
func (UnimplementedStreamrpcServer) Route(Streamrpc_RouteServer) error {
	return status.Errorf(codes.Unimplemented, "method Route not implemented")
}
func (UnimplementedStreamrpcServer) mustEmbedUnimplementedStreamrpcServer() {}

// UnsafeStreamrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamrpcServer will
// result in compilation errors.
type UnsafeStreamrpcServer interface {
	mustEmbedUnimplementedStreamrpcServer()
}

func RegisterStreamrpcServer(s grpc.ServiceRegistrar, srv StreamrpcServer) {
	s.RegisterService(&Streamrpc_ServiceDesc, srv)
}

func _Streamrpc_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamrpcServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Streamrpc_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamrpcServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Streamrpc_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamrpcServer).List(m, &streamrpcListServer{stream})
}

type Streamrpc_ListServer interface {
	Send(*StreamResponse) error
	grpc.ServerStream
}

type streamrpcListServer struct {
	grpc.ServerStream
}

func (x *streamrpcListServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Streamrpc_Record_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamrpcServer).Record(&streamrpcRecordServer{stream})
}

type Streamrpc_RecordServer interface {
	SendAndClose(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamrpcRecordServer struct {
	grpc.ServerStream
}

func (x *streamrpcRecordServer) SendAndClose(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamrpcRecordServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Streamrpc_Route_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamrpcServer).Route(&streamrpcRouteServer{stream})
}

type Streamrpc_RouteServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamrpcRouteServer struct {
	grpc.ServerStream
}

func (x *streamrpcRouteServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamrpcRouteServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Streamrpc_ServiceDesc is the grpc.ServiceDesc for Streamrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Streamrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "streamrpc.Streamrpc",
	HandlerType: (*StreamrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Streamrpc_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Streamrpc_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Record",
			Handler:       _Streamrpc_Record_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Route",
			Handler:       _Streamrpc_Route_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "streamrpc.proto",
}
