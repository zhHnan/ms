package apiProject

type Project struct {
	Id                 int64   `json:"id"`
	Cover              string  `json:"cover"`
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	AccessControlType  string  `json:"access_control_type"`
	WhiteList          string  `json:"white_list"`
	Order              int     `json:"order"`
	Deleted            int     `json:"deleted"`
	TemplateCode       string  `json:"template_code"`
	Schedule           float64 `json:"schedule"`
	CreateTime         string  `json:"create_time"`
	OrganizationCode   string  `json:"organization_code"`
	DeletedTime        string  `json:"deleted_time"`
	Private            int     `json:"private"`
	Prefix             string  `json:"prefix"`
	OpenPrefix         int     `json:"open_prefix"`
	Archive            int     `json:"archive"`
	ArchiveTime        int64   `json:"archive_time"`
	OpenBeginTime      int     `json:"open_begin_time"`
	OpenTaskPrivate    int     `json:"open_task_private"`
	TaskBoardTheme     string  `json:"task_board_theme"`
	BeginTime          int64   `json:"begin_time"`
	EndTime            int64   `json:"end_time"`
	AutoUpdateSchedule int     `json:"auto_update_schedule"`
	Code               string  `json:"code"`
}

type MemberProject struct {
	Id          int64  `json:"id"`
	ProjectCode int64  `json:"project_code"`
	MemberCode  int64  `json:"member_code"`
	JoinTime    string `json:"join_time"`
	IsOwner     int64  `json:"is_owner"`
	Authorize   string `json:"authorize"`
}

type ProjectAndMember struct {
	Project
	ProjectCode int64  `json:"project_code"`
	MemberCode  int64  `json:"member_code"`
	JoinTime    int64  `json:"join_time"`
	IsOwner     int64  `json:"is_owner"`
	Authorize   string `json:"authorize"`
	OwnerName   string `json:"owner_name"`
	Collected   int    `json:"collected"`
}
type ProjectDetail struct {
	Project
	OwnerName   string `json:"owner_name"`
	OwnerAvatar string `json:"owner_avatar"`
	Collected   int    `json:"collected"`
}

type ProjectTemplate struct {
	Id               int                   `json:"id"`
	Name             string                `json:"name"`
	Description      string                `json:"description"`
	Sort             int                   `json:"sort"`
	CreateTime       string                `json:"create_time"`
	OrganizationCode string                `json:"organization_code"`
	Cover            string                `json:"cover"`
	MemberCode       string                `json:"member_code"`
	IsSystem         int                   `json:"is_system"`
	TaskStages       []*TaskStagesOnlyName `json:"task_stages"`
	Code             string                `json:"code"`
}

type TaskStagesOnlyName struct {
	Name string `json:"name"`
}

type SaveProjectRequest struct {
	Name         string `json:"name" form:"name"`
	TemplateCode string `json:"templateCode" form:"templateCode"`
	Description  string `json:"description" form:"description"`
	Id           int    `json:"id" form:"id"`
}

type SaveProject struct {
	Id               int64  `json:"id"`
	Cover            string `json:"cover"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Code             string `json:"code"`
	CreateTime       string `json:"create_time"`
	TaskBoardTheme   string `json:"task_board_theme"`
	OrganizationCode string `json:"organization_code"`
}

type ProjectRequest struct {
	ProjectCode        string `json:"project_code" form:"projectCode"`
	Name               string `json:"name" form:"name"`
	Description        string `json:"description" form:"description"`
	Cover              string `json:"cover" form:"cover"`
	WhiteList          string `json:"white_list" form:"white_list"`
	AccessControlType  string `json:"access_control_type" form:"access_control_type"`
	Private            int    `json:"private" form:"private"`
	Prefix             string `json:"prefix" form:"prefix"`
	TaskBoardTheme     string `json:"task_board_theme" form:"task_board_theme"`
	OpenPrefix         int    `json:"open_prefix" form:"open_prefix"`
	OpenBeginTime      int    `json:"open_begin_time" form:"open_begin_time"`
	OpenTaskPrivate    int    `json:"open_task_private" form:"open_task_private"`
	Schedule           int    `json:"schedule" form:"schedule"`
	AutoUpdateSchedule int    `json:"auto_update_schedule" form:"auto_update_schedule"`
}
type MemberProjectResp struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Avatar  string `json:"avatar"`
	Code    string `json:"code"`
	IsOwner int    `json:"isOwner"`
}
type ProjectLog struct {
	Content      string `json:"content"`
	Remark       string `json:"remark"`
	CreateTime   string `json:"create_time"`
	SourceCode   string `json:"source_code"`
	IsComment    int    `json:"is_comment"`
	ProjectCode  string `json:"project_code"`
	ProjectName  string `json:"project_name"`
	MemberAvatar string `json:"member_avatar"`
	MemberName   string `json:"member_name"`
	TaskName     string `json:"task_name"`
}
