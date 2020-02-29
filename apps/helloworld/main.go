package main

import (
	_ "natsmicro/apps/helloworld/boot"
	_ "natsmicro/apps/helloworld/config"
	_ "natsmicro/apps/helloworld/router"
	"github.com/gogf/gf/frame/g"
	"natsmicro/apps/helloworld/rpc/nrpc_server"
)

func init()  {
	//g.Config().SetFileName("config.toml")
	go nrpc_server.StartServer()
}
func main() {

	g.Server().Run()
}
