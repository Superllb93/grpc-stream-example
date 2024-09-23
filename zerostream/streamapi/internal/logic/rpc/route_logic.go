package rpc

import (
	"context"
	"io"

	"zerostream/streamapi/internal/svc"
	"zerostream/streamapi/internal/types"
	"zerostream/streamrpc/streamrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RouteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRouteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RouteLogic {
	return &RouteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RouteLogic) Route() (resp *types.StreamRpcRouteResponse, err error) {
	stream, err := l.svcCtx.RpcClient.Route(l.ctx)
	if err != nil {
		return nil, err
	}

	list := make([]*types.Point, 0)
	for i := 0; i < 10; i++ {
		if err := stream.Send(&streamrpc.StreamRequest{
			Point: &streamrpc.StreamPoint{
				Name:  "llb-client",
				Value: int32(i),
			},
		}); err != nil {
			return nil, err
		}

		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		list = append(list, &types.Point{Name: msg.Point.Name, Value: msg.Point.Value})

		logx.Infof("client stream recv name: %v, value: %v", msg.Point.Name, msg.Point.Value)
	}

	return &types.StreamRpcRouteResponse{List: list}, nil
}
