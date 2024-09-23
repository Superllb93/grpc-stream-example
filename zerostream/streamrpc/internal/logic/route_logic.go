package logic

import (
	"context"
	"io"

	"zerostream/streamrpc/internal/svc"
	"zerostream/streamrpc/streamrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RouteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRouteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RouteLogic {
	return &RouteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RouteLogic) Route(stream streamrpc.Streamrpc_RouteServer) error {
	i := 0

	for {
		if err := stream.Send(&streamrpc.StreamResponse{
			Point: &streamrpc.StreamPoint{
				Name:  "llb-server",
				Value: int32(i),
			},
		}); err != nil {
			return err
		}

		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		i++

		logx.Infof("server stream recv name: %v, value: %v", msg.Point.Name, msg.Point.Value)
	}

	// return nil
}
