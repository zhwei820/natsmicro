package hello

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"natsmicro/helloworld/rpc/nrpc_client"
	"natsmicro/helloworld/rpc/proto/helloworld"
)

// Hello World
func Handler(r *ghttp.Request) {
	res, _ := nrpc_client.HelloWordNrpcCli.SayHello(helloworld.HelloRequest{Name: "world"})
	glog.Info(res.Message)

	r.Response.Writeln("Hello World!" + res.Message)
}
