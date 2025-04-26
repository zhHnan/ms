package router

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"hnz.com/ms_serve/ms-common/discovery"
	"hnz.com/ms_serve/ms-common/logs"
	"hnz.com/ms_serve/ms-grpc/user/login"
	"hnz.com/ms_serve/ms-user/internal/interceptor"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"hnz.com/ms_serve/ms-user/config"
	login_service "hnz.com/ms_serve/ms-user/pkg/service/login_service.v1"
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
			//loginServiceV1.RegisterLoginServiceServer(g, &loginServiceV1.LoginService{})
			login.RegisterLoginServiceServer(g, login_service.New())
		}}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			otelgrpc.UnaryServerInterceptor(),
			interceptor.NewInterceptor().CacheInterceptor(),
		)))
	//s := grpc.NewServer()
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

func RegisterEtcdServer() {
	etcdRegister := discovery.NewResolver(config.Cfg.Ec.Addrs, logs.Lg)
	resolver.Register(etcdRegister)
	info := discovery.Server{
		Name:    config.Cfg.Gc.Name,
		Addr:    config.Cfg.Gc.Addr,
		Version: config.Cfg.Gc.Version,
		Weight:  config.Cfg.Gc.Weight,
	}
	r := discovery.NewRegister(config.Cfg.Ec.Addrs, logs.Lg)
	_, err := r.Register(info, 2)
	if err != nil {
		log.Fatalln(err)
	}
}
