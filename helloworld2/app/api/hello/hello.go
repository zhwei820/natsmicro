package hello

import (
    "github.com/gogf/gf/net/ghttp"
    "natsmicro/helloworld/rpc/nrpc_client"
    "natsmicro/helloworld/rpc/proto/helloworld"
    "time"
)

// Hello World
func Handler(r *ghttp.Request) {
    t1 := time.Now()
    res, _ := nrpc_client.HelloWordNrpcCli.SayHello(helloworld.HelloRequest{Name: "world"})
	//glog.Info(res.Message)
    println(time.Now().Sub(t1).Seconds())

	r.Response.Writeln("Hello World!" + res.Message)
}
