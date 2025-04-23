package dao

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/project"
	"hnz.com/ms_serve/ms-project/internal/database/gorms"
)

type ProjectLogDao struct {
	conn *gorms.GormConn
}

func NewProjectLogDao() *ProjectLogDao {
	return &ProjectLogDao{
		conn: gorms.New(),
	}
}

func (p *ProjectLogDao) SaveProjectLog(pl *project.ProjectLog) {
	session := p.conn.Session(context.Background())
	session.Save(&pl)
}

func (p *ProjectLogDao) FindLogByTaskCode(ctx context.Context, taskCode int64, comment int) (list []*project.ProjectLog, total int64, err error) {
	session := p.conn.Session(ctx)
	model := session.Model(&project.ProjectLog{})
	if comment == 1 {
		model.Where("source_code=? and is_comment=?", taskCode, comment).Find(&list)
		model.Where("source_code=? and is_comment=?", taskCode, comment).Count(&total)
	} else {
		model.Where("source_code=?", taskCode).Find(&list)
		model.Where("source_code=?", taskCode).Count(&total)
	}
	return
}

func (p *ProjectLogDao) FindLogByTaskCodePage(ctx context.Context, taskCode int64, comment int, page int, pageSize int) (list []*project.ProjectLog, total int64, err error) {
	session := p.conn.Session(ctx)
	model := session.Model(&project.ProjectLog{})
	offset := (page - 1) * pageSize
	if comment == 1 {
		model.Where("source_code=? and is_comment=?", taskCode, comment).Limit(pageSize).Offset(offset).Find(&list)
		model.Where("source_code=? and is_comment=?", taskCode, comment).Count(&total)
	} else {
		model.Where("source_code=?", taskCode).Limit(pageSize).Offset(offset).Find(&list)
		model.Where("source_code=?", taskCode).Count(&total)
	}
	return
}
func (p *ProjectLogDao) FindLogByMemberCode(ctx context.Context, memberId int64, page int64, size int64) (list []*project.ProjectLog, total int64, err error) {
	session := p.conn.Session(ctx)
	offset := (page - 1) * size
	err = session.Model(&project.ProjectLog{}).
		Where("member_code=?", memberId).
		Limit(int(size)).
		Offset(int(offset)).Order("create_time desc").Find(&list).Error
	err = session.Model(&project.ProjectLog{}).
		Where("member_code=?", memberId).Count(&total).Error
	return
}
