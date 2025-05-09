package task

import (
	"github.com/jinzhu/copier"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/times"
	"hnz.com/ms_serve/ms-project/internal/data/common"
	"hnz.com/ms_serve/ms-project/internal/data/project"
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

func CovertProjectMap(tsts []MsTaskStagesTemplate) map[int][]*common.TaskStagesOnlyName {
	var tss = make(map[int][]*common.TaskStagesOnlyName)
	for _, v := range tsts {
		ts := &common.TaskStagesOnlyName{}
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
	ProjectName   string
	StageName     string
	PriText       string
	StatusText    string
}
type Executor struct {
	Name   string
	Avatar string
	Code   string
}

func (t *Task) ToTaskDisplay() *TaskDisplay {
	td := &TaskDisplay{}
	_ = copier.Copy(td, t)
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
	td.StatusText = t.GetStatusStr()
	td.PriText = t.GetPriStr()
	return td
}

type MyTaskDisplay struct {
	Id                 int64
	ProjectCode        string
	Name               string
	Pri                int
	ExecuteStatus      string
	Description        string
	CreateBy           string
	DoneBy             string
	DoneTime           string
	CreateTime         string
	AssignTo           string
	Deleted            int
	StageCode          string
	TaskTag            string
	Done               int
	BeginTime          string
	EndTime            string
	RemindTime         string
	Pcode              string
	Sort               int
	Like               int
	Star               int
	DeletedTime        string
	Private            int
	IdNum              int
	Path               string
	Schedule           int
	VersionCode        string
	FeaturesCode       string
	WorkTime           int
	Status             int
	Code               string
	Cover              string `json:"cover"`
	AccessControlType  string `json:"access_control_type"`
	WhiteList          string `json:"white_list"`
	Order              int    `json:"order"`
	TemplateCode       string `json:"template_code"`
	OrganizationCode   string `json:"organization_code"`
	Prefix             string `json:"prefix"`
	OpenPrefix         int    `json:"open_prefix"`
	Archive            int    `json:"archive"`
	ArchiveTime        string `json:"archive_time"`
	OpenBeginTime      int    `json:"open_begin_time"`
	OpenTaskPrivate    int    `json:"open_task_private"`
	TaskBoardTheme     string `json:"task_board_theme"`
	AutoUpdateSchedule int    `json:"auto_update_schedule"`
	HasUnDone          int    `json:"hasUnDone"`
	ParentDone         int    `json:"parentDone"`
	PriText            string `json:"priText"`
	ProjectName        string
	Executor           *Executor
}

func (t *Task) ToMyTaskDisplay(p *project.Project, name string, avatar string) *MyTaskDisplay {
	td := &MyTaskDisplay{}
	_ = copier.Copy(td, p)
	_ = copier.Copy(td, t)
	td.Executor = &Executor{
		Name:   name,
		Avatar: avatar,
	}
	td.ProjectName = p.Name
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
	td.AccessControlType = GetAccessControlType(p.AccessControlType)
	td.ArchiveTime = times.FormatByMill(p.ArchiveTime)
	td.TemplateCode = encrypts.EncryptNoErr(int64(p.TemplateCode))
	td.OrganizationCode = encrypts.EncryptNoErr(p.OrganizationCode)
	return td
}

func GetAccessControlType(accessControlType int) string {
	if accessControlType == 0 {
		return "open"
	}
	if accessControlType == 1 {
		return "private"
	}
	if accessControlType == 2 {
		return "custom"
	}
	return ""
}
func (t *Task) GetStatusStr() string {
	status := t.Status
	if status == NoStarted {
		return "未开始"
	}
	if status == Started {
		return "开始"
	}
	return ""
}
func (t *Task) GetPriStr() string {
	status := t.Pri
	if status == Normal {
		return "普通"
	}
	if status == Urgent {
		return "紧急"
	}
	if status == VeryUrgent {
		return "非常紧急"
	}
	return ""
}

const (
	Normal = iota
	Urgent
	VeryUrgent
)

const (
	NoStarted = iota
	Started
)
