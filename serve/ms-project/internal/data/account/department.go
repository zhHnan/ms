package account

type Department struct {
	Id               int64
	OrganizationCode int64
	Name             string
	Sort             int
	PCode            int64
	icon             string
	CreateTime       int64
	Path             string
}

func (*Department) TableName() string {
	return "ms_department"
}
