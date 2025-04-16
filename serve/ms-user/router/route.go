package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"hnz.com/ms_serve/user/config"
	loginServiceV1 "hnz.com/ms_serve/user/pkg/service/login_service.v1"
	"log"
	"net"
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

func RegisterGrpc() *grpc.Server {
	c := grpcConfig{
		Addr: config.Cfg.Gc.Addr,
		RegisterFunc: func(g *grpc.Server) {
			loginServiceV1.RegisterLoginServiceServer(g, &loginServiceV1.LoginService{})
		}}
	s := grpc.NewServer()
	c.RegisterFunc(s)

	listen, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Println("grpc server start fail...", err)
	}
	go func() {
		err = s.Serve(listen)
		if err != nil {
			log.Println("grpc server start fail...", err)
			return
		}
	}()
	return s
}
