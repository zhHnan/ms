package repo

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/task"
)

type TaskWorkTimeRepo interface {
	Save(ctx context.Context, twt *task.TaskWorkTime) error
	FindWorkTimeList(ctx context.Context, taskCode int64) (list []*task.TaskWorkTime, err error)
}
