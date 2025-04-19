package project_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-grpc/project"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/data/menu"
	"hnz.com/ms_serve/ms-project/internal/database/tran"
	"hnz.com/ms_serve/ms-project/internal/repo"
	"hnz.com/ms_serve/ms-project/pkg/model"
	"strconv"
)

type ProjectService struct {
	project.UnimplementedProjectServiceServer
	cache       repo.Cache
	transaction tran.Transaction
	menuRepo    repo.MenuRepo
	projectRepo repo.ProjectRepo
}

func New() *ProjectService {
	return &ProjectService{
		cache:       dao.Rc,
		transaction: dao.NewTrans(),
		menuRepo:    dao.NewMenuDao(),
		projectRepo: dao.NewProjectDao(),
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
func (p *ProjectService) FindProjectByMemId(ctx context.Context, in *project.ProjectRpcMessage) (*project.MyProjectResponse, error) {
	id := in.MemberId
	pm, total, err := p.projectRepo.FindProjectByMemId(ctx, id, in.Page, in.PageSize)
	if err != nil {
		zap.L().Error("menu findAll error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	if pm == nil {
		return &project.MyProjectResponse{Pm: []*project.ProjectMessage{}, Total: total}, nil
	}
	var pmm []*project.ProjectMessage
	copier.Copy(&pmm, pm)
	for _, v := range pmm {
		v.Code, _ = encrypts.Encrypt(strconv.FormatInt(v.Id, 10), model.AESKey)
	}
	return &project.MyProjectResponse{Pm: pmm, Total: total}, nil
}
