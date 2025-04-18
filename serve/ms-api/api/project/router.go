package project

import (
	"github.com/gin-gonic/gin"
	"hnz.com/ms_serve/ms-api/router"
	"log"
)

type RouterUser struct {
}

func init() {
	log.Println("init user-api routers...")
	ru := &RouterUser{}
	router.Register(ru)
}

func (*RouterUser) Route(r *gin.Engine) {
	InitProjectRpc()
	h := New()
	r.POST("/project/index", h.index)
}
