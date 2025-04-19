package user

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
	rpc.InitUserRpc()
	h := New()
	r.POST("/project/login/getCaptcha", h.getCaptcha)
	r.POST("/project/login/register", h.register)
	r.POST("/project/login", h.login)
	group := r.Group("/project/")
	group.Use(midd.TokenVerify())
	group.POST("organization/_getOrgList", h.myOrgList)
}
