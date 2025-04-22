package repo

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/task"
	"hnz.com/ms_serve/ms-project/internal/database"
)

type TaskRepo interface {
	FindTaskByStageCode(ctx context.Context, stageId int) ([]*task.Task, error)
	FindTaskMemberByTaskId(ctx context.Context, taskCode int64, memberCode int64) (*task.TaskMember, error)
	FindTaskMaxIdNum(ctx context.Context, projectCode int64) (int64, error)
	FindTaskSort(ctx context.Context, projectCode int64, stageCode int64) (int64, error)
	SaveTask(ctx context.Context, conn database.DBConn, ts *task.Task) error
	SaveTaskMember(ctx context.Context, conn database.DBConn, tm *task.TaskMember) error
	FindTaskById(ctx context.Context, taskCode int64) (*task.Task, error)
	UpdateTaskSort(ctx context.Context, conn database.DBConn, ts *task.Task) error
	FindTaskByStageCodeSmallSort(ctx context.Context, stageCode int, sort int) (ts *task.Task, err error)
	FindTaskByAssignTo(ctx context.Context, memberId int64, done int, page int64, pageSize int64) ([]*task.Task, int64, error)
	FindTaskByMemberCode(ctx context.Context, memberId int64, done int, page int64, pageSize int64) (tList []*task.Task, total int64, err error)
	FindTaskByCreateBy(ctx context.Context, memberId int64, done int, page int64, pageSize int64) (tList []*task.Task, total int64, err error)
}
