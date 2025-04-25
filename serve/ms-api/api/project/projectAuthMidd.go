package project

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	common "hnz.com/ms_serve/ms-common"
	"hnz.com/ms_serve/ms-common/errs"
	"net/http"
)

func ProjectAuth() func(*gin.Context) {
	return func(c *gin.Context) {
		zap.L().Info("开始做项目操作授权认证")
		result := &common.Result{}
		//在接口有权限的基础上，做项目权限，不是这个项目的成员，无权限查看项目和操作项目
		//检查是否有projectCode和taskCode这两个参数
		isProjectAuth := false
		projectCode := c.PostForm("projectCode")
		if projectCode != "" {
			isProjectAuth = true
		}
		taskCode := c.PostForm("taskCode")
		if taskCode != "" {
			isProjectAuth = true
		}
		if isProjectAuth {
			p := New()
			pr, isMember, isOwner, err := p.FindProjectByMemberId(c.GetInt64("memberId"), projectCode, taskCode)
			if err != nil {
				code, msg := errs.ParseGrpcError(err)
				c.JSON(http.StatusOK, result.Failure(code, msg))
				c.Abort()
				return
			}
			if !isMember {
				c.JSON(http.StatusOK, result.Failure(403, "不是项目成员，无操作权限"))
				c.Abort()
				return
			}
			if pr.Private == 1 {
				//私有项目
				if isOwner {
					c.Next()
					return
				} else {
					c.JSON(http.StatusOK, result.Failure(403, "私有项目，无操作权限"))
					c.Abort()
					return
				}
			}
		}
	}
}
