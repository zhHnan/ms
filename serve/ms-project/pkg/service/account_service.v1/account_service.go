package project_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/errs"
	accountRpc "hnz.com/ms_serve/ms-grpc/account"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/database/tran"
	"hnz.com/ms_serve/ms-project/internal/domain"
	"hnz.com/ms_serve/ms-project/internal/repo"
)

type AccountService struct {
	accountRpc.UnimplementedAccountServiceServer
	cache             repo.Cache
	transaction       tran.Transaction
	accountDomain     *domain.AccountDomain
	projectAuthDomain *domain.ProjectAuthDomain
}

// New 初始化
func New() *AccountService {
	return &AccountService{
		cache:             dao.Rc,
		transaction:       dao.NewTrans(),
		accountDomain:     domain.NewAccountDomain(),
		projectAuthDomain: domain.NewProjectAuthDomain(),
	}
}

func (a *AccountService) Account(c context.Context, msg *accountRpc.AccountReqMessage) (*accountRpc.AccountResponse, error) {
	accountList, total, err := a.accountDomain.AccountList(
		msg.OrganizationCode,
		msg.MemberId,
		msg.Page,
		msg.PageSize,
		msg.DepartmentCode,
		msg.SearchType)
	if err != nil {
		return &accountRpc.AccountResponse{}, errs.GrpcError(err)
	}
	authList, err := a.projectAuthDomain.AuthList(encrypts.DecryptToRes(msg.OrganizationCode))
	if err != nil {
		return &accountRpc.AccountResponse{}, errs.GrpcError(err)
	}
	var maList []*accountRpc.MemberAccount
	_ = copier.Copy(&maList, accountList)
	var prList []*accountRpc.ProjectAuth
	_ = copier.Copy(&prList, authList)
	return &accountRpc.AccountResponse{
		AccountList: maList,
		AuthList:    prList,
		Total:       total,
	}, nil
}
