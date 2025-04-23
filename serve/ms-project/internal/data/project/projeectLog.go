package project

import (
	"github.com/jinzhu/copier"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/times"
	"hnz.com/ms_serve/ms-project/internal/data/common"
)

type ProjectLog struct {
	Id           int64
	MemberCode   int64
	Content      string
	Remark       string
	Type         string
	CreateTime   int64
	SourceCode   int64
	ActionType   string
	ToMemberCode int64
	IsComment    int
	ProjectCode  int64
	Icon         string
	IsRobot      int
}

func (*ProjectLog) TableName() string {
	return "ms_project_log"
}

type ProjectLogDisplay struct {
	Id           int64
	MemberCode   string
	Content      string
	Remark       string
	Type         string
	CreateTime   string
	SourceCode   string
	ActionType   string
	ToMemberCode string
	IsComment    int
	ProjectCode  string
	Icon         string
	IsRobot      int
	Member       common.Member
}

func (l *ProjectLog) ToDisplay() *ProjectLogDisplay {
	pd := &ProjectLogDisplay{}
	_ = copier.Copy(pd, l)
	pd.MemberCode = encrypts.EncryptNoErr(l.MemberCode)
	pd.ToMemberCode = encrypts.EncryptNoErr(l.ToMemberCode)
	pd.ProjectCode = encrypts.EncryptNoErr(l.ProjectCode)
	pd.CreateTime = times.FormatByMill(l.CreateTime)
	pd.SourceCode = encrypts.EncryptNoErr(l.SourceCode)
	return pd
}
