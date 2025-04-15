package main

import (
	"github.com/gin-gonic/gin"
	srv "hnz.com/ms_serve/common"
	"hnz.com/ms_serve/common/logs"
	_ "hnz.com/ms_serve/user/api"
	"hnz.com/ms_serve/user/config"
	"hnz.com/ms_serve/user/router"
	"log"
)

func main() {
	r := gin.Default()
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: "E:\\Projects\\go_proj\\ms\\logs\\project-debug.log",
		InfoFileName:  "E:\\Projects\\go_proj\\ms\\logs\\info\\project-info.log",
		WarnFileName:  "E:\\Projects\\go_proj\\ms\\logs\\error\\project-error.log",
		MaxSize:       500,
		MaxAge:        28,
		MaxBackups:    3,
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
	router.InitRouter(r)
	srv.Run(r, config.Cfg.Sc.Addr, config.Cfg.Sc.Name)
}
