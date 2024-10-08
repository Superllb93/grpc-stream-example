// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: streamrpc.proto

package server

import (
	"context"

	"zerostream/streamrpc/internal/logic"
	"zerostream/streamrpc/internal/svc"
	"zerostream/streamrpc/streamrpc"
)

type StreamrpcServer struct {
	svcCtx *svc.ServiceContext
	streamrpc.UnimplementedStreamrpcServer
}

func NewStreamrpcServer(svcCtx *svc.ServiceContext) *StreamrpcServer {
	return &StreamrpcServer{
		svcCtx: svcCtx,
	}
}

func (s *StreamrpcServer) Ping(ctx context.Context, in *streamrpc.Request) (*streamrpc.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *StreamrpcServer) List(in *streamrpc.StreamRequest, stream streamrpc.Streamrpc_ListServer) error {
	l := logic.NewListLogic(stream.Context(), s.svcCtx)
	return l.List(in, stream)
}

func (s *StreamrpcServer) Record(stream streamrpc.Streamrpc_RecordServer) error {
	l := logic.NewRecordLogic(stream.Context(), s.svcCtx)
	return l.Record(stream)
}

func (s *StreamrpcServer) Route(stream streamrpc.Streamrpc_RouteServer) error {
	l := logic.NewRouteLogic(stream.Context(), s.svcCtx)
	return l.Route(stream)
}
