package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var Cfg = InitConfig()

type Config struct {
	viper *viper.Viper
	Sc    *ServerConfig
}
type ServerConfig struct {
	Name string
	Addr string
}

// InitConfig 初始化配置
func InitConfig() *Config {
	conf := &Config{
		viper: viper.New(),
	}
	dir, _ := os.Getwd()
	conf.viper.SetConfigName("app")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath("/etc/ms/user")
	conf.viper.AddConfigPath(dir + "/config")
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln("读取配置文件失败！", err)
	}
	conf.ReadServerConfig()
	return conf
}

func (cfg *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = cfg.viper.GetString("server.name")
	sc.Addr = cfg.viper.GetString("server.addr")
	cfg.Sc = sc
}
