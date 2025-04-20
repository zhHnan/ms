package dao

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/task"
	"hnz.com/ms_serve/ms-project/internal/database/gorms"
)

type TaskStagesTemplateDao struct {
	conn *gorms.GormConn
}

func NewTaskStagesTemplateDao() *TaskStagesTemplateDao {
	return &TaskStagesTemplateDao{
		conn: gorms.New(),
	}
}
func (t *TaskStagesTemplateDao) FindInProTemIds(ctx context.Context, ids []int) ([]task.MsTaskStagesTemplate, error) {
	var tsts []task.MsTaskStagesTemplate
	session := t.conn.Session(ctx)
	err := session.Model(&task.MsTaskStagesTemplate{}).Where("project_template_code in ?", ids).Find(&tsts).Error
	return tsts, err
}

func (t *TaskStagesTemplateDao) FindByProjectId(ctx context.Context, projectId int) (list []task.MsTaskStagesTemplate, err error) {
	session := t.conn.Session(ctx)
	err = session.Model(&task.MsTaskStagesTemplate{}).Where("project_template_code = ?", projectId).Order("sort desc, id asc").Find(&list).Error
	return
}
