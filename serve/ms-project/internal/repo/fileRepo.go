package repo

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/files"
)

type FileRepo interface {
	Save(ctx context.Context, file *files.File) error
	FindByIds(background context.Context, ids []int64) (list []*files.File, err error)
}

type SourceLinkRepo interface {
	Save(ctx context.Context, link *files.SourceLink) error
	FindByTaskCode(ctx context.Context, taskCode int64) (list []*files.SourceLink, err error)
}
