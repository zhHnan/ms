package repo

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/task"
	"hnz.com/ms_serve/ms-project/internal/database"
)

type TaskStagesTemplateRepo interface {
	FindInProTemIds(ctx context.Context, id []int) ([]task.MsTaskStagesTemplate, error)
	FindByProjectId(ctx context.Context, projectId int) ([]task.MsTaskStagesTemplate, error)
}

type TaskStagesRepo interface {
	SaveTaskStages(conn database.DBConn, ctx context.Context, msg *task.TaskStages) error
}
