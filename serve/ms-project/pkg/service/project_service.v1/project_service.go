package project_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-common/times"
	"hnz.com/ms_serve/ms-grpc/project"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/data/menu"
	pro "hnz.com/ms_serve/ms-project/internal/data/project"
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
	page := in.Page
	pageSize := in.PageSize
	var pm []*pro.ProjectAndMember
	var total int64
	var err error
	if in.SelectBy == "" || in.SelectBy == "my" {
		pm, total, err = p.projectRepo.FindProjectByMemId(ctx, id, "", page, pageSize)
	}
	if in.SelectBy == "archive" {
		pm, total, err = p.projectRepo.FindProjectByMemId(ctx, id, "and archive = 1", page, pageSize)
	}
	if in.SelectBy == "deleted" {
		pm, total, err = p.projectRepo.FindProjectByMemId(ctx, id, "and deleted = 1", page, pageSize)
	}
	if in.SelectBy == "collect" {
		pm, total, err = p.projectRepo.FindCollectProjectByMemId(ctx, id, page, pageSize)
	}
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
		pam := pro.ToMap(pm)[v.Id]
		v.AccessControlType = pam.GetAccessControlType()
		v.OrganizationCode, _ = encrypts.Encrypt(strconv.FormatInt(pam.OrganizationCode, 10), model.AESKey)
		v.JoinTime = times.FormatByMill(pam.JoinTime)
		v.OwnerName = in.MemberName
		v.Order = int32(pam.Sort)
		v.CreateTime = times.FormatByMill(pam.CreateTime)
	}
	return &project.MyProjectResponse{Pm: pmm, Total: total}, nil
}
