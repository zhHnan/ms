package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"hnz.com/ms_serve/ms-project/internal/data/task"
	"hnz.com/ms_serve/ms-project/internal/database/gorms"
)

type TaskDao struct {
	conn *gorms.GormConn
}

func NewTaskDao() *TaskDao {
	return &TaskDao{
		conn: gorms.New(),
	}
}

func (t *TaskDao) FindTaskByStageCode(ctx context.Context, stageCode int) ([]*task.Task, error) {
	var taskList []*task.Task
	session := t.conn.Session(ctx)
	err := session.Model(&task.Task{}).Where("stage_code=? and deleted = 0", stageCode).
		Order("sort asc").
		Find(&taskList).Error
	return taskList, err
}

func (t *TaskDao) FindTaskMemberByTaskId(ctx context.Context, taskCode int64, memberCode int64) (*task.TaskMember, error) {
	var tm *task.TaskMember
	err := t.conn.Session(ctx).Where("task_code=? and member_code=?", taskCode, memberCode).Limit(1).Find(&tm).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return tm, err
}
