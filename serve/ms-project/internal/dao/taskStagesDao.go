package dao

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/task"
	"hnz.com/ms_serve/ms-project/internal/database"
	"hnz.com/ms_serve/ms-project/internal/database/gorms"
)

type TaskStagesDao struct {
	conn *gorms.GormConn
}

func NewTaskStagesDao() *TaskStagesDao {
	return &TaskStagesDao{
		conn: gorms.New(),
	}
}

func (t *TaskStagesDao) SaveTaskStages(conn database.DBConn, ctx context.Context, msg *task.TaskStages) error {
	t.conn = conn.(*gorms.GormConn)
	err := t.conn.Tx(ctx).Save(&msg).Error
	return err
}
