package tasks

type TaskStagesResp struct {
	Name         string `json:"name"`
	ProjectCode  string `json:"project_code"`
	Sort         int    `json:"sort"`
	Description  string `json:"description"`
	CreateTime   string `json:"create_time"`
	Code         string `json:"code"`
	Deleted      int    `json:"deleted"`
	TasksLoading bool   `json:"tasksLoading"`
	FixedCreator bool   `json:"fixedCreator"`
	ShowTaskCard bool   `json:"showTaskCard"`
	Tasks        []int  `json:"tasks"`
	DoneTasks    []int  `json:"doneTasks"`
	UnDoneTasks  []int  `json:"unDoneTasks"`
}

type TaskDisplay struct {
	ProjectCode   string   `json:"project_code"`
	Name          string   `json:"name"`
	Pri           int      `json:"pri"`
	ExecuteStatus string   `json:"execute_status"`
	Description   string   `json:"description"`
	CreateBy      string   `json:"create_by"`
	DoneBy        string   `json:"done_by"`
	DoneTime      string   `json:"done_time"`
	CreateTime    string   `json:"create_time"`
	AssignTo      string   `json:"assign_to"`
	Deleted       int      `json:"deleted"`
	StageCode     string   `json:"stage_code"`
	TaskTag       string   `json:"task_tag"`
	Done          int      `json:"done"`
	BeginTime     string   `json:"begin_time"`
	EndTime       string   `json:"end_time"`
	RemindTime    string   `json:"remind_time"`
	Pcode         string   `json:"pcode"`
	Sort          int      `json:"sort"`
	Like          int      `json:"like"`
	Star          int      `json:"star"`
	DeletedTime   string   `json:"deleted_time"`
	Private       int      `json:"private"`
	IdNum         int      `json:"id_num"`
	Path          string   `json:"path"`
	Schedule      int      `json:"schedule"`
	VersionCode   string   `json:"version_code"`
	FeaturesCode  string   `json:"features_code"`
	WorkTime      int      `json:"work_time"`
	Status        int      `json:"status"`
	Code          string   `json:"code"`
	CanRead       int      `json:"canRead"`
	HasUnDone     int      `json:"hasUnDone"`
	ParentDone    int      `json:"parentDone"`
	HasComment    int      `json:"hasComment"`
	HasSource     int      `json:"hasSource"`
	Executor      Executor `json:"executor"`
	PriText       string   `json:"priText"`
	StatusText    string   `json:"statusText"`
	Liked         int      `json:"liked"`
	Stared        int      `json:"stared"`
	Tags          []int    `json:"tags"`
	ChildCount    []int    `json:"childCount"`
}

type Executor struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
type TaskSaveReq struct {
	Name        string `form:"name"`
	StageCode   string `form:"stage_code"`
	ProjectCode string `form:"project_code"`
	AssignTo    string `form:"assign_to"`
}
