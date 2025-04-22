package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"hnz.com/ms_serve/ms-project/internal/data/task"
	"hnz.com/ms_serve/ms-project/internal/database"
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
	if err != nil {
		return nil, err
	}
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
func (t *TaskDao) FindTaskMaxIdNum(ctx context.Context, projectCode int64) (v int64, err error) {
	session := t.conn.Session(ctx)
	err = session.Model(&task.Task{}).
		Where("project_code=?", projectCode).
		Select("COALESCE(max(id_num), 0) as maxIdNum").
		Scan(&v).Error
	return
}

func (t *TaskDao) FindTaskSort(ctx context.Context, projectCode int64, stageCode int64) (v int64, err error) {
	session := t.conn.Session(ctx)
	err = session.Model(&task.Task{}).
		Where("project_code=? and stage_code=?", projectCode, stageCode).
		Select("COALESCE(max(sort), 0) as sort").
		Scan(&v).Error
	return
}
func (t *TaskDao) SaveTask(ctx context.Context, conn database.DBConn, ts *task.Task) error {
	t.conn = conn.(*gorms.GormConn)
	return t.conn.Tx(ctx).Save(&ts).Error
}

func (t *TaskDao) SaveTaskMember(ctx context.Context, conn database.DBConn, tm *task.TaskMember) error {
	t.conn = conn.(*gorms.GormConn)
	return t.conn.Tx(ctx).Save(&tm).Error
}
func (t *TaskDao) FindTaskById(ctx context.Context, taskCode int64) (ts *task.Task, err error) {
	session := t.conn.Session(ctx)
	err = session.Where("id=?", taskCode).Take(&ts).Error
	return
}

func (t *TaskDao) UpdateTaskSort(ctx context.Context, conn database.DBConn, ts *task.Task) error {
	t.conn = conn.(*gorms.GormConn)
	err := t.conn.Tx(ctx).Model(&task.Task{}).
		Where("id=?", ts.Id).
		Select("sort", "stage_code").
		Updates(&ts).
		Error
	return err
}
