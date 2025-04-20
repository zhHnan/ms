package dao

import (
	"context"
	"fmt"
	"hnz.com/ms_serve/ms-project/internal/data/project"
	"hnz.com/ms_serve/ms-project/internal/database"
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

func (p *ProjectDao) FindProjectByMemId(ctx context.Context, memId int64, condition string, page int64, size int64) ([]*project.ProjectAndMember, int64, error) {
	session := p.conn.Session(ctx)
	index := (page - 1) * size
	sql := fmt.Sprintf("select * from ms_project a, ms_project_member b where a.id=b.project_code and b.member_code=? %s order by sort limit ?,?", condition)
	db := session.Raw(sql, memId, index, size)
	var mp []*project.ProjectAndMember
	err := db.Scan(&mp).Error
	var total int64
	query := fmt.Sprintf("select count(*) from ms_project a, ms_project_member b where a.id=b.project_code and b.member_code=? %s", condition)
	err = session.Raw(query, memId).Scan(&total).Error
	return mp, total, err
}
func (p *ProjectDao) FindCollectProjectByMemId(ctx context.Context, id int64, page int64, size int64) ([]*project.ProjectAndMember, int64, error) {
	session := p.conn.Session(ctx)
	index := (page - 1) * size
	sql := fmt.Sprintf("select * from ms_project where id in (select project_code from ms_project_collection where member_code = ?) order by sort limit ?,?")
	db := session.Raw(sql, id, index, size)
	var mp []*project.ProjectAndMember
	err := db.Scan(&mp).Error
	var total int64
	query := fmt.Sprintf("member_code=?")
	session.Model(&project.ProjectMember{}).Where(query, id).Count(&total)
	return mp, total, err
}
func (p *ProjectDao) SaveProject(conn database.DBConn, ctx context.Context, pr *project.Project) error {
	p.conn = conn.(*gorms.GormConn)
	return p.conn.Tx(ctx).Save(&pr).Error
}

func (p *ProjectDao) SaveProjectMember(conn database.DBConn, ctx context.Context, pm *project.ProjectMember) error {
	p.conn = conn.(*gorms.GormConn)
	return p.conn.Tx(ctx).Save(&pm).Error
}
