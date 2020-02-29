package hello

import (
	"github.com/gogf/gf/net/ghttp"
	"natsmicro/apps/helloworld/rpc/nrpc_client"
	"natsmicro/apps/helloworld/rpc/proto/helloworld"
	"time"
)

// Hello World
func Handler(r *ghttp.Request) {
	t1 := time.Now()
	res, _ := nrpc_client.HelloWordNrpcCli.SayHello(helloworld.HelloRequest{Name: "world"})
	//glog.Info(res.Message)
	println(time.Now().Sub(t1).Seconds())

	//u, err := gadmin_user.FindOne("id", 1)
	//if err != nil {
	//	glog.Debug(err)
	//}
	//g.Dump(u)

	r.Response.Writeln("Hello World!" + res.Message)
}
