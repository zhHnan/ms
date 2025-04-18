package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"hnz.com/ms_serve/ms-api/pkg/model/user"
	common "hnz.com/ms_serve/ms-common"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-grpc/user/login"
	"net/http"
	"time"
)

type HandlerUser struct {
}

func New() *HandlerUser {

	return &HandlerUser{}
}

func (h *HandlerUser) getCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	mobile := ctx.PostForm("mobile")
	c, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	resp, err := UserClient.GetCaptcha(c, &login.CaptchaMessage{
		Mobile: mobile,
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(200, result.Failure(code, msg))
		return
	}
	ctx.JSON(200, result.Success(resp.Code))
}

// Register 注册
func (h *HandlerUser) register(c *gin.Context) {
	// 接收参数
	result := &common.Result{}
	var req user.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, result.Failure(http.StatusBadRequest, "参数格式有误"))
		return
	}
	// 校验参数
	if err := req.Verify(); err != nil {
		c.JSON(http.StatusOK, result.Failure(http.StatusBadRequest, err.Error()))
		return
	}
	// 调用rpc
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	msg := &login.RegisterMessage{}
	if err := copier.Copy(msg, &req); err != nil {
		c.JSON(http.StatusOK, result.Failure(http.StatusBadRequest, "参数格式有误"))
		return
	}
	_, err := UserClient.Register(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, result.Success(nil))
}

// Login 登录
func (h *HandlerUser) login(c *gin.Context) {
	// 接收参数
	result := &common.Result{}
	var req user.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, result.Failure(http.StatusBadRequest, "参数格式有误"))
		return
	}
	// 调用grpc 完成登录
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	msg := &login.LoginMessage{}
	if err := copier.Copy(msg, &req); err != nil {
		c.JSON(http.StatusOK, result.Failure(http.StatusBadRequest, "参数格式有误"))
		return
	}
	res, err := UserClient.Login(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Failure(code, msg))
		return
	}
	var resp user.LoginRsp
	err = copier.Copy(&resp, res)
	// 返回结果
	c.JSON(http.StatusOK, result.Success(resp))
}
