package project

import (
	"github.com/gin-gonic/gin"
	"hnz.com/ms_serve/ms-api/api/midd"
	"hnz.com/ms_serve/ms-api/api/rpc"
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
	rpc.InitProjectRpc()
	h := New()
	group := r.Group("/project")
	group.Use(midd.TokenVerify())
	group.POST("/index", h.index)
	group.POST("/project/selfList", h.projectList)
	group.POST("/project", h.projectList)
	group.POST("/project_template", h.projectTemplate)
	group.POST("/project/save", h.projectSave)
	group.POST("/project/read", h.projectRead)
	group.POST("/project/recycle", h.projectRecycle)
	group.POST("/project/recovery", h.projectRecovery)
}
