package repo

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/task"
)

type TaskRepo interface {
	FindTaskByStageCode(ctx context.Context, stageId int) ([]*task.Task, error)
	FindTaskMemberByTaskId(ctx context.Context, taskCode int64, memberCode int64) (*task.TaskMember, error)
}
