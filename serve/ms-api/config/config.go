package config

import (
	"bytes"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"hnz.com/ms_serve/ms-common/logs"
	"log"
	"os"
)

var Cfg = InitConfig()

type Config struct {
	viper *viper.Viper
	Sc    *ServerConfig
	Ec    *EtcdConfig
}
type ServerConfig struct {
	Name string
	Addr string
}
type EtcdConfig struct {
	Addrs []string
}

// InitConfig 初始化配置
func InitConfig() *Config {
	conf := &Config{
		viper: viper.New(),
	}
	// 先从nacos 读取配置
	nacosClient := InitNacosClient()
	configYaml, err2 := nacosClient.configClient.GetConfig(vo.ConfigParam{DataId: "config.yaml", Group: nacosClient.Group})
	if err2 != nil {
		log.Fatalln("读取配置文件失败！", err2)
	}
	conf.viper.SetConfigType("yaml")
	if configYaml != "" {
		err := conf.viper.ReadConfig(bytes.NewBuffer([]byte(configYaml)))
		if err != nil {
			log.Fatalln("读取配置文件失败！", err)
		}
		// 监听配置变化
		err2 = nacosClient.configClient.ListenConfig(vo.ConfigParam{
			DataId: "config.yaml",
			Group:  nacosClient.Group,
			OnChange: func(namespace, group, dataId, data string) {
				log.Printf("config change content 【%s】\n", data)
				err := conf.viper.ReadConfig(bytes.NewBuffer([]byte(data)))
				if err != nil {
					log.Printf("load nacos config change err！【%s】\n", err.Error())
				}
				// 所有配置发生变化时，重新读取配置文件
				conf.ReLoadAllConfig()
			},
		})
		if err2 != nil {
			log.Fatalln("监听配置文件失败！", err2)
		}
	} else {
		dir, _ := os.Getwd()
		conf.viper.SetConfigName("config")
		conf.viper.SetConfigType("yaml")
		conf.viper.AddConfigPath("/etc/ms/ms-user")
		conf.viper.AddConfigPath(dir + "/config")
		err := conf.viper.ReadInConfig()
		if err != nil {
			log.Fatalln("读取配置文件失败！", err)
		}
	}
	conf.ReLoadAllConfig()
	return conf
}

func (cfg *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = cfg.viper.GetString("server.name")
	sc.Addr = cfg.viper.GetString("server.addr")
	cfg.Sc = sc
}
func (c *Config) InitZapLog() {
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("maxSize"),
		MaxAge:        c.viper.GetInt("maxAge"),
		MaxBackups:    c.viper.GetInt("maxBackups"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}

func (cfg *Config) ReadEtcdConfig() {
	ec := &EtcdConfig{}
	var addrs []string
	err := cfg.viper.UnmarshalKey("etcd.addrs", &addrs)
	if err != nil {
		log.Fatalln("读取etcd配置失败！", err)
	}
	ec.Addrs = addrs
	cfg.Ec = ec
}

func (c *Config) ReLoadAllConfig() {
	c.ReadServerConfig()
	c.InitZapLog()
	c.ReadEtcdConfig()
}
