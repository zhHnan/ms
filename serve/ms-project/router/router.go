package router

import (
	"hnz.com/ms_serve/ms-common/discovery"
	"hnz.com/ms_serve/ms-common/logs"
	account_service "hnz.com/ms_serve/ms-grpc/account"
	auth_service "hnz.com/ms_serve/ms-grpc/auth"
	department_service "hnz.com/ms_serve/ms-grpc/department"
	"hnz.com/ms_serve/ms-grpc/project"
	task_service "hnz.com/ms_serve/ms-grpc/task"
	"hnz.com/ms_serve/ms-project/internal/interceptor"
	"hnz.com/ms_serve/ms-project/internal/rpc"
	authServiceV1 "hnz.com/ms_serve/ms-project/pkg/service/auth_service.v1"
	departmentServiceV1 "hnz.com/ms_serve/ms-project/pkg/service/department_service.v1"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"hnz.com/ms_serve/ms-project/config"
	accountServiceV1 "hnz.com/ms_serve/ms-project/pkg/service/account_service.v1"
	projectServiceV1 "hnz.com/ms_serve/ms-project/pkg/service/project_service.v1"
	taskServiceV1 "hnz.com/ms_serve/ms-project/pkg/service/task_service.v1"
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
			project.RegisterProjectServiceServer(g, projectServiceV1.New())
			task_service.RegisterTaskServiceServer(g, taskServiceV1.New())
			account_service.RegisterAccountServiceServer(g, accountServiceV1.New())
			department_service.RegisterDepartmentServiceServer(g, departmentServiceV1.New())
			auth_service.RegisterAuthServiceServer(g, authServiceV1.New())
		}}
	s := grpc.NewServer(interceptor.NewInterceptor().CacheInterceptor())
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

func InitUserRpc() {
	rpc.InitUserRpc()
}
