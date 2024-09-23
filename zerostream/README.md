# grpc streaming client and server

> https://mp.weixin.qq.com/s?__biz=MzUxMDI4MDc1NA==&mid=2247483721&idx=3&sn=b61db0379afd96e0149c279564d8efea&chksm=f9041414ce739d02c1554318a6e86942a0450266f27360913882860f24bc59268d315142f79b&cur_album_id=1383472721040064512&scene=189#wechat_redirect

### 根据proto文件生成对应的代码
```sh
goctl rpc protoc streamrpc.proto --go_out=. --go-grpc_out=. --zrpc_out=. --client=true --style=go_zero
```

### 格式化api文件
```sh
goctl api format --dir=.
```

### 根据api文件生成代码
```sh
goctl api go --api=streamapi.api --dir=. --style=go_zero
```

# gprc-metadata

> https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md