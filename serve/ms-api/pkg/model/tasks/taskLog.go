package tasks

type TaskLogReq struct {
	TaskCode string `form:"taskCode"`
	PageSize int    `form:"pageSize"`
	Page     int    `form:"page"`
	All      int    `form:"all"`
	Comment  int    `form:"comment"`
}
type ProjectLogDisplay struct {
	Id           int64  `json:"id"`
	MemberCode   string `json:"member_code"`
	Content      string `json:"content"`
	Remark       string `json:"remark"`
	Type         string `json:"type"`
	CreateTime   string `json:"create_time"`
	SourceCode   string `json:"source_code"`
	ActionType   string `json:"action_type"`
	ToMemberCode string `json:"to_member_code"`
	IsComment    int    `json:"is_comment"`
	ProjectCode  string `json:"project_code"`
	Icon         string `json:"icon"`
	IsRobot      int    `json:"is_robot"`
	Member       Member `json:"member"`
}

type Member struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Avatar string `json:"avatar"`
}

type TaskWorkTime struct {
	Id         int64  `json:"id"`
	TaskCode   string `json:"task_code"`
	MemberCode string `json:"member_code"`
	CreateTime string `json:"create_time"`
	Content    string `json:"content"`
	BeginTime  string `json:"begin_time"`
	Num        int    `json:"num"`
	Code       string `json:"code"`
	Member     Member `json:"member"`
}

type SaveTaskWorkTimeReq struct {
	TaskCode  string `json:"task_code" form:"taskCode"`
	Content   string `form:"content"`
	Num       int    `form:"num"`
	BeginTime string `form:"beginTime"`
}
