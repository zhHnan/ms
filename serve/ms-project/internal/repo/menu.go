package repo

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/menu"
)

type MenuRepo interface {
	FindAll(ctx context.Context) ([]*menu.ProjectMenu, error)
}
