package logic

import (
	"context"
	"io"

	"zerostream/streamrpc/internal/svc"
	"zerostream/streamrpc/streamrpc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
)

type RecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecordLogic {
	return &RecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RecordLogic) Record(stream streamrpc.Streamrpc_RecordServer) error {
	defer func() {
		trailer := metadata.Pairs("server-record", "record")
		stream.SetTrailer(trailer)
	}()

	// receive metadata - server side
	md, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		logx.Infof("record req header: %+v", md)
	}

	header := metadata.Pairs("server-record", "record")
	stream.SendHeader(header) // can't use SetHeader here

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&streamrpc.StreamResponse{Point: &streamrpc.StreamPoint{Name: "llb", Value: 1}})
		}
		if err != nil {
			return err
		}

		logx.Infof("stream recv name: %v, value: %v", msg.Point.Name, msg.Point.Value)
	}

	// return nil
}
