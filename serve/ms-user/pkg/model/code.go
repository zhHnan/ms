package model

import (
	"hnz.com/ms_serve/ms-common/errs"
)

var (
	RedisError    = errs.NewError(9001, "redis错误")
	DataBaseError = errs.NewError(9002, "数据库错误")
	NoLegalMobile = errs.NewError(2001, "手机号不合法")
	CaptchaError  = errs.NewError(2002, "验证码不正确")
	EmailExist    = errs.NewError(2003, "邮箱已存在")
	MobileExist   = errs.NewError(2004, "手机号已存在")
	NameExist     = errs.NewError(2005, "用户名已存在")
)
