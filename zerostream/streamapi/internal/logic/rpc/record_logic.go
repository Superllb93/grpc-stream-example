package rpc

import (
	"context"

	"zerostream/streamapi/internal/svc"
	"zerostream/streamapi/internal/types"
	"zerostream/streamrpc/streamrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
)

type RecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecordLogic {
	return &RecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecordLogic) Record(req *types.StreamRpcRecordRequest) (resp *types.StreamRpcRecordResponse, err error) {
	ctx := metadata.AppendToOutgoingContext(l.ctx, "client-record", "record")

	stream, err := l.svcCtx.RpcClient.Record(ctx)
	if err != nil {
		return nil, err
	}

	header, _ := stream.Header()
	logx.Infof("record resp header: %+v", header)

	for i := 0; i < 10; i++ {
		if err := stream.Send(&streamrpc.StreamRequest{
			Point: &streamrpc.StreamPoint{
				Name:  req.Name,
				Value: req.Value + int32(i),
			},
		}); err != nil {
			return nil, err
		}
	}

	msg, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	// Read the trailer after the RPC is finished
	trailer := stream.Trailer()
	logx.Infof("record resp trailer: %+v", trailer)

	return &types.StreamRpcRecordResponse{List: []*types.Point{{Name: msg.Point.Name, Value: msg.Point.Value}}}, nil
}
