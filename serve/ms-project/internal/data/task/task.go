package task

import (
	"github.com/jinzhu/copier"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/times"
)

type MsTaskStagesTemplate struct {
	Id                  int
	Name                string
	ProjectTemplateCode int
	CreateTime          int64
	Sort                int
}

func (*MsTaskStagesTemplate) TableName() string {
	return "ms_task_stages_template"
}

type TaskStagesOnlyName struct {
	Name string
}

func CovertProjectMap(tsts []MsTaskStagesTemplate) map[int][]*TaskStagesOnlyName {
	var tss = make(map[int][]*TaskStagesOnlyName)
	for _, v := range tsts {
		ts := &TaskStagesOnlyName{}
		ts.Name = v.Name
		tss[v.ProjectTemplateCode] = append(tss[v.ProjectTemplateCode], ts)
	}
	return tss
}

type TaskStages struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	ProjectCode int64  `json:"project_code"`
	Sort        int    `json:"sort"`
	Description string `json:"description"`
	CreateTime  int64  `json:"create_time"`
	Deleted     int    `json:"deleted"`
}

func (*TaskStages) TableName() string {
	return "ms_task_stages"
}

func ToTaskStagesMap(tss []*TaskStages) map[int]*TaskStages {
	m := make(map[int]*TaskStages)
	for _, v := range tss {
		m[v.Id] = v
	}
	return m
}

// Task
type Task struct {
	Id            int64
	ProjectCode   int64
	Name          string
	Pri           int
	ExecuteStatus int
	Description   string
	CreateBy      int64
	DoneBy        int64
	DoneTime      int64
	CreateTime    int64
	AssignTo      int64
	Deleted       int
	StageCode     int
	TaskTag       string
	Done          int
	BeginTime     int64
	EndTime       int64
	RemindTime    int64
	Pcode         int64
	Sort          int
	Like          int
	Star          int
	DeletedTime   int64
	Private       int
	IdNum         int
	Path          string
	Schedule      int
	VersionCode   int64
	FeaturesCode  int64
	WorkTime      int
	Status        int
}

func (*Task) TableName() string {
	return "ms_task"
}

type TaskMember struct {
	Id         int64
	TaskCode   int64
	IsExecutor int
	MemberCode int64
	JoinTime   int64
	IsOwner    int
}

func (*TaskMember) TableName() string {
	return "ms_task_member"
}

const (
	Wait = iota
	Doing
	Done
	Pause
	Cancel
	Closed
)

func (t *Task) GetExecuteStatusStr() string {
	status := t.ExecuteStatus
	if status == Wait {
		return "wait"
	}
	if status == Doing {
		return "doing"
	}
	if status == Done {
		return "done"
	}
	if status == Pause {
		return "pause"
	}
	if status == Cancel {
		return "cancel"
	}
	if status == Closed {
		return "closed"
	}
	return ""
}

type TaskDisplay struct {
	Id            int64
	ProjectCode   string
	Name          string
	Pri           int
	ExecuteStatus string
	Description   string
	CreateBy      string
	DoneBy        string
	DoneTime      string
	CreateTime    string
	AssignTo      string
	Deleted       int
	StageCode     string
	TaskTag       string
	Done          int
	BeginTime     string
	EndTime       string
	RemindTime    string
	Pcode         string
	Sort          int
	Like          int
	Star          int
	DeletedTime   string
	Private       int
	IdNum         int
	Path          string
	Schedule      int
	VersionCode   string
	FeaturesCode  string
	WorkTime      int
	Status        int
	Code          string
	CanRead       int
	Executor      Executor
}
type Executor struct {
	Name   string
	Avatar string
	Code   string
}

func (t *Task) ToTaskDisplay() *TaskDisplay {
	td := &TaskDisplay{}
	copier.Copy(td, t)
	td.CreateTime = times.FormatByMill(t.CreateTime)
	td.DoneTime = times.FormatByMill(t.DoneTime)
	td.BeginTime = times.FormatByMill(t.BeginTime)
	td.EndTime = times.FormatByMill(t.EndTime)
	td.RemindTime = times.FormatByMill(t.RemindTime)
	td.DeletedTime = times.FormatByMill(t.DeletedTime)
	td.CreateBy = encrypts.EncryptNoErr(t.CreateBy)
	td.ProjectCode = encrypts.EncryptNoErr(t.ProjectCode)
	td.DoneBy = encrypts.EncryptNoErr(t.DoneBy)
	td.AssignTo = encrypts.EncryptNoErr(t.AssignTo)
	td.StageCode = encrypts.EncryptNoErr(int64(t.StageCode))
	td.Pcode = encrypts.EncryptNoErr(t.Pcode)
	td.VersionCode = encrypts.EncryptNoErr(t.VersionCode)
	td.FeaturesCode = encrypts.EncryptNoErr(t.FeaturesCode)
	td.ExecuteStatus = t.GetExecuteStatusStr()
	td.Code = encrypts.EncryptNoErr(t.Id)
	td.CanRead = 1
	return td
}
