package logic

import (
	"context"

	"zerostream/streamrpc/internal/svc"
	"zerostream/streamrpc/streamrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *streamrpc.Request) (*streamrpc.Response, error) {
	defer func() {
		trailer := metadata.Pairs("server-ping", "ping")
		grpc.SetTrailer(l.ctx, trailer)
	}()

	// receive metada - server side
	md, ok := metadata.FromIncomingContext(l.ctx)
	if ok {
		logx.Infof("ping req header: %+v", md)
	}

	// send metadata - server side
	header := metadata.Pairs("server-ping", "ping")
	grpc.SetHeader(l.ctx, header)

	return &streamrpc.Response{Pong: in.Ping}, nil
}
