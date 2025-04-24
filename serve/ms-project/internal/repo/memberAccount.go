package repo

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/account"
)

type MemberAccountRepo interface {
	FindList(ctx context.Context, condition string, organizationCode int64, departmentCode int64, page int64, pageSize int64) ([]*account.MemberAccount, int64, error)
}
