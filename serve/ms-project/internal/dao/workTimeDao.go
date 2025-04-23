package dao

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/task"
	"hnz.com/ms_serve/ms-project/internal/database/gorms"
)

type TaskWorkTimeDao struct {
	conn *gorms.GormConn
}

func NewTaskWorkTimeDao() *TaskWorkTimeDao {
	return &TaskWorkTimeDao{
		conn: gorms.New(),
	}
}

func (t *TaskWorkTimeDao) Save(ctx context.Context, twt *task.TaskWorkTime) error {
	session := t.conn.Session(ctx)
	err := session.Save(&twt).Error
	return err
}

func (t *TaskWorkTimeDao) FindWorkTimeList(ctx context.Context, taskCode int64) (list []*task.TaskWorkTime, err error) {
	session := t.conn.Session(ctx)
	err = session.Model(&task.TaskWorkTime{}).Where("task_code=?", taskCode).Find(&list).Error
	return
}
