package config

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"hnz.com/ms_serve/ms-user/internal/database/gorms"
)

var _db *gorm.DB

func (c *Config) ReConnMysql() {
	if c.Dc.Separation {
		//读写分离配置
		username := c.Dc.Master.Username //账号
		password := c.Dc.Master.Password //密码
		host := c.Dc.Master.Host         //数据库地址，可以是Ip或者域名
		port := c.Dc.Master.Port         //数据库端口
		Dbname := c.Dc.Master.Db         //数据库名
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
		var err error
		_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			zap.L().Error("连接数据库失败", zap.Error(err))
			return
		}
		var replicas []gorm.Dialector
		for _, v := range c.Dc.Slave {
			username := v.Username //账号
			password := v.Password //密码
			host := v.Host         //数据库地址，可以是Ip或者域名
			port := v.Port         //数据库端口
			Dbname := v.Db         //数据库名
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
			cfg := mysql.Config{
				DSN: dsn,
			}
			replicas = append(replicas, mysql.New(cfg))
		}
		err = _db.Use(dbresolver.Register(dbresolver.Config{
			Sources: []gorm.Dialector{mysql.New(mysql.Config{
				DSN: dsn,
			})},
			Replicas: replicas,
			Policy:   dbresolver.RandomPolicy{},
		}).
			SetMaxIdleConns(10).
			SetMaxOpenConns(200))
		if err != nil {
			zap.L().Error("use slave error", zap.Error(err))
			return
		}
	} else {
		//配置MySQL连接参数
		username := c.Mc.Username //账号
		password := c.Mc.Password //密码
		host := c.Mc.Host         //数据库地址，可以是Ip或者域名
		port := c.Mc.Port         //数据库端口
		Dbname := c.Mc.Db         //数据库名
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
		var err error
		_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			panic("连接数据库失败, error=" + err.Error())
		}
	}
	gorms.SetDB(_db)
}
