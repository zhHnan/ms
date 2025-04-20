package rpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"hnz.com/ms_serve/ms-common/discovery"
	"hnz.com/ms_serve/ms-common/logs"
	"hnz.com/ms_serve/ms-grpc/project"
	"hnz.com/ms_serve/ms-grpc/task"
	"hnz.com/ms_serve/ms-user/config"
	"log"
)

var ProjectClient project.ProjectServiceClient
var TaskClient task_service_v1.TaskServiceClient

func InitProjectRpc() {
	etcdResolver := discovery.NewResolver(config.Cfg.Ec.Addrs, logs.Lg)
	resolver.Register(etcdResolver)
	conn, err := grpc.Dial("etcd:///project", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ProjectClient = project.NewProjectServiceClient(conn)
	TaskClient = task_service_v1.NewTaskServiceClient(conn)
}
