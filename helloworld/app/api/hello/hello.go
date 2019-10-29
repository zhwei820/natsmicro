package hello

import (
	"github.com/gogf/gf/net/ghttp"
	"natsmicro/helloworld2/rpc/nrpc_client"
	"natsmicro/helloworld2/rpc/proto/helloworld"
)

// Hello World
func Handler(r *ghttp.Request) {
	nrpc_client.HelloWordNrpcCli.SayHello(helloworld.HelloRequest2{Name: "world"})

	r.Response.Writeln("Hello World!")
}
