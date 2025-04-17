package config

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"hnz.com/ms_serve/ms-common/logs"
	"log"
	"os"
)

var Cfg = InitConfig()

type Config struct {
	viper *viper.Viper
	Sc    *ServerConfig
	Gc    *GrpcConfig
	Ec    *EtcdConfig
	Mc    *MysqlConfig
}
type ServerConfig struct {
	Name string
	Addr string
}
type GrpcConfig struct {
	Name    string
	Addr    string
	Version string
	Weight  int64
}
type EtcdConfig struct {
	Addrs []string
}
type MysqlConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Db       string
}

// InitConfig 初始化配置
func InitConfig() *Config {
	conf := &Config{
		viper: viper.New(),
	}
	dir, _ := os.Getwd()
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath("/etc/ms/ms-user")
	conf.viper.AddConfigPath(dir + "/config")
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln("读取配置文件失败！", err)
	}
	conf.ReadServerConfig()
	conf.InitZapLog()
	//conf.InitRedisOptions()
	conf.ReadGrpcConfig()
	conf.ReadEtcdConfig()
	conf.InitMysqlConfig()
	return conf
}

// ReadServerConfig 初始化server配置
func (cfg *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = cfg.viper.GetString("server.name")
	sc.Addr = cfg.viper.GetString("server.addr")
	cfg.Sc = sc
}

// ReadGrpcConfig 初始化grpc配置
func (cfg *Config) ReadGrpcConfig() {
	gc := &GrpcConfig{}
	gc.Name = cfg.viper.GetString("grpc.name")
	gc.Addr = cfg.viper.GetString("grpc.addr")
	gc.Version = cfg.viper.GetString("grpc.version")
	gc.Weight = cfg.viper.GetInt64("grpc.weight")
	cfg.Gc = gc
}

// InitZapLog 初始化日志
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

// ReadRedisConfig 初始化redis配置
func (c *Config) ReadRedisConfig() *redis.Options {
	return &redis.Options{
		Addr:     c.viper.GetString("redis.host") + ":" + c.viper.GetString("redis.port"),
		Password: c.viper.GetString("redis.password"), // no password set
		DB:       c.viper.GetInt("db"),                // use default DB
	}
}

// ReadEtcdConfig 初始化etcd配置
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

// InitMysqlConfig 初始化mysql配置
func (cfg *Config) InitMysqlConfig() {
	mc := &MysqlConfig{
		Username: cfg.viper.GetString("mysql.username"),
		Password: cfg.viper.GetString("mysql.password"),
		Host:     cfg.viper.GetString("mysql.host"),
		Port:     cfg.viper.GetInt("mysql.port"),
		Db:       cfg.viper.GetString("mysql.db"),
	}
	cfg.Mc = mc
}
