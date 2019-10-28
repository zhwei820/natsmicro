package main

import (
	_ "natsmicro/helloworld/boot"
	_ "natsmicro/helloworld/router"
	"github.com/gogf/gf/frame/g"
)

func init()  {
	g.Config().SetFileName("config.toml")

}
func main() {

	g.Server().Run()
}
