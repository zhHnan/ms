package project

import (
	"context"
	"github.com/gin-gonic/gin"
	common "hnz.com/ms_serve/ms-common"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-grpc/project"
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
