package main

import (
	"github.com/gin-gonic/gin"
	_ "hnz.com/ms_serve/ms-api/api"
	"hnz.com/ms_serve/ms-api/api/midd"
	"hnz.com/ms_serve/ms-api/config"
	"hnz.com/ms_serve/ms-api/router"
	srv "hnz.com/ms_serve/ms-common"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(midd.RequestLog())
	r.StaticFS("/upload", http.Dir("./upload"))
	//从配置中读取日志配置，初始化日志
	config.Cfg.InitZapLog()
	router.InitRouter(r)

	srv.Run(r, config.Cfg.Sc.Addr, config.Cfg.Sc.Name, nil)
}
