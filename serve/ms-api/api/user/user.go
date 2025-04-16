package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"hnz.com/ms_serve/common"
	"hnz.com/ms_serve/common/errs"
	loginServiceV1 "hnz.com/ms_serve/ms-user/pkg/service/login_service.v1"
	"time"
)

type HandlerUser struct {
}

func New() *HandlerUser {

	return &HandlerUser{}
}

func (h *HandlerUser) GetCaptcha(ctx *gin.Context) {

	result := &common.Result{}
	mobile := ctx.Query("mobile")
	c, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	resp, err := UserClient.GetCaptcha(c, &loginServiceV1.CaptchaMessage{
		Mobile: mobile,
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(200, result.Failure(code, msg))
		return
	}
	ctx.JSON(200, result.Success(resp.Code))
}
