package hello

import (
	"github.com/gogf/gf/net/ghttp"
	"natsmicro/helloworld/rpc/nrpc_client"
	"natsmicro/helloworld/rpc/proto/helloworld"
)

// Hello World
func Handler(r *ghttp.Request) {
	nrpc_client.HelloWordNrpcCli.SayHello(helloworld.HelloRequest{Name: "world"})
	r.Response.Writeln("Hello World!")
}
