package project_service_v1

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/errs"
	authRpc "hnz.com/ms_serve/ms-grpc/auth"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/database"
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
	if msg.Action == "save" {
		//保存
		nodes := msg.Nodes
		fmt.Println("nodes:", nodes)
		//先删在存 加事务
		authId := msg.AuthId
		err := a.transaction.Action(func(conn database.DBConn) error {
			err := a.projectAuthDomain.Save(conn, authId, nodes)
			return err
		})
		if err != nil {
			// 检查错误类型，避免nil指针异常
			berr, ok := err.(*errs.BError)
			if !ok || berr == nil {
				if err == nil {
					return &authRpc.ApplyResponse{}, nil
				}
				return nil, err
			}
			return nil, errs.GrpcError(berr)
		}
	}
	return &authRpc.ApplyResponse{}, nil
}

func (a *AuthService) AuthNodesByMemberId(ctx context.Context, msg *authRpc.AuthReqMessage) (*authRpc.AuthNodesResponse, error) {
	list, err := a.projectAuthDomain.AuthNodes(msg.MemberId)
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	return &authRpc.AuthNodesResponse{List: list}, nil
}
