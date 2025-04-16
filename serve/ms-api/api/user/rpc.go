package user

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	loginServiceV1 "hnz.com/ms_serve/ms-user/pkg/service/login_service.v1"
	"log"
)

var UserClient loginServiceV1.LoginServiceClient

func InitUserRpc() {

	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	UserClient = loginServiceV1.NewLoginServiceClient(conn)
}
