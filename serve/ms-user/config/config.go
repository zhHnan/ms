package config

import (
	"bytes"
	"github.com/go-redis/redis"
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
	Gc    *GrpcConfig
	Ec    *EtcdConfig
	Mc    *MysqlConfig
	Jc    *JwtConfig
	Dc    *DbConfig
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
	Name     string
}
type DbConfig struct {
	// 是否开启数据库分离
	Separation bool
	Master     MysqlConfig
	Slave      []MysqlConfig
}

type JwtConfig struct {
	AccessSecret  string
	RefreshSecret string
	AccessExp     int
	RefreshExp    int
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
func (c *Config) ReadJwtConfig() {
	jc := &JwtConfig{
		AccessSecret:  c.viper.GetString("jwt.accessSecret"),
		RefreshSecret: c.viper.GetString("jwt.refreshSecret"),
		AccessExp:     c.viper.GetInt("jwt.accessExp"),
		RefreshExp:    c.viper.GetInt("jwt.refreshExp"),
	}
	c.Jc = jc
}

func (c *Config) InitDbConfig() {
	mc := &DbConfig{}
	mc.Separation = c.viper.GetBool("db.separation")
	var slaves []MysqlConfig
	err := c.viper.UnmarshalKey("db.slave", &slaves)
	if err != nil {
		panic(err)
	}
	master := MysqlConfig{
		Username: c.viper.GetString("db.master.username"),
		Password: c.viper.GetString("db.master.password"),
		Host:     c.viper.GetString("db.master.host"),
		Port:     c.viper.GetInt("db.master.port"),
		Db:       c.viper.GetString("db.master.db"),
	}
	mc.Master = master
	mc.Slave = slaves
	c.Dc = mc
}
func (c *Config) ReLoadAllConfig() {
	c.ReadServerConfig()
	c.InitZapLog()
	c.ReadGrpcConfig()
	c.ReadEtcdConfig()
	c.InitMysqlConfig()
	c.ReadJwtConfig()
	c.InitDbConfig()
	//重新创建相关的客户端
	c.ReConnRedis()
	c.ReConnMysql()
}
