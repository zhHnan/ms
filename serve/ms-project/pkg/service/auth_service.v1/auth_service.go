package project_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/errs"
	authRpc "hnz.com/ms_serve/ms-grpc/auth"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/database/tran"
	"hnz.com/ms_serve/ms-project/internal/domain"
	"hnz.com/ms_serve/ms-project/internal/repo"
)

type AuthService struct {
	authRpc.UnimplementedAuthServiceServer
	cache             repo.Cache
	transaction       tran.Transaction
	projectAuthDomain *domain.ProjectAuthDomain
}

func New() *AuthService {
	return &AuthService{
		cache:             dao.Rc,
		transaction:       dao.NewTrans(),
		projectAuthDomain: domain.NewProjectAuthDomain(),
	}
}

func (a *AuthService) AuthList(ctx context.Context, msg *authRpc.AuthReqMessage) (*authRpc.ListAuthMessage, error) {
	organizationCode := encrypts.DecryptToRes(msg.OrganizationCode)
	listPage, total, err := a.projectAuthDomain.AuthListPage(organizationCode, msg.Page, msg.PageSize)
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	var prList []*authRpc.ProjectAuth
	_ = copier.Copy(&prList, listPage)
	return &authRpc.ListAuthMessage{List: prList, Total: total}, nil
}

func (a *AuthService) Apply(ctx context.Context, msg *authRpc.AuthReqMessage) (*authRpc.ApplyResponse, error) {
	if msg.Action == "getnode" {
		//获取列表
		list, checkedList, err := a.projectAuthDomain.AllNodeAndAuth(msg.AuthId)
		if err != nil {
			return nil, errs.GrpcError(err)
		}
		var prList []*authRpc.ProjectNodeMessage
		_ = copier.Copy(&prList, list)
		return &authRpc.ApplyResponse{List: prList, CheckedList: checkedList}, nil
	}
	return &authRpc.ApplyResponse{}, nil
}
