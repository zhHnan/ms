package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"hnz.com/ms_serve/ms-api/api/rpc"
	"hnz.com/ms_serve/ms-api/pkg/model/menu"
	common "hnz.com/ms_serve/ms-common"
	"hnz.com/ms_serve/ms-common/errs"
	menuRpc "hnz.com/ms_serve/ms-grpc/menu"
	"net/http"
	"time"
)

type HandlerMenu struct {
}

func NewMenu() *HandlerMenu {
	return &HandlerMenu{}
}

func (d *HandlerMenu) menuList(c *gin.Context) {
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := rpc.MenuClient.MenuList(ctx, &menuRpc.MenuReqMessage{})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
	}
	var list []*menu.Menu
	_ = copier.Copy(&list, res.List)
	if list == nil {
		list = []*menu.Menu{}
	}
	c.JSON(http.StatusOK, result.Success(list))
}
