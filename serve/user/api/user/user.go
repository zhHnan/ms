package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hnz.com/ms_serve/common"
	"hnz.com/ms_serve/user/pkg/dao"
	"hnz.com/ms_serve/user/pkg/model"
	"hnz.com/ms_serve/user/pkg/repo"
	"log"
	"time"
)

type HandlerLogin struct {
	cache repo.Cache
}

func New() *HandlerLogin {
	return &HandlerLogin{
		cache: dao.Rc,
	}
}

// GetCaptcha 获取手机验证码
func (h *HandlerLogin) GetCaptcha(ctx *gin.Context) {
	res := &common.Result{}
	mobile := ctx.PostForm("mobile")
	if mobile == "" {
		res.Failure(model.NoLegalMobile, "手机号不能为空")
		return
	}
	// 生成验证码
	code := "123456"
	go func() {
		// 发送验证码
		time.Sleep(time.Second * 2)
		zap.L().Info("验证码发送成功！")
		c, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		err := h.cache.Put(c, "REGISTER_"+mobile, code, time.Minute*5)
		if err != nil {
			log.Println("验证码存入redis失败！", err)
		}
		log.Printf("手机号和验证码存入到redis中：REGISTER_%s:%s", mobile, code)
	}()
	ctx.JSON(200, res.Success(code))
}
