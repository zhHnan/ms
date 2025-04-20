package project_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-common/times"
	task_service "hnz.com/ms_serve/ms-grpc/task"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/data/task"
	"hnz.com/ms_serve/ms-project/internal/database/tran"
	"hnz.com/ms_serve/ms-project/internal/repo"
	"hnz.com/ms_serve/ms-project/pkg/model"
	"time"
)

type TaskService struct {
	task_service.UnimplementedTaskServiceServer
	cache                  repo.Cache
	transaction            tran.Transaction
	projectRepo            repo.ProjectRepo
	proTemplateRepo        repo.ProjectTemplateRepo
	taskStagesTemplateRepo repo.TaskStagesTemplateRepo
	taskStagesRepo         repo.TaskStagesRepo
}

// 初始化
func New() *TaskService {
	return &TaskService{
		cache:                  dao.Rc,
		transaction:            dao.NewTrans(),
		projectRepo:            dao.NewProjectDao(),
		proTemplateRepo:        dao.NewProjectTemplateDao(),
		taskStagesRepo:         dao.NewTaskStagesDao(),
		taskStagesTemplateRepo: dao.NewTaskStagesTemplateDao(),
	}
}

func (t *TaskService) TaskStages(c context.Context, msg *task_service.TaskReqMessage) (*task_service.TaskStagesResponse, error) {
	projectCode := encrypts.DecryptToRes(msg.ProjectCode)
	page := msg.Page
	pageSize := msg.PageSize
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	taskStages, total, err := t.taskStagesRepo.FindByProjectCode(ctx, projectCode, page, pageSize)
	if err != nil {
		zap.L().Error("project task TaskStages FindByProjectCode error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	var resp []*task_service.TaskStagesMessage
	_ = copier.Copy(&resp, taskStages)
	if resp == nil {
		return &task_service.TaskStagesResponse{
			List:  resp,
			Total: 0,
		}, nil
	}
	tsMap := task.ToTaskStagesMap(taskStages)
	for _, v := range resp {
		stages := tsMap[int(v.Id)]
		v.Code, _ = encrypts.EncryptInt64(int64(v.Id), model.AESKey)
		v.CreateTime = times.FormatByMill(stages.CreateTime)
		v.ProjectCode = msg.ProjectCode
	}
	return &task_service.TaskStagesResponse{
		List:  resp,
		Total: total,
	}, nil
}
