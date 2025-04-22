package project_service_v1

import (
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-common/times"
	taskRpc "hnz.com/ms_serve/ms-grpc/task"
	"hnz.com/ms_serve/ms-grpc/user/login"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/data/project"
	"hnz.com/ms_serve/ms-project/internal/data/task"
	"hnz.com/ms_serve/ms-project/internal/database"
	"hnz.com/ms_serve/ms-project/internal/database/tran"
	"hnz.com/ms_serve/ms-project/internal/repo"
	"hnz.com/ms_serve/ms-project/internal/rpc"
	"hnz.com/ms_serve/ms-project/pkg/model"
)

type TaskService struct {
	taskRpc.UnimplementedTaskServiceServer
	cache                  repo.Cache
	transaction            tran.Transaction
	projectRepo            repo.ProjectRepo
	proTemplateRepo        repo.ProjectTemplateRepo
	taskStagesTemplateRepo repo.TaskStagesTemplateRepo
	taskStagesRepo         repo.TaskStagesRepo
	taskRepo               repo.TaskRepo
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
		taskRepo:               dao.NewTaskDao(),
	}
}

func (t *TaskService) TaskStages(c context.Context, msg *taskRpc.TaskReqMessage) (*taskRpc.TaskStagesResponse, error) {
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
	var resp []*taskRpc.TaskStagesMessage
	_ = copier.Copy(&resp, taskStages)
	if resp == nil {
		return &taskRpc.TaskStagesResponse{
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
	return &taskRpc.TaskStagesResponse{
		List:  resp,
		Total: total,
	}, nil
}
func (t *TaskService) MemberProjectList(c context.Context, msg *taskRpc.TaskReqMessage) (*taskRpc.MemberProjectResponse, error) {
	projectCode := encrypts.DecryptToRes(msg.ProjectCode)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	memberInfos, total, err := t.projectRepo.FindMemberByProjectId(ctx, projectCode)
	if err != nil {
		zap.L().Error("project task TaskStages FindMemberByProjectId error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	if memberInfos == nil || len(memberInfos) <= 0 {
		return &taskRpc.MemberProjectResponse{
			List:  nil,
			Total: 0,
		}, nil
	}

	var memberIds []int64
	pmMap := make(map[int64]*project.ProjectMember)
	for _, v := range memberInfos {
		memberIds = append(memberIds, v.MemberCode)
		pmMap[v.MemberCode] = v
	}
	userMessage := &login.UserMessage{
		MemberIds: memberIds,
	}
	//fmt.Printf("\n userMessage--memberIds:【%s】\n", userMessage.MemberIds)
	members, err := rpc.UserClient.FindMemberByIds(ctx, userMessage)
	if err != nil {
		zap.L().Error("project task TaskStages userClient.FindMemberInfoByIds error", zap.Error(err))
		return nil, err
	}
	var list []*taskRpc.MemberProjectMessage
	for _, v := range members.MemberList {
		owner := pmMap[v.Id].IsOwner
		mpm := &taskRpc.MemberProjectMessage{
			MemberCode: v.Id,
			Name:       v.Name,
			Avatar:     v.Avatar,
			Email:      v.Email,
			Code:       v.Code,
		}
		if v.Id == owner {
			mpm.IsOwner = int32(model.Owner)
		} else {
			mpm.IsOwner = int32(model.Normal)
		}
		list = append(list, mpm)
	}
	return &taskRpc.MemberProjectResponse{
		List:  list,
		Total: total,
	}, nil
}

func (t *TaskService) TaskList(ctx context.Context, msg *taskRpc.TaskReqMessage) (*taskRpc.TaskListResponse, error) {
	stageCode := encrypts.DecryptToRes(msg.StageCode)
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	fmt.Println("【TaskList】stageCode:", stageCode)
	taskList, err := t.taskRepo.FindTaskByStageCode(c, int(stageCode))
	if err != nil {
		zap.L().Error("project task TaskList FindTaskByStageCode error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	//fmt.Println("【TaskList】taskList:", taskList)
	var taskDisplayList []*task.TaskDisplay
	taskDisplayMap := make(map[int64]*task.TaskDisplay) // 使用map替代数组索引
	var mIds []int64
	for _, v := range taskList {
		display := v.ToTaskDisplay()
		if v.Private == model.Private {
			tm, err := t.taskRepo.FindTaskMemberByTaskId(ctx, v.Id, msg.MemberId)
			if err != nil {
				zap.L().Error("project task TaskList FindTaskMemberByTaskId error", zap.Error(err))
				return nil, errs.GrpcError(model.DataBaseError)
			}
			if tm == nil {
				display.CanRead = model.NoCanRead
			} else {
				display.CanRead = model.CanRead
			}
		}
		taskDisplayList = append(taskDisplayList, display)
		taskDisplayMap[v.Id] = display // 将任务添加到map中，以ID为键
		mIds = append(mIds, v.AssignTo)
	}
	if len(mIds) <= 0 {
		return &taskRpc.TaskListResponse{
			List: nil,
		}, nil
	}
	memberList, err := rpc.UserClient.FindMemberByIds(c, &login.UserMessage{MemberIds: mIds})
	if err != nil {
		zap.L().Error("project task TaskList UserClient.FindMemberByIds error", zap.Error(err))
		return nil, err
	}

	memberMap := make(map[int64]*login.MemberMessage)
	for _, v := range memberList.MemberList {
		memberMap[v.Id] = v
	}
	for _, v := range taskDisplayList {
		assignTo := encrypts.DecryptToRes(v.AssignTo)
		message := memberMap[assignTo]
		executor := task.Executor{
			Name:   message.Name,
			Avatar: message.Avatar,
		}
		// 使用map而不是数组索引来更新executor
		taskDisplayMap[v.Id].Executor = executor
	}
	var taskMessageList []*taskRpc.TaskMessage
	_ = copier.Copy(&taskMessageList, taskDisplayList)
	return &taskRpc.TaskListResponse{List: taskMessageList}, nil
}

func (t *TaskService) SaveTask(ctx context.Context, msg *taskRpc.TaskReqMessage) (*taskRpc.TaskMessage, error) {
	//先检查业务
	if msg.Name == "" {
		return nil, errs.GrpcError(model.TaskNameNotNull)
	}
	stageCode := encrypts.DecryptToRes(msg.StageCode)
	taskStages, err := t.taskStagesRepo.FindById(ctx, int(stageCode))
	if err != nil {
		zap.L().Error("project task SaveTask taskStagesRepo.FindById error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	if taskStages == nil {
		return nil, errs.GrpcError(model.TaskStagesNotNull)
	}
	projectCode := encrypts.DecryptToRes(msg.ProjectCode)
	projectById, err := t.projectRepo.FindProjectById(ctx, projectCode)
	if err != nil {
		zap.L().Error("project task SaveTask projectRepo.FindProjectById error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	if projectById.Deleted == model.Deleted {
		return nil, errs.GrpcError(model.ProjectAlreadyDeleted)
	}
	maxIdNum, err := t.taskRepo.FindTaskMaxIdNum(ctx, projectCode)
	if err != nil {
		zap.L().Error("project task SaveTask taskRepo.FindTaskMaxIdNum error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	maxSort, err := t.taskRepo.FindTaskSort(ctx, projectCode, stageCode)
	if err != nil {
		zap.L().Error("project task SaveTask taskRepo.FindTaskSort error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	assignTo := encrypts.DecryptToRes(msg.AssignTo)
	ts := &task.Task{
		Name:        msg.Name,
		CreateTime:  time.Now().UnixMilli(),
		CreateBy:    msg.MemberId,
		AssignTo:    assignTo,
		ProjectCode: projectCode,
		StageCode:   int(stageCode),
		IdNum:       int(maxIdNum + 1),
		Private:     projectById.OpenTaskPrivate,
		Sort:        int(maxSort + 1),
		BeginTime:   time.Now().UnixMilli(),
		EndTime:     time.Now().Add(2 * 24 * time.Hour).UnixMilli(),
	}
	err = t.transaction.Action(func(conn database.DBConn) error {
		err = t.taskRepo.SaveTask(ctx, conn, ts)
		if err != nil {
			zap.L().Error("project task SaveTask taskRepo.SaveTask error", zap.Error(err))
			return errs.GrpcError(model.DataBaseError)
		}
		tm := &task.TaskMember{
			MemberCode: msg.MemberId,
			TaskCode:   ts.Id,
			IsExecutor: model.Executor,
			JoinTime:   time.Now().UnixMilli(),
			IsOwner:    model.Owner,
		}
		err = t.taskRepo.SaveTaskMember(ctx, conn, tm)
		if err != nil {
			zap.L().Error("project task SaveTask taskRepo.SaveTaskMember error", zap.Error(err))
			return errs.GrpcError(model.DataBaseError)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	display := ts.ToTaskDisplay()
	tm := &taskRpc.TaskMessage{}
	_ = copier.Copy(tm, display)
	return tm, nil
}
