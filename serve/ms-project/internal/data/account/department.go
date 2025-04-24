package account

import (
	"github.com/jinzhu/copier"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/times"
)

type Department struct {
	Id               int64  `gorm:"column:id"`
	OrganizationCode int64  `gorm:"column:organization_code"`
	Name             string `gorm:"column:name"`
	Sort             int    `gorm:"column:sort"`
	PCode            int64  `gorm:"column:pcode"` // 显式指定列名为 p_code
	Icon             string `gorm:"column:icon"`
	CreateTime       int64  `gorm:"column:create_time"`
	Path             string `gorm:"column:path"`
}

func (*Department) TableName() string {
	return "ms_department"
}

type DepartmentDisplay struct {
	Id               int64
	OrganizationCode string
	Name             string
	Sort             int
	Pcode            string
	icon             string
	CreateTime       string
	Path             string
}

func (d *Department) ToDisplay() *DepartmentDisplay {
	dp := &DepartmentDisplay{}
	_ = copier.Copy(dp, d)
	dp.CreateTime = times.FormatByMill(d.CreateTime)
	dp.OrganizationCode = encrypts.EncryptNoErr(d.OrganizationCode)
	if d.PCode > 0 {
		dp.Pcode = encrypts.EncryptNoErr(d.PCode)
	} else {
		dp.Pcode = ""
	}
	return dp
}
