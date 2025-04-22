package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"hnz.com/ms_serve/ms-api/api/rpc"
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
	resp, err := rpc.UserClient.GetCaptcha(c, &login.CaptchaMessage{
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
	if err := c.ShouldBind(&req); err != nil {
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
	_, err := rpc.UserClient.Register(ctx, msg)
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
	//// 读取原始请求体
	//body, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	log.Printf("Failed to read request body: %v", err)
	//	c.JSON(http.StatusBadRequest, result.Failure(http.StatusBadRequest, "无法读取请求体"))
	//	return
	//}
	//log.Printf("Raw request body: %s", string(body))
	//
	//// 重新设置请求体以便后续解析
	//c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	var req user.LoginReq
	if err := c.ShouldBind(&req); err != nil {
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
	ip := GetIp(c)
	msg.Ip = ip
	res, err := rpc.UserClient.Login(ctx, msg)
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

func (p *HandlerUser) myOrgList(c *gin.Context) {
	result := &common.Result{}
	token := c.GetHeader("Authorization")
	//验证用户是否已经登录
	mem, err2 := rpc.UserClient.TokenVerify(context.Background(), &login.LoginMessage{Token: token})
	if err2 != nil {
		code, msg := errs.ParseGrpcError(err2)
		c.JSON(http.StatusOK, result.Failure(code, msg))
		return
	}
	list, err2 := rpc.UserClient.MyOrgList(context.Background(), &login.UserMessage{MemId: mem.Member.Id})
	if err2 != nil {
		code, msg := errs.ParseGrpcError(err2)
		c.JSON(http.StatusOK, result.Failure(code, msg))
		return
	}
	if list.OrganizationList == nil {
		c.JSON(http.StatusOK, result.Success([]*user.OrganizationList{}))
		return
	}
	var orgs []*user.OrganizationList
	_ = copier.Copy(&orgs, list.OrganizationList)
	c.JSON(http.StatusOK, result.Success(orgs))
}

// GetIp 获取ip函数
func GetIp(c *gin.Context) string {
	ip := c.ClientIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	return ip
}
