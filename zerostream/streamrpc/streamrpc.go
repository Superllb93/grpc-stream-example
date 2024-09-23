package main

import (
	"context"
	"flag"
	"fmt"

	"zerostream/streamrpc/internal/config"
	"zerostream/streamrpc/internal/server"
	"zerostream/streamrpc/internal/svc"
	"zerostream/streamrpc/streamrpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/streamrpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		streamrpc.RegisterStreamrpcServer(grpcServer, server.NewStreamrpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 添加拦截器
	s.AddUnaryInterceptors(exampleUnaryInterceptor)
	s.AddStreamInterceptors(exampleStreamInterceptor)

	s.AddOptions(grpc.MaxRecvMsgSize(1024*1024*10), grpc.MaxSendMsgSize(1024*1024*10))

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func exampleUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// TODO: fill your logic here
	return handler(ctx, req)
}
func exampleStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// TODO: fill your logic here
	return handler(srv, ss)
}
