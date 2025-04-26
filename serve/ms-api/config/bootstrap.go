package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var BC = InitBootstrap()

type BootConf struct {
	viper       *viper.Viper
	NacosConfig *NacosConfig
}

func (c *BootConf) ReadNacosConfig() {
	nc := &NacosConfig{}
	c.viper.UnmarshalKey("nacos", nc)
	c.NacosConfig = nc
}

type NacosConfig struct {
	Namespace   string
	Group       string
	IpAddr      string
	Port        int
	ContextPath string
	Scheme      string
}

func InitBootstrap() *BootConf {
	conf := &BootConf{viper: viper.New()}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("bootstrap")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath(workDir + "/config")
	conf.viper.AddConfigPath("E:/Projects/go_proj/ms/serve/ms-api")
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	conf.ReadNacosConfig()
	return conf
}
