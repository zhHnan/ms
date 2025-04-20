package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"hnz.com/ms_serve/ms-api/api/rpc"
	"hnz.com/ms_serve/ms-api/pkg/model"
	"hnz.com/ms_serve/ms-api/pkg/model/apiProject"
	"hnz.com/ms_serve/ms-api/pkg/model/menu"
	common "hnz.com/ms_serve/ms-common"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-grpc/project"
	"net/http"
	"strconv"
	"time"
)

type HandlerProject struct {
}

func New() *HandlerProject {

	return &HandlerProject{}
}

func (p *HandlerProject) index(ctx *gin.Context) {
	result := &common.Result{}
	c, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	msg := &project.IndexMessage{}
	resp, err := rpc.ProjectClient.Index(c, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(200, result.Failure(code, msg))
		return
	}
	var menus []*menu.Menu
	_ = copier.Copy(&menus, resp.Menus)
	ctx.JSON(200, result.Success(menus))
}

func (p *HandlerProject) projectList(c *gin.Context) {
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	id := c.GetInt64("memberId")
	memberName := c.GetString("memberName")
	var page = &model.Page{}
	page.Bind(c)
	selectBy := c.PostForm("selectBy")

	msg := &project.ProjectRpcMessage{
		MemberId:   id,
		MemberName: memberName,
		Page:       page.Page,
		PageSize:   page.PageSize,
		SelectBy:   selectBy,
	}
	res, err := rpc.ProjectClient.FindProjectByMemId(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	var pam []*apiProject.ProjectAndMember
	_ = copier.Copy(&pam, res.Pm)
	if pam == nil {
		pam = []*apiProject.ProjectAndMember{}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  pam,
		"total": res.Total,
	}))
}

func (p *HandlerProject) projectTemplate(c *gin.Context) {
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	id := c.GetInt64("memberId")
	memberName := c.GetString("memberName")
	var page = &model.Page{}
	page.Bind(c)
	viewTypeStr := c.PostForm("viewType")
	viewType, _ := strconv.ParseInt(viewTypeStr, 10, 64)

	msg := &project.ProjectRpcMessage{
		MemberId:         id,
		MemberName:       memberName,
		Page:             page.Page,
		PageSize:         page.PageSize,
		ViewType:         int32(viewType),
		OrganizationCode: c.GetString("organizationCode"),
	}
	res, err := rpc.ProjectClient.FindProjectTemplate(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	var pam []*apiProject.ProjectTemplate
	_ = copier.Copy(&pam, res.Ptm)
	if pam == nil {
		pam = []*apiProject.ProjectTemplate{}
	}
	for _, v := range pam {
		if v.TaskStages == nil {
			v.TaskStages = []*apiProject.TaskStagesOnlyName{}
		}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  pam,
		"total": res.Total,
	}))
}

func (p *HandlerProject) projectSave(c *gin.Context) {
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberId := c.GetInt64("memberId")
	organizationCode := c.GetString("organizationCode")
	var req *apiProject.SaveProjectRequest
	_ = c.ShouldBind(&req)
	msg := &project.ProjectRpcMessage{
		Id:               int64(req.Id),
		Name:             req.Name,
		Description:      req.Description,
		TemplateCode:     req.TemplateCode,
		OrganizationCode: organizationCode,
		MemberId:         memberId,
	}
	res, err := rpc.ProjectClient.SaveProject(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	var resp *project.SaveProjectMessage
	_ = copier.Copy(&resp, res)
	c.JSON(http.StatusOK, result.Success(resp))
}

func (p *HandlerProject) projectRead(c *gin.Context) {
	result := &common.Result{}
	projectCode := c.PostForm("projectCode")
	memberId := c.GetInt64("memberId")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	detail, err := rpc.ProjectClient.GetProjectDetail(ctx, &project.ProjectRpcMessage{ProjectCode: projectCode, MemberId: memberId})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	pd := &apiProject.ProjectDetail{}
	_ = copier.Copy(pd, detail)
	c.JSON(http.StatusOK, result.Success(pd))
}

func (p *HandlerProject) projectRecycle(c *gin.Context) {
	result := &common.Result{}
	projectCode := c.PostForm("projectCode")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := rpc.ProjectClient.UpdateDeletedProject(ctx, &project.ProjectRpcMessage{ProjectCode: projectCode, Deleted: true})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	c.JSON(http.StatusOK, result.Success([]int{}))
}

func (p *HandlerProject) projectRecovery(c *gin.Context) {
	result := &common.Result{}
	projectCode := c.PostForm("projectCode")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := rpc.ProjectClient.UpdateDeletedProject(ctx, &project.ProjectRpcMessage{ProjectCode: projectCode, Deleted: false})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	c.JSON(http.StatusOK, result.Success([]int{}))
}
