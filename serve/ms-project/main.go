package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	srv "hnz.com/ms_serve/ms-common"
	"hnz.com/ms_serve/ms-project/config"
	"hnz.com/ms_serve/ms-project/router"
)

func main() {
	r := gin.Default()
	//从配置中读取日志配置，初始化日志
	router.InitRouter(r)
	grpc := router.RegisterGrpc()
	router.RegisterEtcdServer()
	stop := func() {
		grpc.Stop()
	}
	// 初始化rpc
	router.InitUserRpc()
	// 开启pprof
	pprof.Register(r)
	srv.Run(r, config.Cfg.Sc.Addr, config.Cfg.Sc.Name, stop)
}
