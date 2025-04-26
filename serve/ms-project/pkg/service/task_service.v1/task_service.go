package project_service_v1

import (
	"context"
	"fmt"
	"hnz.com/ms_serve/ms-common/kafkas"
	"hnz.com/ms_serve/ms-project/config"
	"hnz.com/ms_serve/ms-project/internal/data/files"
	"hnz.com/ms_serve/ms-project/internal/domain"
	"time"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-common/times"
	taskRpc "hnz.com/ms_serve/ms-grpc/task"
	"hnz.com/ms_serve/ms-grpc/user/login"
	"hnz.com/ms_serve/ms-project/internal/dao"
	modelCom "hnz.com/ms_serve/ms-project/internal/data/common"
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
	proLogRepo             repo.ProjectLogRepo
	taskWorkTimeRepo       repo.TaskWorkTimeRepo
	fileRepo               repo.FileRepo
	sourceLinkRepo         repo.SourceLinkRepo
	taskWorkTimeDomain     *domain.TaskWorkTimeDomain
}

// New 初始化
func New() *TaskService {
	return &TaskService{
		cache:                  dao.Rc,
		transaction:            dao.NewTrans(),
		projectRepo:            dao.NewProjectDao(),
		proTemplateRepo:        dao.NewProjectTemplateDao(),
		taskStagesRepo:         dao.NewTaskStagesDao(),
		taskStagesTemplateRepo: dao.NewTaskStagesTemplateDao(),
		taskRepo:               dao.NewTaskDao(),
		proLogRepo:             dao.NewProjectLogDao(),
		taskWorkTimeRepo:       dao.NewTaskWorkTimeDao(),
		fileRepo:               dao.NewFileDao(),
		sourceLinkRepo:         dao.NewSourceLinkDao(),
		taskWorkTimeDomain:     domain.NewTaskWorkTimeDomain(),
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
		Sort:        int(maxSort + 65535),
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
			JoinTime:   time.Now().UnixMilli(),
			IsOwner:    model.Owner,
		}
		if assignTo == msg.MemberId {
			tm.IsOwner = model.Executor
		} else {
			tm.IsOwner = model.NoExecutor
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
	member, err := rpc.UserClient.FindMemberInfoById(ctx, &login.UserMessage{MemId: assignTo})
	if err != nil {
		return nil, err
	}
	display.Executor = task.Executor{
		Name:   member.Name,
		Avatar: member.Avatar,
		Code:   member.Code,
	}
	//添加任务动态
	createProjectLog(t.proLogRepo, ts.ProjectCode, ts.Id, ts.Name, ts.AssignTo, "create", "task")

	tm := &taskRpc.TaskMessage{}
	_ = copier.Copy(tm, display)
	return tm, nil
}

func (t *TaskService) TaskSort(ctx context.Context, msg *taskRpc.TaskReqMessage) (*taskRpc.TaskSortResponse, error) {
	preTaskCode := encrypts.DecryptToRes(msg.PreTaskCode)
	stageCode := encrypts.DecryptToRes(msg.ToStageCode)
	if msg.PreTaskCode == msg.NextTaskCode {
		return &taskRpc.TaskSortResponse{}, nil
	}
	err := t.sortTask(preTaskCode, msg.NextTaskCode, stageCode)
	if err != nil {
		return nil, err
	}
	return &taskRpc.TaskSortResponse{}, nil
}

func (t *TaskService) sortTask(preTaskCode int64, nextTaskCode string, stageCode int64) error {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	ts, err := t.taskRepo.FindTaskById(c, preTaskCode)
	if err != nil {
		zap.L().Error("project task TaskSort taskRepo.FindTaskById error", zap.Error(err))
		return errs.GrpcError(model.DataBaseError)
	}
	err = t.transaction.Action(func(conn database.DBConn) error {
		//如果相等是不需要进行改变的
		ts.StageCode = int(stageCode)
		if nextTaskCode != "" {
			//意味着要进行排序的替换
			nextTaskCode := encrypts.DecryptToRes(nextTaskCode)
			next, err := t.taskRepo.FindTaskById(c, nextTaskCode)
			if err != nil {
				zap.L().Error("project task TaskSort taskRepo.FindTaskById error", zap.Error(err))
				return errs.GrpcError(model.DataBaseError)
			}
			// next.Sort 要找到比它小的那个任务
			prepre, err := t.taskRepo.FindTaskByStageCodeSmallSort(c, next.StageCode, next.Sort)
			if err != nil {
				zap.L().Error("project task TaskSort taskRepo.FindTaskByStageCodeLtSort error", zap.Error(err))
				return errs.GrpcError(model.DataBaseError)
			}
			if prepre != nil {
				ts.Sort = (prepre.Sort + next.Sort) / 2
			}
			if prepre == nil {
				ts.Sort = 0
			}
		} else {
			maxSort, err := t.taskRepo.FindTaskSort(c, ts.ProjectCode, int64(ts.StageCode))
			if err != nil {
				zap.L().Error("project task TaskSort taskRepo.FindTaskSort error", zap.Error(err))
				return errs.GrpcError(model.DataBaseError)
			}
			ts.Sort = int(maxSort + 65535)
		}
		if ts.Sort < 50 {
			//重置排序
			err = t.resetSort(stageCode)
			if err != nil {
				zap.L().Error("project task TaskSort resetSort error", zap.Error(err))
				return errs.GrpcError(model.DataBaseError)
			}
			return t.sortTask(preTaskCode, nextTaskCode, stageCode)
		}
		err = t.taskRepo.UpdateTaskSort(c, conn, ts)
		if err != nil {
			zap.L().Error("project task TaskSort taskRepo.UpdateTaskSort error", zap.Error(err))
			return errs.GrpcError(model.DataBaseError)
		}
		return nil
	})
	return err
}
func (t *TaskService) resetSort(stageCode int64) error {
	list, err := t.taskRepo.FindTaskByStageCode(context.Background(), int(stageCode))
	if err != nil {
		return err
	}
	return t.transaction.Action(func(conn database.DBConn) error {
		iSort := 65535
		for index, v := range list {
			v.Sort = (index + 1) * iSort
			return t.taskRepo.UpdateTaskSort(context.Background(), conn, v)
		}
		return nil
	})
}
func (t *TaskService) MyTaskList(ctx context.Context, msg *taskRpc.TaskReqMessage) (*taskRpc.MyTaskListResponse, error) {
	var tsList []*task.Task
	var err error
	var total int64
	if msg.TaskType == 1 {
		//我执行的
		tsList, total, err = t.taskRepo.FindTaskByAssignTo(ctx, msg.MemberId, int(msg.Type), msg.Page, msg.PageSize)
		if err != nil {
			zap.L().Error("project task MyTaskList taskRepo.FindTaskByAssignTo error", zap.Error(err))
			return nil, errs.GrpcError(model.DataBaseError)
		}
	}
	if msg.TaskType == 2 {
		//我执行的
		tsList, total, err = t.taskRepo.FindTaskByMemberCode(ctx, msg.MemberId, int(msg.Type), msg.Page, msg.PageSize)
		if err != nil {
			zap.L().Error("project task MyTaskList taskRepo.FindTaskByMemberCode error", zap.Error(err))
			return nil, errs.GrpcError(model.DataBaseError)
		}
	}
	if msg.TaskType == 3 {
		//我执行的
		tsList, total, err = t.taskRepo.FindTaskByCreateBy(ctx, msg.MemberId, int(msg.Type), msg.Page, msg.PageSize)
		if err != nil {
			zap.L().Error("project task MyTaskList taskRepo.FindTaskByCreateBy error", zap.Error(err))
			return nil, errs.GrpcError(model.DataBaseError)
		}
	}
	if tsList == nil || len(tsList) <= 0 {
		return &taskRpc.MyTaskListResponse{List: nil, Total: 0}, nil
	}
	var pids []int64
	var mids []int64
	for _, v := range tsList {
		pids = append(pids, v.ProjectCode)
		mids = append(mids, v.AssignTo)
	}
	pListChan := make(chan []*project.Project)
	defer close(pListChan)
	mListChan := make(chan *login.MemberListResponse)
	defer close(mListChan)
	go func() {
		pList, _ := t.projectRepo.FindProjectByIds(ctx, pids)
		pListChan <- pList
	}()
	go func() {
		mList, _ := rpc.UserClient.FindMemberByIds(ctx, &login.UserMessage{
			MemberIds: mids,
		})
		mListChan <- mList
	}()
	pList := <-pListChan
	projectMap := project.ToProjectMap(pList)
	mList := <-mListChan

	mMap := make(map[int64]*login.MemberMessage)
	for _, v := range mList.MemberList {
		mMap[v.Id] = v
	}
	var mtdList []*task.MyTaskDisplay
	for _, v := range tsList {
		memberMessage := mMap[v.AssignTo]
		name := memberMessage.Name
		avatar := memberMessage.Avatar
		mtd := v.ToMyTaskDisplay(projectMap[v.ProjectCode], name, avatar)
		mtdList = append(mtdList, mtd)
	}
	var myMsgs []*taskRpc.MyTaskMessage
	_ = copier.Copy(&myMsgs, mtdList)
	return &taskRpc.MyTaskListResponse{List: myMsgs, Total: total}, nil
}
func (t *TaskService) ReadTask(ctx context.Context, msg *taskRpc.TaskReqMessage) (*taskRpc.TaskMessage, error) {
	taskCode := encrypts.DecryptToRes(msg.TaskCode)
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	taskInfo, err := t.taskRepo.FindTaskById(c, taskCode)
	if err != nil {
		zap.L().Error("project task ReadTask taskRepo FindTaskById error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	if taskInfo == nil {
		return &taskRpc.TaskMessage{}, nil
	}
	display := taskInfo.ToTaskDisplay()
	if taskInfo.Private == 1 {
		//代表隐私模式
		taskMember, err := t.taskRepo.FindTaskMemberByTaskId(ctx, taskInfo.Id, msg.MemberId)
		if err != nil {
			zap.L().Error("project task TaskList taskRepo.FindTaskMemberByTaskId error", zap.Error(err))
			return nil, errs.GrpcError(model.DataBaseError)
		}
		if taskMember != nil {
			display.CanRead = model.CanRead
		} else {
			display.CanRead = model.NoCanRead
		}
	}
	pj, err := t.projectRepo.FindProjectById(c, taskInfo.ProjectCode)
	if err != nil {
		zap.L().Error("project task TaskList FindProjectById error", zap.Error(err))
		return nil, err
	}
	display.ProjectName = pj.Name
	taskStages, err := t.taskStagesRepo.FindById(c, taskInfo.StageCode)
	if err != nil {
		zap.L().Error("project task TaskList FindById error", zap.Error(err))
		return nil, err
	}
	display.StageName = taskStages.Name
	// in ()
	memberMessage, err := rpc.UserClient.FindMemberInfoById(ctx, &login.UserMessage{MemId: taskInfo.AssignTo})
	if err != nil {
		zap.L().Error("project task TaskList UserClient.FindMemberInfoById error", zap.Error(err))
		return nil, err
	}
	e := task.Executor{
		Name:   memberMessage.Name,
		Avatar: memberMessage.Avatar,
	}
	display.Executor = e
	var taskMessage = &taskRpc.TaskMessage{}
	_ = copier.Copy(taskMessage, display)
	return taskMessage, nil
}

func (t *TaskService) ListTaskMember(ctx context.Context, msg *taskRpc.TaskReqMessage) (*taskRpc.TaskMemberList, error) {
	taskCode := encrypts.DecryptToRes(msg.TaskCode)
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	taskMemberPage, total, err := t.taskRepo.FindTaskMemberPage(c, taskCode, msg.Page, msg.PageSize)
	if err != nil {
		zap.L().Error("project task TaskList taskRepo.FindTaskMemberPage error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	var mids []int64
	for _, v := range taskMemberPage {
		mids = append(mids, v.MemberCode)
	}
	messageList, err := rpc.UserClient.FindMemberByIds(ctx, &login.UserMessage{MemberIds: mids})
	mMap := make(map[int64]*login.MemberMessage, len(messageList.MemberList))
	for _, v := range messageList.MemberList {
		mMap[v.Id] = v
	}
	var taskMemberMessages []*taskRpc.TaskMemberMessage
	for _, v := range taskMemberPage {
		tm := &taskRpc.TaskMemberMessage{}
		tm.Code = encrypts.EncryptNoErr(v.MemberCode)
		tm.Id = v.Id
		message := mMap[v.MemberCode]
		tm.Name = message.Name
		tm.Avatar = message.Avatar
		tm.IsExecutor = int32(v.IsExecutor)
		tm.IsOwner = int32(v.IsOwner)
		taskMemberMessages = append(taskMemberMessages, tm)
	}
	return &taskRpc.TaskMemberList{List: taskMemberMessages, Total: total}, nil
}
func (t *TaskService) TaskLog(ctx context.Context, msg *taskRpc.TaskReqMessage) (*taskRpc.TaskLogList, error) {
	config.SendLog(kafkas.Info("TaskLog", "TaskLog", kafkas.FieldMap{
		"msg": msg,
	}))
	taskCode := encrypts.DecryptToRes(msg.TaskCode)
	all := msg.All
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var list []*project.ProjectLog
	var total int64
	var err error
	if all == 1 {
		//显示全部
		list, total, err = t.proLogRepo.FindLogByTaskCode(c, taskCode, int(msg.Comment))
	}
	if all == 0 {
		//分页
		list, total, err = t.proLogRepo.FindLogByTaskCodePage(c, taskCode, int(msg.Comment), int(msg.Page), int(msg.PageSize))
	}
	if err != nil {
		zap.L().Error("project task TaskLog projectLogRepo.FindLogByTaskCodePage error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	if total == 0 {
		return &taskRpc.TaskLogList{}, nil
	}
	var displayList []*project.ProjectLogDisplay
	var mIdList []int64
	for _, v := range list {
		mIdList = append(mIdList, v.MemberCode)
	}
	messageList, err := rpc.UserClient.FindMemberByIds(c, &login.UserMessage{MemberIds: mIdList})
	mMap := make(map[int64]*login.MemberMessage)
	for _, v := range messageList.MemberList {
		mMap[v.Id] = v
	}
	for _, v := range list {
		display := v.ToDisplay()
		message := mMap[v.MemberCode]
		m := modelCom.Member{}
		m.Name = message.Name
		m.Id = message.Id
		m.Avatar = message.Avatar
		m.Code = message.Code
		display.Member = m
		displayList = append(displayList, display)
	}
	var l []*taskRpc.TaskLog
	_ = copier.Copy(&l, displayList)
	return &taskRpc.TaskLogList{List: l, Total: total}, nil
}
func (t *TaskService) TaskWorkTimeList(ctx context.Context, msg *taskRpc.TaskReqMessage) (*taskRpc.TaskWorkTimeResponse, error) {
	taskCode := encrypts.DecryptToRes(msg.TaskCode)
	displayList, bError := t.taskWorkTimeDomain.TaskWorkTimeList(taskCode)
	if bError != nil {
		zap.L().Error("project task TaskWorkTimeList userRpcDomain.MemberList error", zap.Error(bError))
		return nil, bError
	}
	var l []*taskRpc.TaskWorkTime
	_ = copier.Copy(&l, displayList)
	return &taskRpc.TaskWorkTimeResponse{List: l, Total: int64(len(l))}, nil
}

func (t *TaskService) SaveTaskWorkTime(ctx context.Context, msg *taskRpc.TaskReqMessage) (*taskRpc.SaveTaskWorkTimeResponse, error) {
	tmt := &task.TaskWorkTime{}
	tmt.BeginTime = msg.BeginTime
	tmt.Num = int(msg.Num)
	tmt.Content = msg.Content
	tmt.TaskCode = encrypts.DecryptToRes(msg.TaskCode)
	tmt.MemberCode = msg.MemberId
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := t.taskWorkTimeRepo.Save(c, tmt)
	if err != nil {
		zap.L().Error("project task SaveTaskWorkTime taskWorkTimeRepo.Save error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	return &taskRpc.SaveTaskWorkTimeResponse{}, nil
}

func (t *TaskService) SaveTaskFile(ctx context.Context, msg *taskRpc.TaskFileReqMessage) (*taskRpc.TaskFileResponse, error) {
	taskCode := encrypts.DecryptToRes(msg.TaskCode)
	//存file表
	f := &files.File{
		PathName:         msg.PathName,
		Title:            msg.FileName,
		Extension:        msg.Extension,
		Size:             int(msg.Size),
		ObjectType:       "",
		OrganizationCode: encrypts.DecryptToRes(msg.OrganizationCode),
		TaskCode:         encrypts.DecryptToRes(msg.TaskCode),
		ProjectCode:      encrypts.DecryptToRes(msg.ProjectCode),
		CreateBy:         msg.MemberId,
		CreateTime:       time.Now().UnixMilli(),
		Downloads:        0,
		Extra:            "",
		Deleted:          model.NoDeleted,
		FileType:         msg.FileType,
		FileUrl:          msg.FileUrl,
		DeletedTime:      0,
	}
	err := t.fileRepo.Save(context.Background(), f)
	if err != nil {
		zap.L().Error("project task SaveTaskFile fileRepo.Save error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	//存入source_link
	sl := &files.SourceLink{
		SourceType:       "file",
		SourceCode:       f.Id,
		LinkType:         "task",
		LinkCode:         taskCode,
		OrganizationCode: encrypts.DecryptToRes(msg.OrganizationCode),
		CreateBy:         msg.MemberId,
		CreateTime:       time.Now().UnixMilli(),
		Sort:             0,
	}
	err = t.sourceLinkRepo.Save(context.Background(), sl)
	if err != nil {
		zap.L().Error("project task SaveTaskFile sourceLinkRepo.Save error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	return &taskRpc.TaskFileResponse{}, nil
}

func (t *TaskService) TaskSources(ctx context.Context, msg *taskRpc.TaskReqMessage) (*taskRpc.TaskSourceResponse, error) {
	taskCode := encrypts.DecryptToRes(msg.TaskCode)
	sourceLinks, err := t.sourceLinkRepo.FindByTaskCode(context.Background(), taskCode)
	if err != nil {
		zap.L().Error("project task SaveTaskFile sourceLinkRepo.FindByTaskCode error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	if len(sourceLinks) == 0 {
		return &taskRpc.TaskSourceResponse{}, nil
	}
	var fIdList []int64
	for _, v := range sourceLinks {
		fIdList = append(fIdList, v.SourceCode)
	}
	fileRes, err := t.fileRepo.FindByIds(context.Background(), fIdList)
	if err != nil {
		zap.L().Error("project task SaveTaskFile fileRepo.FindByIds error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	fMap := make(map[int64]*files.File)
	for _, v := range fileRes {
		fMap[v.Id] = v
	}
	var list []*files.SourceLinkDisplay
	for _, v := range sourceLinks {
		list = append(list, v.ToDisplay(fMap[v.SourceCode]))
	}
	var slMsg []*taskRpc.TaskSourceMessage
	_ = copier.Copy(&slMsg, list)
	return &taskRpc.TaskSourceResponse{List: slMsg}, nil
}
func (t *TaskService) CreateComment(ctx context.Context, msg *taskRpc.TaskReqMessage) (*taskRpc.CreateCommentResponse, error) {
	taskCode := encrypts.DecryptToRes(msg.TaskCode)
	taskById, err := t.taskRepo.FindTaskById(context.Background(), taskCode)
	if err != nil {
		zap.L().Error("project task CreateComment fileRepo.FindTaskById error", zap.Error(err))
		return nil, errs.GrpcError(model.DataBaseError)
	}
	pl := &project.ProjectLog{
		MemberCode:   msg.MemberId,
		Content:      msg.CommentContent,
		Remark:       msg.CommentContent,
		Type:         "createComment",
		CreateTime:   time.Now().UnixMilli(),
		SourceCode:   taskCode,
		ActionType:   "task",
		ToMemberCode: 0,
		IsComment:    model.Comment,
		ProjectCode:  taskById.ProjectCode,
		Icon:         "plus",
		IsRobot:      0,
	}
	t.proLogRepo.SaveProjectLog(pl)
	return &taskRpc.CreateCommentResponse{}, nil
}

func createProjectLog(
	logRepo repo.ProjectLogRepo,
	projectCode int64,
	taskCode int64,
	taskName string,
	toMemberCode int64,
	logType string,
	actionType string) {
	remark := ""
	if logType == "create" {
		remark = "创建了任务"
	}
	pl := &project.ProjectLog{
		MemberCode:  toMemberCode,
		SourceCode:  taskCode,
		Content:     taskName,
		Remark:      remark,
		ProjectCode: projectCode,
		CreateTime:  time.Now().UnixMilli(),
		Type:        logType,
		ActionType:  actionType,
		Icon:        "plus",
		IsComment:   0,
		IsRobot:     0,
	}
	logRepo.SaveProjectLog(pl)
}
