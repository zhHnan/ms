package config

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
)

type NacosClient struct {
	configClient config_client.IConfigClient
	Group        string
}

func InitNacosClient() *NacosClient {
	bootConf := InitBootstrap()
	clientConfig := constant.ClientConfig{
		NamespaceId:         bootConf.NacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      bootConf.NacosConfig.IpAddr,
			Port:        uint64(bootConf.NacosConfig.Port),
			ContextPath: bootConf.NacosConfig.ContextPath,
			Scheme:      bootConf.NacosConfig.Scheme,
		},
	}
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
	return &NacosClient{
		configClient: configClient,
		Group:        bootConf.NacosConfig.Group,
	}
}
