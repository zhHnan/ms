package repo

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/account"
)

type DepartmentRepo interface {
	FindDepartmentById(ctx context.Context, id int64) (*account.Department, error)
	ListDepartment(organizationCode int64, parentDepartmentCode int64, page int64, size int64) (list []*account.Department, total int64, err error)
}
