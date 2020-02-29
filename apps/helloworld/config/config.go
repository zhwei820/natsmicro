package config

import (
	"natsmicro/common/conf"
	"os"
	"strconv"
)

func init() {
	// 可以没有，采用默认值
	nacusIp:=os.Getenv("NacusIp")
	nacusPort:=os.Getenv("NacusPort")
	intNacusPort, _ := strconv.Atoi(nacusPort)
	dataId := "natsmicro_dev"
	group := "natsmicro"

	conf.InitConfig(nacusIp, uint64(intNacusPort), dataId, group)
}
