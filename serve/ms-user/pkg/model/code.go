package model

import (
	"hnz.com/ms_serve/common/errs"
)

var (
	NoLegalMobile = errs.NewError(2001, "手机号不合法")
)
