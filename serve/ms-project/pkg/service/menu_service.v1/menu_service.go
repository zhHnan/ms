package project_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"hnz.com/ms_serve/ms-common/errs"
	menuRpc "hnz.com/ms_serve/ms-grpc/menu"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/database/tran"
	"hnz.com/ms_serve/ms-project/internal/domain"
	"hnz.com/ms_serve/ms-project/internal/repo"
)

type MenuService struct {
	menuRpc.UnimplementedMenuServiceServer
	cache       repo.Cache
	transaction tran.Transaction
	menuDomain  *domain.MenuDomain
}

// New 初始化
func New() *MenuService {
	return &MenuService{
		cache:       dao.Rc,
		transaction: dao.NewTrans(),
		menuDomain:  domain.NewMenuDomain(),
	}
}

func (m *MenuService) MenuList(context.Context, *menuRpc.MenuReqMessage) (*menuRpc.MenuResponseMessage, error) {
	treeList, err := m.menuDomain.MenuTreeList()
	if err != nil {
		zap.L().Error("MenuList error", zap.Error(err))
		return nil, errs.GrpcError(err)
	}
	var list []*menuRpc.MenuMessage
	_ = copier.Copy(&list, treeList)
	return &menuRpc.MenuResponseMessage{List: list}, nil
}
