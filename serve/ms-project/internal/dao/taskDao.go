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

func (t *TaskDao) FindTaskByStageCodeSmallSort(ctx context.Context, stageCode int, sort int) (ts *task.Task, err error) {
	session := t.conn.Session(ctx)
	err = session.Where("stage_code=? and sort < ?", stageCode, sort).Order("sort desc").Limit(1).Find(&ts).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}
func (t *TaskDao) FindTaskByCreateBy(ctx context.Context, memberId int64, done int, page int64, pageSize int64) (tList []*task.Task, total int64, err error) {
	session := t.conn.Session(ctx)
	offset := (page - 1) * pageSize
	err = session.Model(&task.Task{}).Where("create_by=? and deleted=0 and done=?", memberId, done).Limit(int(pageSize)).Offset(int(offset)).Find(&tList).Error
	err = session.Model(&task.Task{}).Where("create_by=? and deleted=0 and done=?", memberId, done).Count(&total).Error
	return
}

func (t *TaskDao) FindTaskByMemberCode(ctx context.Context, memberId int64, done int, page int64, pageSize int64) (tList []*task.Task, total int64, err error) {
	session := t.conn.Session(ctx)
	offset := (page - 1) * pageSize
	sql := "select a.* from ms_task a,ms_task_member b where a.id=b.task_code and member_code=? and a.deleted=0 and a.done=? limit ?,?"
	raw := session.Model(&task.Task{}).Raw(sql, memberId, done, offset, pageSize)
	err = raw.Scan(&tList).Error
	if err != nil {
		return nil, 0, err
	}
	sqlCount := "select count(*) from ms_task a,ms_task_member b where a.id=b.task_code and member_code=? and a.deleted=0 and a.done=?"
	rawCount := session.Model(&task.Task{}).Raw(sqlCount, memberId, done)
	err = rawCount.Scan(&total).Error
	return
}

func (t *TaskDao) FindTaskByAssignTo(ctx context.Context, memberId int64, done int, page int64, pageSize int64) (tsList []*task.Task, total int64, err error) {
	session := t.conn.Session(ctx)
	offset := (page - 1) * pageSize
	err = session.Model(&task.Task{}).Where("assign_to=? and deleted=0 and done=?", memberId, done).Limit(int(pageSize)).Offset(int(offset)).Find(&tsList).Error
	err = session.Model(&task.Task{}).Where("assign_to=? and deleted=0 and done=?", memberId, done).Count(&total).Error
	return
}
