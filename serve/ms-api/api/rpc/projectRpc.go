package rpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"hnz.com/ms_serve/ms-common/discovery"
	"hnz.com/ms_serve/ms-common/logs"
	account_service_v1 "hnz.com/ms_serve/ms-grpc/account"
	auth_service_v1 "hnz.com/ms_serve/ms-grpc/auth"
	department_service_v1 "hnz.com/ms_serve/ms-grpc/department"
	menu_service_v1 "hnz.com/ms_serve/ms-grpc/menu"
	"hnz.com/ms_serve/ms-grpc/project"
	"hnz.com/ms_serve/ms-grpc/task"
	"hnz.com/ms_serve/ms-user/config"
	"log"
)

var ProjectClient project.ProjectServiceClient
var TaskClient task_service_v1.TaskServiceClient
var AccountClient account_service_v1.AccountServiceClient
var DepartmentClient department_service_v1.DepartmentServiceClient
var AuthClient auth_service_v1.AuthServiceClient
var MenuClient menu_service_v1.MenuServiceClient

func InitProjectRpc() {
	etcdResolver := discovery.NewResolver(config.Cfg.Ec.Addrs, logs.Lg)
	resolver.Register(etcdResolver)
	conn, err := grpc.Dial("etcd:///project", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ProjectClient = project.NewProjectServiceClient(conn)
	TaskClient = task_service_v1.NewTaskServiceClient(conn)
	AccountClient = account_service_v1.NewAccountServiceClient(conn)
	DepartmentClient = department_service_v1.NewDepartmentServiceClient(conn)
	AuthClient = auth_service_v1.NewAuthServiceClient(conn)
	MenuClient = menu_service_v1.NewMenuServiceClient(conn)
}
