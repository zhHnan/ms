package project

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"hnz.com/ms_serve/ms-api/api/rpc"
	"hnz.com/ms_serve/ms-api/pkg/model"
	"hnz.com/ms_serve/ms-api/pkg/model/apiProject"
	"hnz.com/ms_serve/ms-api/pkg/model/comment"
	"hnz.com/ms_serve/ms-api/pkg/model/file"
	"hnz.com/ms_serve/ms-api/pkg/model/tasks"
	common "hnz.com/ms_serve/ms-common"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-common/files"
	"hnz.com/ms_serve/ms-common/times"
	taskRpc "hnz.com/ms_serve/ms-grpc/task"
	"net/http"
	"os"
	"path"
	"time"
)

type HandlerTask struct {
}

func NewTask() *HandlerTask {

	return &HandlerTask{}
}

func (t *HandlerTask) taskStages(c *gin.Context) {
	result := &common.Result{}
	projectCode := c.PostForm("projectCode")
	page := &model.Page{}
	page.Bind(c)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &taskRpc.TaskReqMessage{
		MemberId:    c.GetInt64("memberId"),
		ProjectCode: projectCode,
		Page:        page.Page,
		PageSize:    page.PageSize,
	}
	stages, err := rpc.TaskClient.TaskStages(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	var resp []*tasks.TaskStagesResp
	copier.Copy(&resp, stages.List)
	if resp == nil {
		resp = []*tasks.TaskStagesResp{}
	}
	for _, v := range resp {
		v.TasksLoading = true  //任务加载状态
		v.FixedCreator = false //添加任务按钮定位
		v.ShowTaskCard = false //是否显示创建卡片
		v.Tasks = []int{}
		v.DoneTasks = []int{}
		v.UnDoneTasks = []int{}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  resp,
		"total": stages.Total,
		"page":  page.Page,
	}))
}

func (t *HandlerTask) taskMemberList(c *gin.Context) {
	result := &common.Result{}
	projectCode := c.PostForm("projectCode")
	page := &model.Page{}
	page.Bind(c)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &taskRpc.TaskReqMessage{
		ProjectCode: projectCode,
		Page:        page.Page,
		PageSize:    page.PageSize,
	}
	memberResp, err := rpc.TaskClient.MemberProjectList(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	var resp []*apiProject.MemberProjectResp
	copier.Copy(&resp, memberResp.List)
	if resp == nil {
		resp = []*apiProject.MemberProjectResp{}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  resp,
		"total": memberResp.Total,
		"page":  page.Page,
	}))
}

func (t *HandlerTask) taskList(c *gin.Context) {
	result := &common.Result{}
	stageCode := c.PostForm("stageCode")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	list, err := rpc.TaskClient.TaskList(ctx, &taskRpc.TaskReqMessage{StageCode: stageCode, MemberId: c.GetInt64("memberId")})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	var taskDisplayList []*tasks.TaskDisplay
	_ = copier.Copy(&taskDisplayList, list.List)
	if taskDisplayList == nil {
		taskDisplayList = []*tasks.TaskDisplay{}
	}
	for _, v := range taskDisplayList {
		if v.Tags == nil {
			v.Tags = []int{}
		}
		if v.ChildCount == nil {
			v.ChildCount = []int{}
		}
	}
	c.JSON(http.StatusOK, result.Success(taskDisplayList))
}

func (t *HandlerTask) taskSave(c *gin.Context) {
	result := &common.Result{}
	var req *tasks.TaskSaveReq
	err := c.ShouldBind(&req)
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &taskRpc.TaskReqMessage{
		ProjectCode: req.ProjectCode,
		Name:        req.Name,
		StageCode:   req.StageCode,
		AssignTo:    req.AssignTo,
		MemberId:    c.GetInt64("memberId"),
	}
	taskMessage, err := rpc.TaskClient.SaveTask(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	td := &tasks.TaskDisplay{}
	_ = copier.Copy(td, taskMessage)
	if td != nil {
		if td.Tags == nil {
			td.Tags = []int{}
		}
		if td.ChildCount == nil {
			td.ChildCount = []int{}
		}
	}
	c.JSON(http.StatusOK, result.Success(td))
}

func (t *HandlerTask) taskSort(c *gin.Context) {
	result := &common.Result{}
	var req *tasks.TaskSortReq
	err := c.ShouldBind(&req)
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &taskRpc.TaskReqMessage{
		PreTaskCode:  req.PreTaskCode,
		NextTaskCode: req.NextTaskCode,
		ToStageCode:  req.ToStageCode,
	}
	_, err = rpc.TaskClient.TaskSort(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	c.JSON(http.StatusOK, result.Success([]int{}))
}

func (t *HandlerTask) selfList(c *gin.Context) {
	result := &common.Result{}
	var req *tasks.MyTaskReq
	err := c.ShouldBind(&req)
	if err != nil {
		return
	}
	memberId := c.GetInt64("memberId")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &taskRpc.TaskReqMessage{
		MemberId: memberId,
		TaskType: int32(req.TaskType),
		Type:     int32(req.Type),
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	myTaskListResponse, err := rpc.TaskClient.MyTaskList(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	var myTaskList []*tasks.MyTaskDisplay
	_ = copier.Copy(&myTaskList, myTaskListResponse.List)
	if myTaskList == nil {
		myTaskList = []*tasks.MyTaskDisplay{}
	}
	for _, v := range myTaskList {
		v.ProjectInfo = tasks.ProjectInfo{
			Name: v.ProjectName,
			Code: v.ProjectCode,
		}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  myTaskList,
		"total": myTaskListResponse.Total,
	}))
}

func (t *HandlerTask) taskRead(c *gin.Context) {
	result := &common.Result{}
	taskCode := c.PostForm("taskCode")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &taskRpc.TaskReqMessage{
		TaskCode: taskCode,
		MemberId: c.GetInt64("memberId"),
	}
	taskMessage, err := rpc.TaskClient.ReadTask(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	td := &tasks.TaskDisplay{}
	_ = copier.Copy(td, taskMessage)
	if td != nil {
		if td.Tags == nil {
			td.Tags = []int{}
		}
		if td.ChildCount == nil {
			td.ChildCount = []int{}
		}
	}
	c.JSON(200, result.Success(td))
}

func (t *HandlerTask) listTaskMember(c *gin.Context) {
	result := &common.Result{}
	taskCode := c.PostForm("taskCode")
	page := &model.Page{}
	page.Bind(c)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &taskRpc.TaskReqMessage{
		TaskCode: taskCode,
		MemberId: c.GetInt64("memberId"),
		Page:     page.Page,
		PageSize: page.PageSize,
	}
	taskMemberResponse, err := rpc.TaskClient.ListTaskMember(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	var tms []*tasks.TaskMember
	_ = copier.Copy(&tms, taskMemberResponse.List)
	if tms == nil {
		tms = []*tasks.TaskMember{}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  tms,
		"total": taskMemberResponse.Total,
		"page":  page.Page,
	}))
}

func (t *HandlerTask) taskLog(c *gin.Context) {
	result := &common.Result{}
	var req *tasks.TaskLogReq
	err := c.ShouldBind(&req)
	if err != nil {
		return
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &taskRpc.TaskReqMessage{
		TaskCode: req.TaskCode,
		MemberId: c.GetInt64("memberId"),
		Page:     int64(req.Page),
		PageSize: int64(req.PageSize),
		All:      int32(req.All),
		Comment:  int32(req.Comment),
	}
	taskLogResponse, err := rpc.TaskClient.TaskLog(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	var tms []*tasks.ProjectLogDisplay
	_ = copier.Copy(&tms, taskLogResponse.List)
	if tms == nil {
		tms = []*tasks.ProjectLogDisplay{}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  tms,
		"total": taskLogResponse.Total,
		"page":  req.Page,
	}))
}

func (t *HandlerTask) taskWorkTimeList(c *gin.Context) {
	taskCode := c.PostForm("taskCode")
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &taskRpc.TaskReqMessage{
		TaskCode: taskCode,
		MemberId: c.GetInt64("memberId"),
	}
	taskWorkTimeResponse, err := rpc.TaskClient.TaskWorkTimeList(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	var tms []*tasks.TaskWorkTime
	_ = copier.Copy(&tms, taskWorkTimeResponse.List)
	if tms == nil {
		tms = []*tasks.TaskWorkTime{}
	}
	c.JSON(http.StatusOK, result.Success(tms))
}

func (t *HandlerTask) saveTaskWorkTime(c *gin.Context) {
	result := &common.Result{}
	var req *tasks.SaveTaskWorkTimeReq
	err := c.ShouldBind(&req)
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &taskRpc.TaskReqMessage{
		TaskCode:  req.TaskCode,
		MemberId:  c.GetInt64("memberId"),
		Content:   req.Content,
		Num:       int32(req.Num),
		BeginTime: times.ParseTime(req.BeginTime),
	}
	_, err = rpc.TaskClient.SaveTaskWorkTime(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	c.JSON(http.StatusOK, result.Success([]int{}))
}

func (t *HandlerTask) uploadFiles(c *gin.Context) {
	result := &common.Result{}
	req := file.UploadFileReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		return
	}

	multipartForm, _ := c.MultipartForm()
	fileForm := multipartForm.File["file"]
	key := ""
	// 假设只上传一个文件
	uploadFile := fileForm[0]
	// 若没有达成分片条件
	if req.TotalChunks == 1 {
		path := "upload/" + req.ProjectCode + "/" + req.TaskCode + "/" + times.FormatYMD(time.Now())
		if !files.IsExist(path) {
			_ = os.MkdirAll(path, os.ModePerm)
		}
		dst := path + "/" + req.Filename
		key = dst
		err := c.SaveUploadedFile(uploadFile, dst)
		if err != nil {
			c.JSON(http.StatusOK, result.Failure(9999, err.Error()))
			return
		}
	}
	if req.TotalChunks > 1 {
		// 分片上传
		path := "upload/" + req.ProjectCode + "/" + req.TaskCode + "/" + times.FormatYMD(time.Now())
		if !files.IsExist(path) {
			_ = os.MkdirAll(path, os.ModePerm)
		}
		fileName := path + "/" + req.Identifier
		// O_RDWR 作设置文件打开模式；ModePerm 是文件权限，0777 表示文件权限为 777，即所有用户都有读、写、执行权限。
		openFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusOK, result.Failure(9999, err.Error()))
			return
		}
		open, err := uploadFile.Open()
		if err != nil {
			c.JSON(http.StatusOK, result.Failure(9999, err.Error()))
			return
		}
		defer open.Close()
		buf := make([]byte, req.CurrentChunkSize)
		_, _ = open.Read(buf)
		_, err = openFile.Write(buf)
		fmt.Println(err)
		err = openFile.Close()
		fmt.Println(err)
		key = fileName
		if req.TotalChunks == req.ChunkNumber {
			//最后一块 重命名文件名
			newPath := path + "/" + req.Filename
			key = newPath
			err := os.Rename(fileName, newPath)
			fmt.Println(err)
		}
	}
	//调用服务 存入file表
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	fileUrl := "http://localhost/" + key
	msg := &taskRpc.TaskFileReqMessage{
		TaskCode:         req.TaskCode,
		ProjectCode:      req.ProjectCode,
		OrganizationCode: c.GetString("organizationCode"),
		PathName:         key,
		FileName:         req.Filename,
		Size:             int64(req.TotalSize),
		Extension:        path.Ext(key),
		FileUrl:          fileUrl,
		FileType:         fileForm[0].Header.Get("Content-Type"),
		MemberId:         c.GetInt64("memberId"),
	}
	if req.TotalChunks == req.ChunkNumber {
		_, err = rpc.TaskClient.SaveTaskFile(ctx, msg)
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(http.StatusOK, result.Failure(code, msg))
		}
	}

	c.JSON(http.StatusOK, result.Success(gin.H{
		"file":        key,
		"hash":        "",
		"key":         key,
		"url":         "http://localhost/" + key,
		"projectName": req.ProjectName,
	}))
	return
}

func (t *HandlerTask) taskSources(c *gin.Context) {
	result := &common.Result{}
	taskCode := c.PostForm("taskCode")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	sources, err := rpc.TaskClient.TaskSources(ctx, &taskRpc.TaskReqMessage{TaskCode: taskCode})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	var slList []*file.SourceLink
	_ = copier.Copy(&slList, sources.List)
	if slList == nil {
		slList = []*file.SourceLink{}
	}
	c.JSON(http.StatusOK, result.Success(slList))
}

func (t *HandlerTask) createComment(c *gin.Context) {
	result := &common.Result{}
	req := comment.CommentReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &taskRpc.TaskReqMessage{
		TaskCode:       req.TaskCode,
		CommentContent: req.Comment,
		Mentions:       req.Mentions,
		MemberId:       c.GetInt64("memberId"),
	}
	_, err = rpc.TaskClient.CreateComment(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	c.JSON(http.StatusOK, result.Success(true))
}
