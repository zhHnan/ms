package project_service_v1

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-common/times"
	"hnz.com/ms_serve/ms-grpc/project"
	"hnz.com/ms_serve/ms-grpc/user/login"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/data/menu"
	pro "hnz.com/ms_serve/ms-project/internal/data/project"
	"hnz.com/ms_serve/ms-project/internal/data/task"
	"hnz.com/ms_serve/ms-project/internal/database"
	"hnz.com/ms_serve/ms-project/internal/database/tran"
	"hnz.com/ms_serve/ms-project/internal/repo"
	"hnz.com/ms_serve/ms-project/internal/rpc"
	"hnz.com/ms_serve/ms-project/pkg/model"
	"strconv"
	"time"
)

type ProjectService struct {
	project.UnimplementedProjectServiceServer
	cache           repo.Cache
	transaction     tran.Transaction
	menuRepo        repo.MenuRepo
	projectRepo     repo.ProjectRepo
	proTemplateRepo repo.ProjectTemplateRepo
	taskStagesRepo  repo.TaskStagesTemplateRepo
}

// 初始化
func New() *ProjectService {
	return &ProjectService{
		cache:           dao.Rc,
		transaction:     dao.NewTrans(),
		menuRepo:        dao.NewMenuDao(),
		projectRepo:     dao.NewProjectDao(),
		proTemplateRepo: dao.NewProjectTemplateDao(),
		taskStagesRepo:  dao.NewTaskStagesTemplateDao(),
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
	if in.SelectBy == "collect" {
		pm, total, err = p.projectRepo.FindCollectProjectByMemId(ctx, id, page, pageSize)
		if err != nil {
			zap.L().Error("FindCollectProjectByMemId error", zap.Error(err))
			return nil, errs.GrpcError(model.DataBaseError)
		}
		for _, v := range pm {
			v.Collected = model.Collected
		}
	} else {
		condition := "and deleted = 0"
		if in.SelectBy == "archive" {
			condition = "and archive = 1"
		} else if in.SelectBy == "deleted" {
			condition = "and deleted = 1"
		}
		pm, total, err = p.projectRepo.FindProjectByMemId(ctx, id, condition, page, pageSize)
		if err != nil {
			zap.L().Error("menu FindProjectByMemId error", zap.Error(err))
			return nil, errs.GrpcError(model.DataBaseError)
		}
		// 获取所有收藏项目（不分页）
		collectPms, _, err := p.projectRepo.FindCollectProjectByMemId(ctx, id, 1, 100000)
		if err != nil {
			zap.L().Error("FindCollectProjectByMemId error", zap.Error(err))
			return nil, errs.GrpcError(model.DataBaseError)
		}
		// 构建收藏项目映射表
		collectedMap := make(map[int64]struct{})
		for _, v := range collectPms {
			collectedMap[v.Id] = struct{}{}
		}
		// 标记已收藏项目
		for _, v := range pm {
			if _, ok := collectedMap[v.Id]; ok {
				v.Collected = model.Collected
			}
		}
	}

	if pm == nil {
		return &project.MyProjectResponse{Pm: []*project.ProjectMessage{}, Total: total}, nil
	}

	var pmm []*project.ProjectMessage
	_ = copier.Copy(&pmm, pm)
	for _, v := range pmm {
		v.Code, _ = encrypts.Encrypt(strconv.FormatInt(v.ProjectCode, 10), model.AESKey)
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

func (p *ProjectService) FindProjectTemplate(ctx context.Context, in *project.ProjectRpcMessage) (*project.ProjectTemplateResponse, error) {
	// 根据viewType去查询模板表
	var proTems []pro.ProjectTemplate
	var total int64
	var err error
	page := in.Page
	pageSize := in.PageSize
	orgCodeStr, _ := encrypts.Decrypt(in.OrganizationCode, model.AESKey)
	orgCode, _ := strconv.ParseInt(orgCodeStr, 10, 64)
	switch in.ViewType {
	case -1:
		// -1 查询全部
		proTems, total, err = p.proTemplateRepo.FindProjectTemplateAll(ctx, orgCode, page, pageSize)
	case 0:
		proTems, total, err = p.proTemplateRepo.FindProjectTemplateCustom(ctx, in.MemberId, orgCode, page, pageSize)
	case 1:
		proTems, total, err = p.proTemplateRepo.FindProjectTemplateSystem(ctx, page, pageSize)
	default:
		zap.L().Error("menu findAll error", zap.Error(err))
		return nil, fmt.Errorf("unsupported ViewType: %d", in.ViewType)
	}
	if err != nil {
		zap.L().Error("menu findAll error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	// 模型转换，拿到模版id列表在任务步骤模板表中查询
	tasks, err := p.taskStagesRepo.FindInProTemIds(ctx, pro.ToProjectTemplateIds(proTems))
	if err != nil {
		zap.L().Error("menu convertInProTemIds error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}

	var ptas []*pro.ProjectTemplateAll
	for _, v := range proTems {
		ptas = append(ptas, v.Convert(task.CovertProjectMap(tasks)[v.Id]))
	}
	// 组装数据
	var res []*project.ProjectTemplateMessage
	_ = copier.Copy(&res, ptas)
	return &project.ProjectTemplateResponse{Ptm: res, Total: total}, nil
}

func (p *ProjectService) SaveProject(ctx context.Context, msg *project.ProjectRpcMessage) (*project.SaveProjectMessage, error) {
	organizationCodeStr, _ := encrypts.Decrypt(msg.OrganizationCode, model.AESKey)
	organizationCode, _ := strconv.ParseInt(organizationCodeStr, 10, 64)
	templateCodeStr, _ := encrypts.Decrypt(msg.TemplateCode, model.AESKey)
	templateCode, _ := strconv.ParseInt(templateCodeStr, 10, 64)
	pr := &pro.Project{
		Name:              msg.Name,
		Description:       msg.Description,
		TemplateCode:      int(templateCode),
		CreateTime:        time.Now().UnixMilli(),
		Cover:             "https://img2.baidu.com/it/u=792555388,2449797505&fm=253&fmt=auto&app=138&f=JPEG?w=667&h=500",
		Deleted:           model.NoDeleted,
		Archive:           model.NoArchive,
		OrganizationCode:  organizationCode,
		AccessControlType: model.Open,
		TaskBoardTheme:    model.Simple,
	}
	err := p.transaction.Action(func(conn database.DBConn) error {
		err := p.projectRepo.SaveProject(conn, ctx, pr)
		if err != nil {
			zap.L().Error("project SaveProject error", zap.Error(err))
			return errs.GrpcError(model.DataBaseError)
		}
		pm := &pro.ProjectMember{
			ProjectCode: pr.Id,
			MemberCode:  msg.MemberId,
			JoinTime:    time.Now().UnixMilli(),
			IsOwner:     msg.MemberId,
			Authorize:   "",
		}
		fmt.Printf("save ProjectCode:【%s】/n", pm.ProjectCode)
		err = p.projectRepo.SaveProjectMember(conn, ctx, pm)
		if err != nil {
			zap.L().Error("project SaveProjectMember error", zap.Error(err))
			return errs.GrpcError(model.DataBaseError)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	code, _ := encrypts.Encrypt(strconv.FormatInt(pr.Id, 10), model.AESKey)
	rsp := &project.SaveProjectMessage{
		Id:               pr.Id,
		Code:             code,
		OrganizationCode: organizationCodeStr,
		Name:             pr.Name,
		Cover:            pr.Cover,
		CreateTime:       times.FormatByMill(pr.CreateTime),
		TaskBoardTheme:   pr.TaskBoardTheme,
	}
	return rsp, nil
}

func (p *ProjectService) GetProjectDetail(ctx context.Context, msg *project.ProjectRpcMessage) (*project.ProjectDetailMessage, error) {
	projectCodeStr, _ := encrypts.Decrypt(msg.ProjectCode, model.AESKey)
	projectCode, _ := strconv.ParseInt(projectCodeStr, 10, 64)
	memberId := msg.MemberId
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	projectAndMember, err := p.projectRepo.FindProjectByPIdAndMemId(c, projectCode, memberId)
	if err != nil {
		zap.L().Error("project FindProjectDetail FindProjectByPIdAndMemId error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	ownerId := projectAndMember.IsOwner
	member, err := rpc.UserClient.FindMemberInfoById(c, &login.UserMessage{MemId: ownerId})
	if err != nil {
		zap.L().Error("project rpc FindProjectDetail FindMemberInfoById error", zap.Error(err))
		return nil, err
	}
	// todo 放于redis中
	isCollect, err := p.projectRepo.FindCollectByPidAndMemId(c, projectCode, memberId)
	if err != nil {
		zap.L().Error("project FindProjectDetail FindCollectByPidAndMemId error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	if isCollect {
		projectAndMember.Collected = model.Collected
	}
	var detailMsg = &project.ProjectDetailMessage{}
	_ = copier.Copy(&detailMsg, projectAndMember)
	detailMsg.OwnerAvatar = member.Avatar
	detailMsg.OwnerName = member.Name
	detailMsg.Code, _ = encrypts.EncryptInt64(projectAndMember.Id, model.AESKey)
	detailMsg.AccessControlType = projectAndMember.GetAccessControlType()
	detailMsg.OrganizationCode, _ = encrypts.EncryptInt64(projectAndMember.OrganizationCode, model.AESKey)
	detailMsg.Order = int32(projectAndMember.Sort)
	detailMsg.CreateTime = times.FormatByMill(projectAndMember.CreateTime)
	return detailMsg, nil
}

func (p *ProjectService) UpdateDeletedProject(ctx context.Context, msg *project.ProjectRpcMessage) (*project.ProjectUpdateDeletedResponse, error) {
	projectCodeStr, _ := encrypts.Decrypt(msg.ProjectCode, model.AESKey)
	projectCode, _ := strconv.ParseInt(projectCodeStr, 10, 64)
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := p.projectRepo.UpdateDeletedProject(c, projectCode, msg.Deleted)
	if err != nil {
		zap.L().Error("project UpdateDeletedProject DeleteProject error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	return &project.ProjectUpdateDeletedResponse{}, nil
}
func (p *ProjectService) UpdateCollectProject(ctx context.Context, msg *project.ProjectRpcMessage) (*project.UpdateCollectResponse, error) {
	projectCodeStr, _ := encrypts.Decrypt(msg.ProjectCode, model.AESKey)
	fmt.Println("【rpc server】projectCodeStr:", projectCodeStr)
	projectCode, _ := strconv.ParseInt(projectCodeStr, 10, 64)
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var err error
	if "collect" == msg.CollectType {
		pc := &pro.ProjectCollection{
			ProjectCode: projectCode,
			MemberCode:  msg.MemberId,
			CreateTime:  time.Now().UnixMilli(),
		}
		err = p.projectRepo.SaveProjectCollect(c, pc)
	}
	if "cancel" == msg.CollectType {
		err = p.projectRepo.DeleteProjectCollect(c, msg.MemberId, projectCode)
	}
	if err != nil {
		zap.L().Error("project UpdateCollectProject SaveProjectCollect error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	return &project.UpdateCollectResponse{}, nil
}
