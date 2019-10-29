package hello

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"natsmicro/helloworld2/rpc/nrpc_client"
	"natsmicro/helloworld2/rpc/proto/helloworld"
)

// Hello World
func Handler(r *ghttp.Request) {
	res, _ := nrpc_client.HelloWordNrpcCli.SayHello(helloworld.HelloRequest2{Name: "world"})
	glog.Info(res.Message)
	r.Response.Writeln(res.Message)
}
