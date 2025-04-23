package task

import (
	"github.com/jinzhu/copier"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/times"
	"hnz.com/ms_serve/ms-project/internal/data/common"
)

type TaskWorkTime struct {
	Id         int64
	TaskCode   int64
	MemberCode int64
	CreateTime int64
	Content    string
	BeginTime  int64
	Num        int
}

func (*TaskWorkTime) TableName() string {
	return "ms_task_work_time"
}

type TaskWorkTimeDisplay struct {
	Id         int64
	TaskCode   string
	MemberCode string
	CreateTime string
	Content    string
	BeginTime  string
	Num        int
	Member     common.Member
}

func (t *TaskWorkTime) ToDisplay() *TaskWorkTimeDisplay {
	td := &TaskWorkTimeDisplay{}
	_ = copier.Copy(td, t)
	td.MemberCode = encrypts.EncryptNoErr(t.MemberCode)
	td.TaskCode = encrypts.EncryptNoErr(t.TaskCode)
	td.CreateTime = times.FormatByMill(t.CreateTime)
	td.BeginTime = times.FormatByMill(t.BeginTime)
	return td
}
