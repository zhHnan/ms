package midd

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hnz.com/ms_serve/ms-api/api/rpc"
	common "hnz.com/ms_serve/ms-common"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-grpc/user/login"
)

func TokenVerify() func(c *gin.Context) {
	return func(c *gin.Context) {
		result := &common.Result{}
		token := c.GetHeader("Authorization")
		//验证用户是否已经登录
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		ip := GetIp(c)
		resp, err := rpc.UserClient.TokenVerify(ctx, &login.LoginMessage{Token: token, Ip: ip})
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(200, result.Failure(code, msg))
			c.Abort()
			return
		}
		c.Set("memberId", resp.Member.Id)
		c.Set("memberName", resp.Member.Name)
		c.Set("organizationCode", resp.Member.OrganizationCode)
		c.Next()
	}
}

// GetIp 获取ip函数
func GetIp(c *gin.Context) string {
	ip := c.ClientIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	return ip
}
func RequestLog() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 记录毫秒级别时间
		start := time.Now()
		c.Next()
		diff := time.Since(start).Milliseconds()
		zap.L().Info("请求信息",
			zap.String("ip", GetIp(c)),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("cost", time.Duration(diff)*time.Millisecond),
		)
	}
}
