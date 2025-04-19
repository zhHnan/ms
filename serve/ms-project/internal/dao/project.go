package dao

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/project"
	"hnz.com/ms_serve/ms-project/internal/database/gorms"
)

type ProjectDao struct {
	conn *gorms.GormConn
}

func NewProjectDao() *ProjectDao {
	return &ProjectDao{
		conn: gorms.New(),
	}
}

func (p *ProjectDao) FindProjectByMemId(ctx context.Context, memId int64, page int64, size int64) ([]*project.ProjectAndMember, int64, error) {
	session := p.conn.Session(ctx)
	index := (page - 1) * size
	db := session.Raw("select * from ms_project a, ms_project_member b where a.id=b.project_code and b.member_code=? limit ?,?", memId, index, size)
	var mp []*project.ProjectAndMember
	err := db.Scan(&mp).Error
	var total int64
	session.Model(&project.ProjectMember{}).Where("member_code=?", memId).Count(&total)
	return mp, total, err
}
