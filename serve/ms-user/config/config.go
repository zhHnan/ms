package config

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"hnz.com/ms_serve/common/logs"
	"log"
	"os"
)

var Cfg = InitConfig()

type Config struct {
	viper *viper.Viper
	Sc    *ServerConfig
	Gc    *GrpcConfig
}
type ServerConfig struct {
	Name string
	Addr string
}
type GrpcConfig struct {
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
	conf.viper.AddConfigPath("/etc/ms/ms-user")
	conf.viper.AddConfigPath(dir + "/config")
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln("读取配置文件失败！", err)
	}
	conf.ReadServerConfig()
	conf.InitZapLog()
	conf.InitRedisOptions()
	conf.ReadGrpcConfig()
	return conf
}

func (cfg *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = cfg.viper.GetString("server.name")
	sc.Addr = cfg.viper.GetString("server.addr")
	cfg.Sc = sc
}
func (cfg *Config) ReadGrpcConfig() {
	gc := &GrpcConfig{}
	gc.Name = cfg.viper.GetString("grpc.name")
	gc.Addr = cfg.viper.GetString("grpc.addr")
	cfg.Gc = gc
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
func (c *Config) InitRedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     c.viper.GetString("redis.host") + ":" + c.viper.GetString("redis.port"),
		Password: c.viper.GetString("redis.password"), // no password set
		DB:       c.viper.GetInt("db"),                // use default DB
	}
}
