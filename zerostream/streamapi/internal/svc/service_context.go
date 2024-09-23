package svc

import (
	"zerostream/streamapi/internal/config"
	"zerostream/streamrpc/streamrpc"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	RpcClient streamrpc.StreamrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		RpcClient: streamrpc.NewStreamrpcClient(zrpc.MustNewClient(c.Rpc).Conn()),
	}
}
