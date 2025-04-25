package repo

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/node"
)

type ProjectNodeRepo interface {
	FindAll(ctx context.Context) (pms []*node.ProjectNode, err error)
}
