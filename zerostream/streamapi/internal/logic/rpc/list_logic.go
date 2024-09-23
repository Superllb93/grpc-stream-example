package rpc

import (
	"context"
	"io"

	"zerostream/streamapi/internal/svc"
	"zerostream/streamapi/internal/types"
	"zerostream/streamrpc/streamrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.StreamRpcListRequest) (resp *types.StreamRpcListResponse, err error) {
	// send metadata - client side
	ctx := metadata.AppendToOutgoingContext(l.ctx, "client-list", "list")

	stream, err := l.svcCtx.RpcClient.List(ctx, &streamrpc.StreamRequest{
		Point: &streamrpc.StreamPoint{
			Name:  req.Name,
			Value: req.Value,
		},
	})
	if err != nil {
		return nil, err
	}

	// receive metadata from streaming call - client side
	header, _ := stream.Header()
	logx.Infof("list resp header: %+v", header)

	list := make([]*types.Point, 0)
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		list = append(list, &types.Point{Name: msg.Point.Name, Value: msg.Point.Value})
	}

	// Read the trailer after the RPC is finished
	trailer := stream.Trailer()
	logx.Infof("list resp trailer: %+v", trailer)

	return &types.StreamRpcListResponse{List: list}, nil
}
