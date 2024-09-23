package logic

import (
	"context"

	"zerostream/streamapi/internal/svc"
	"zerostream/streamapi/internal/types"
	"zerostream/streamrpc/streamrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type StreamapiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStreamapiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StreamapiLogic {
	return &StreamapiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StreamapiLogic) Streamapi(req *types.Request) (resp *types.Response, err error) {
	// send metadata - client side
	ctx := metadata.AppendToOutgoingContext(l.ctx, "client-ping", "ping")
	// ctx = metadata.AppendToOutgoingContext(ctx, "k2", "v2")

	// receive metadata from unary call - client side
	var header, trailer metadata.MD
	r, err := l.svcCtx.RpcClient.Ping(ctx, &streamrpc.Request{Ping: "a"}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return nil, err
	}

	logx.Infof("ping resp header: %+v", header)
	logx.Infof("ping resp trailer: %+v", trailer)

	resp = &types.Response{
		Message: r.Pong,
	}

	return
}
