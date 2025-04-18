package project_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-grpc/project"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/data/menu"
	"hnz.com/ms_serve/ms-project/internal/database/tran"
	"hnz.com/ms_serve/ms-project/internal/repo"
	"hnz.com/ms_serve/ms-project/pkg/model"
)

type ProjectService struct {
	project.UnimplementedProjectServiceServer
	cache       repo.Cache
	transaction tran.Transaction
	menuRepo    repo.MenuRepo
}

func New() *ProjectService {
	return &ProjectService{
		cache:       dao.Rc,
		transaction: dao.NewTrans(),
		menuRepo:    dao.NewMenuDao(),
	}
}

func (p *ProjectService) Index(ctx context.Context, in *project.IndexMessage) (*project.IndexResponse, error) {
	pms, err := p.menuRepo.FindAll(context.Background())
	if err != nil {
		zap.L().Error("index database findMenus error！", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	child := menu.CovertChild(pms)
	var menuMessage []*project.MenuMessage
	_ = copier.Copy(&menuMessage, child)
	return &project.IndexResponse{
		Menus: menuMessage,
	}, nil
}
