package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"hnz.com/ms_serve/ms-api/api/rpc"
	"hnz.com/ms_serve/ms-api/pkg/model/account"
	common "hnz.com/ms_serve/ms-common"
	"hnz.com/ms_serve/ms-common/errs"
	accountRpc "hnz.com/ms_serve/ms-grpc/account"
	"net/http"
	"time"
)

type HandlerAccount struct {
}

func NewAccount() *HandlerAccount {

	return &HandlerAccount{}
}

func (a *HandlerAccount) account(c *gin.Context) {
	result := &common.Result{}
	var req account.AccountReq
	err := c.ShouldBind(&req)
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &accountRpc.AccountReqMessage{
		MemberId:         c.GetInt64("memberId"),
		OrganizationCode: c.GetString("organizationCode"),
		Page:             int64(req.Page),
		PageSize:         int64(req.PageSize),
		SearchType:       int32(req.SearchType),
		DepartmentCode:   req.DepartmentCode,
	}
	response, err := rpc.AccountClient.Account(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	var list []*account.MemberAccount
	_ = copier.Copy(&list, response.AccountList)
	if list == nil {
		list = []*account.MemberAccount{}
	}
	var authList []*account.ProjectAuth
	_ = copier.Copy(&authList, response.AuthList)
	if authList == nil {
		authList = []*account.ProjectAuth{}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"total":    response.Total,
		"page":     req.Page,
		"list":     list,
		"authList": authList,
	}))
}
