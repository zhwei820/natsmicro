package main

import (
	_ "natsmicro/helloworld/boot"
	_ "natsmicro/helloworld/router"
	"github.com/gogf/gf/frame/g"
	"natsmicro/helloworld/rpc/nrpc_server"
)

func init()  {
	g.Config().SetFileName("config.toml")
	go nrpc_server.StartServer()
}
func main() {

	g.Server().Run()
}
