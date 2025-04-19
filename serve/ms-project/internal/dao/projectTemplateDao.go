package dao

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/project"
	"hnz.com/ms_serve/ms-project/internal/database/gorms"
)

type ProjectTemplateDao struct {
	conn *gorms.GormConn
}

func NewProjectTemplateDao() *ProjectTemplateDao {
	return &ProjectTemplateDao{
		conn: gorms.New(),
	}
}
func (p ProjectTemplateDao) FindProjectTemplateSystem(ctx context.Context, page int64, size int64) ([]project.ProjectTemplate, int64, error) {
	var pts []project.ProjectTemplate
	session := p.conn.Session(ctx)
	err := session.Where("is_system=?", 1).
		Model(&project.ProjectTemplate{}).
		Limit(int(size)).Offset(int((page - 1) * size)).
		Find(&pts).Error
	var total int64
	if err != nil {
		return nil, total, err
	}
	err = session.Model(&project.ProjectTemplate{}).Where("is_system=?", 1).Count(&total).Error
	return pts, total, err
}

func (p ProjectTemplateDao) FindProjectTemplateCustom(ctx context.Context, memId int64, organizationCode int64, page int64, size int64) ([]project.ProjectTemplate, int64, error) {
	var pts []project.ProjectTemplate
	session := p.conn.Session(ctx)
	err := session.
		Model(&project.ProjectTemplate{}).
		Where("is_system=? and member_code=? and organization_code=?", 0, memId, organizationCode).
		Limit(int(size)).
		Offset(int((page - 1) * size)).
		Find(&pts).Error
	var total int64
	if err != nil {
		return nil, total, err
	}
	err = session.Model(&project.ProjectTemplate{}).Where("is_system=? and member_code=? and organization_code=?", 0, memId, organizationCode).Count(&total).Error
	return pts, total, err
}

func (p ProjectTemplateDao) FindProjectTemplateAll(ctx context.Context, organizationCode int64, page int64, size int64) ([]project.ProjectTemplate, int64, error) {
	var pts []project.ProjectTemplate
	session := p.conn.Session(ctx)
	err := session.
		Model(&project.ProjectTemplate{}).
		Where("organization_code=?", organizationCode).
		Limit(int(size)).
		Offset(int((page - 1) * size)).
		Find(&pts).Error
	var total int64
	if err != nil {
		return nil, total, err
	}
	err = session.Model(&project.ProjectTemplate{}).Where("organization_code=?", organizationCode).Count(&total).Error
	return pts, total, err
}
