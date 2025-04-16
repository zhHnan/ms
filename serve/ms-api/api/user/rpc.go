package user

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"hnz.com/ms_serve/common/discovery"
	"hnz.com/ms_serve/common/logs"
	"hnz.com/ms_serve/ms-user/config"
	loginServiceV1 "hnz.com/ms_serve/ms-user/pkg/service/login_service.v1"
	"log"
)

var UserClient loginServiceV1.LoginServiceClient

func InitUserRpc() {
	etcdResolver := discovery.NewResolver(config.Cfg.Ec.Addrs, logs.Lg)
	resolver.Register(etcdResolver)
	conn, err := grpc.Dial("etcd:///user", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	UserClient = loginServiceV1.NewLoginServiceClient(conn)
}
