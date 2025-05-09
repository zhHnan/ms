package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Router interface {
	Route(r *gin.Engine)
}
type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}
func (*RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

var routers []Router

func InitRouter(r *gin.Engine) {
	for _, ro := range routers {
		ro.Route(r)
	}
}

func Register(ros ...Router) {
	routers = append(routers, ros...)
}

type grpcConfig struct {
	Addr         string
	RegisterFunc func(server *grpc.Server)
}
