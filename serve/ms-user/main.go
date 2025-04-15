package main

import (
	"github.com/gin-gonic/gin"
	srv "hnz.com/ms_serve/ms-common"
	_ "hnz.com/ms_serve/user/api"
	"hnz.com/ms_serve/user/config"
	"hnz.com/ms_serve/user/router"
)

func main() {
	r := gin.Default()
	//从配置中读取日志配置，初始化日志
	config.Cfg.InitZapLog()
	router.InitRouter(r)
	grpc := router.RegisterGrpc()
	stop := func() {
		grpc.Stop()
	}
	srv.Run(r, config.Cfg.Sc.Addr, config.Cfg.Sc.Name, stop)
}
