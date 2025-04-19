package midd

import (
	"context"
	"github.com/gin-gonic/gin"
	"hnz.com/ms_serve/ms-api/api/rpc"
	common "hnz.com/ms_serve/ms-common"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-grpc/user/login"
	"time"
)

func TokenVerify() func(c *gin.Context) {
	return func(c *gin.Context) {
		result := &common.Result{}
		token := c.GetHeader("Authorization")
		//验证用户是否已经登录
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		resp, err := rpc.UserClient.TokenVerify(ctx, &login.LoginMessage{Token: token})
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(200, result.Failure(code, msg))
			c.Abort()
			return
		}
		c.Set("memberId", resp.Member.Id)
		c.Next()
	}
}
