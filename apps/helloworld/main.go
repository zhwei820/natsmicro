package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	_ "natsmicro/apps/helloworld/config"
	"natsmicro/apps/helloworld2/rpc/nrpc_client"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "natsmicro/apps/helloworld/boot"
	_ "natsmicro/apps/helloworld/router"
	"natsmicro/apps/helloworld/rpc/nrpc_server"
)

func init() {
	//g.Config().SetFileName("config.toml")
	go nrpc_server.StartServer()
}
func main() {

	//g.Server().Run()
	go func() {
		g.Server().Start()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		glog.Info("get a signal: ", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			g.Server().Shutdown()
			nrpc_server.StopServer()
			nrpc_client.Close()
			glog.Info("kratos-demo exit")
			time.Sleep(time.Second)

			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
