package repo

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/project"
)

type ProjectLogRepo interface {
	FindLogByTaskCode(ctx context.Context, taskCode int64, comment int) (list []*project.ProjectLog, total int64, err error)
	FindLogByTaskCodePage(ctx context.Context, taskCode int64, comment int, page int, pageSize int) (list []*project.ProjectLog, total int64, err error)
	SaveProjectLog(pl *project.ProjectLog)
}
