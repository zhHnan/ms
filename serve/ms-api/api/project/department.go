package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"hnz.com/ms_serve/ms-api/api/rpc"
	"hnz.com/ms_serve/ms-api/pkg/model/account"
	common "hnz.com/ms_serve/ms-common"
	"hnz.com/ms_serve/ms-common/errs"
	departmentRpc "hnz.com/ms_serve/ms-grpc/department"
	"net/http"
	"time"
)

type HandlerDepartment struct {
}

func NewDepartment() *HandlerDepartment {

	return &HandlerDepartment{}
}

func (d *HandlerDepartment) department(c *gin.Context) {
	result := &common.Result{}
	var req *account.DepartmentReq
	err := c.ShouldBind(&req)
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &departmentRpc.DepartmentReqMessage{
		Page:                 req.Page,
		PageSize:             req.PageSize,
		ParentDepartmentCode: req.Pcode,
		OrganizationCode:     c.GetString("organizationCode"),
	}
	listDepartmentMessage, err := rpc.DepartmentClient.List(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	var list []*account.Department
	_ = copier.Copy(&list, listDepartmentMessage.List)
	if list == nil {
		list = []*account.Department{}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"total": listDepartmentMessage.Total,
		"page":  req.Page,
		"list":  list,
	}))
}
