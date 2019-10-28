package boot

import "github.com/gogf/gf/frame/g"

func init() {
    //g.Server().SetPort(8199)
    g.Server().SetPort(g.Config().GetInt("port", 8199))

}

