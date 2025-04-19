package project

import (
	"github.com/gin-gonic/gin"
	"hnz.com/ms_serve/ms-api/api/midd"
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
	group := r.Group("/project/")
	group.Use(midd.TokenVerify())
	group.POST("index", h.index)
	group.POST("project/selfList", h.projectList)
}
