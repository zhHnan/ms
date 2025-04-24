package dao

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/account"
	"hnz.com/ms_serve/ms-project/internal/database/gorms"
)

type DepartmentDao struct {
	conn *gorms.GormConn
}

func NewDepartmentDao() *DepartmentDao {
	return &DepartmentDao{
		conn: gorms.New(),
	}
}

func (d *DepartmentDao) FindDepartmentById(ctx context.Context, id int64) (dt *account.Department, err error) {
	session := d.conn.Session(ctx)
	err = session.Where("id=?", id).Find(&dt).Error
	return
}

func (d *DepartmentDao) ListDepartment(organizationCode int64, parentDepartmentCode int64, page int64, size int64) (list []*account.Department, total int64, err error) {
	session := d.conn.Session(context.Background())
	session.Model(&account.Department{})
	session.Where("organization_code=?", organizationCode)
	if parentDepartmentCode > 0 {
		session.Where("pcode=?", parentDepartmentCode)
	}
	err = session.Count(&total).Error
	err = session.Limit(int(size)).Offset(int((page - 1) * size)).Find(&list).Error
	return
}
