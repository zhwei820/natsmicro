package conf

import (
	"github.com/gogf/gf/os/gcfg"
	"github.com/gogf/gf/os/glog"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func InitConfig(nacusIp string, nacusPort uint64, dataId, group string) {
	// 可以没有，采用默认值
	clientConfig := constant.ClientConfig{
		TimeoutMs:      1 * 1000,
		ListenInterval: 2 * 1000,
		BeatInterval:   1 * 1000,
		LogDir:         "./logs/nacos",
		CacheDir:       "./nacos",
	}

	// 至少一个
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      nacusIp,
			ContextPath: "/nacos",
			Port:        nacusPort,
		},
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		glog.Error(err)
	}
	res, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})

	if err != nil {
		glog.Error(err)
	}
	glog.Info(res)
	gcfg.SetContent(res)
	configClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			glog.Info("group:" + group + ", dataId:" + dataId + ", data:" + data)
			gcfg.SetContent(data)

		},
	})
}
