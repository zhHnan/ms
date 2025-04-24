package repo

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/account"
)

type DepartmentRepo interface {
	FindDepartmentById(ctx context.Context, id int64) (*account.Department, error)
}
