package main

import (
	"github.com/gin-gonic/gin"
	srv "hnz.com/ms_serve/common"
	_ "hnz.com/ms_serve/ms-api/api"
	"hnz.com/ms_serve/ms-api/config"
	"hnz.com/ms_serve/ms-api/router"
)

func main() {
	r := gin.Default()
	//从配置中读取日志配置，初始化日志
	config.Cfg.InitZapLog()
	router.InitRouter(r)

	srv.Run(r, config.Cfg.Sc.Addr, config.Cfg.Sc.Name, nil)
}
