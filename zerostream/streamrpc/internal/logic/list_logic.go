package logic

import (
	"context"

	"zerostream/streamrpc/internal/svc"
	"zerostream/streamrpc/streamrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *streamrpc.StreamRequest, stream streamrpc.Streamrpc_ListServer) error {
	defer func() {
		trailer := metadata.Pairs("server-list", "list")
		stream.SetTrailer(trailer)
	}()

	// receive metada - server side
	md, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		logx.Infof("list req header: %+v", md)
	}

	header := metadata.Pairs("server-list", "list")
	stream.SetHeader(header)

	for i := 0; i < 10; i++ {
		if err := stream.Send(&streamrpc.StreamResponse{
			Point: &streamrpc.StreamPoint{
				Name:  in.Point.Name,
				Value: in.Point.Value + int32(i),
			},
		}); err != nil {
			return err
		}
	}

	return nil
}
