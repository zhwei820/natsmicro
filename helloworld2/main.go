package main

import (
	_ "natsmicro/helloworld2/boot"
	_ "natsmicro/helloworld2/router"
	"github.com/gogf/gf/frame/g"
	"natsmicro/helloworld2/rpc/nrpc_server"
)

func init()  {
	g.Config().SetFileName("config.toml")
	go nrpc_server.StartServer()
}
func main() {

	g.Server().Run()
}
