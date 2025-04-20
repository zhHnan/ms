package task

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
