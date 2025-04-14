package user

import (
	"github.com/gin-gonic/gin"
	"hnz.com/ms_serve/user/router"
	"log"
)

func init() {
	log.Println("init user routers...")
	router.Register(&RouterUser{})
}

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	h := &HandlerLogin{}
	r.POST("/api/project/login/getCaptcha", h.GetCaptcha)
}
