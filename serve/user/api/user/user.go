package user

import (
	"github.com/gin-gonic/gin"
	"hnz.com/ms_serve/common"
)

type HandlerLogin struct {
}

// GetCaptcha 获取手机验证码
func (HandlerLogin) GetCaptcha(ctx *gin.Context) {
	res := &common.Result{}
	ctx.JSON(200, res.Success("123456"))
}
