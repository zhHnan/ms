package main

import (
	"github.com/gin-gonic/gin"
	srv "hnz.com/ms_serve/common"
	_ "hnz.com/ms_serve/user/api"
	"hnz.com/ms_serve/user/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	srv.Run(r, ":8088", "ms_user")
}
