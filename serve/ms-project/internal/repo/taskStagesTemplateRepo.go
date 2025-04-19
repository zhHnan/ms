package repo

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/task"
)

type TaskStagesTemplateRepo interface {
	FindInProTemIds(ctx context.Context, id []int) ([]task.MsTaskStagesTemplate, error)
}
