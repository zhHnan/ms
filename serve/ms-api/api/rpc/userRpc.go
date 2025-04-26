package rpc

import (
	"log"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"hnz.com/ms_serve/ms-common/discovery"
	"hnz.com/ms_serve/ms-common/logs"
	"hnz.com/ms_serve/ms-grpc/user/login"
	"hnz.com/ms_serve/ms-user/config"
)

var UserClient login.LoginServiceClient

func InitUserRpc() {
	etcdResolver := discovery.NewResolver(config.Cfg.Ec.Addrs, logs.Lg)
	resolver.Register(etcdResolver)
	conn, err := grpc.NewClient("etcd:///user",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	UserClient = login.NewLoginServiceClient(conn)
}
