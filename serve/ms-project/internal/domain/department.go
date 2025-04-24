package domain

import (
	"context"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/data/account"
	"hnz.com/ms_serve/ms-project/internal/repo"
	"hnz.com/ms_serve/ms-project/pkg/model"
	"time"
)

type DepartmentDomain struct {
	departmentRepo repo.DepartmentRepo
}

func (d *DepartmentDomain) FindDepartmentById(id int64) (*account.Department, error) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return d.departmentRepo.FindDepartmentById(c, id)
}

func (d *DepartmentDomain) List(organizationCode int64, parentDepartmentCode int64, page int64, size int64) ([]*account.DepartmentDisplay, int64, *errs.BError) {
	list, total, err := d.departmentRepo.ListDepartment(organizationCode, parentDepartmentCode, page, size)
	if err != nil {
		return nil, 0, model.DataBaseError
	}
	var dList []*account.DepartmentDisplay
	for _, v := range list {
		dList = append(dList, v.ToDisplay())
	}
	return dList, total, nil
}

func NewDepartmentDomain() *DepartmentDomain {
	return &DepartmentDomain{
		departmentRepo: dao.NewDepartmentDao(),
	}
}
