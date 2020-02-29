package router

import (
	"github.com/gogf/gf/frame/g"
	"natsmicro/apps/helloworld2/app/api/hello"
)

// 统一路由注册.
func init() {
	g.Server().BindHandler("/", hello.Handler)
}
