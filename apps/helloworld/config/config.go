package config

import (
	"github.com/gogf/gf/os/glog"
	"natsmicro/common/conf"
	"os"
	"strconv"
)

func init() {
	// 可以没有，采用默认值
	nacusIp := os.Getenv("NacusIp")
	if nacusIp == "" {
		nacusIp = "127.0.0.1"
	}
	nacusPort := os.Getenv("NacusPort")
	if nacusPort == "" {
		nacusPort = "8848"
	}
	intNacusPort, _ := strconv.Atoi(nacusPort)
	dataId := "natsmicro_dev"
	group := "natsmicro"

	glog.Info("InitConfig: ", nacusIp, uint64(intNacusPort), dataId, group)
	conf.InitConfig(nacusIp, uint64(intNacusPort), dataId, group)
}
