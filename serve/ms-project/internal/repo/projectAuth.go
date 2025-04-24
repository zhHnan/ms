package repo

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/account"
)

type ProjectAuthRepo interface {
	FindAuthList(ctx context.Context, orgCode int64) ([]*account.ProjectAuth, error)
}
