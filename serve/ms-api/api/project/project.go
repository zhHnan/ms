package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"hnz.com/ms_serve/ms-api/pkg/model"
	"hnz.com/ms_serve/ms-api/pkg/model/apiProject"
	common "hnz.com/ms_serve/ms-common"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-grpc/project"
	"net/http"
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
	resp, err := ProjectClient.Index(c, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(200, result.Failure(code, msg))
		return
	}
	ctx.JSON(200, result.Success(resp.Menus))
}

func (p *HandlerProject) projectList(c *gin.Context) {
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	idAny, _ := c.Get("memberId")
	id := idAny.(int64)
	var page = &model.Page{}
	page.Bind(c)
	res, err := ProjectClient.FindProjectByMemId(ctx, &project.ProjectRpcMessage{MemberId: id, Page: page.Page, PageSize: page.PageSize})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	if res.Pm == nil {
		res.Pm = []*project.ProjectMessage{}
	}
	var pam []*apiProject.ProjectAndMember
	_ = copier.Copy(&pam, res.Pm)
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  pam,
		"total": res.Total,
	}))
}
