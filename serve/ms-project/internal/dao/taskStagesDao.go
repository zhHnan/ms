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
func (t *TaskStagesDao) FindByProjectCode(ctx context.Context, projectCode int64, page int64, size int64) ([]*task.TaskStages, int64, error) {
	session := t.conn.Session(ctx)
	var stages []*task.TaskStages
	err := session.Model(&task.TaskStages{}).Where("project_code=? and deleted=?", projectCode, 0).Order("sort asc").Limit(int(size)).Offset(int((page - 1) * size)).Find(&stages).Error
	var total int64
	err = session.Model(&task.TaskStages{}).Where("project_code=?", projectCode).Count(&total).Error
	return stages, total, err
}
func (t *TaskStagesDao) FindById(ctx context.Context, stageCode int) (*task.TaskStages, error) {
	var stages task.TaskStages
	err := t.conn.Session(ctx).Model(&task.TaskStages{}).Where("id=?", stageCode).Find(&stages).Error
	return &stages, err
}
