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
	sql := fmt.Sprintf("select * from ms_project mp, ms_project_member mpm where mp.id = mpm.project_code and mp.id in (select project_code from ms_project_collection mpc where mpc.member_code = ?) order by sort limit ?,?")
	db := session.Raw(sql, id, index, size)
	var mp []*project.ProjectAndMember
	err := db.Scan(&mp).Error
	var total int64
	query := fmt.Sprintf("select count(*) from ms_project where id in (select project_code from ms_project_collection where member_code = ?)")
	err = session.Raw(query, id).Scan(&total).Error
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
func (p *ProjectDao) FindProjectByPIdAndMemId(ctx context.Context, code int64, id int64) (*project.ProjectAndMember, error) {
	var pam *project.ProjectAndMember
	session := p.conn.Session(ctx)
	sql := fmt.Sprintf("select a.*, b.project_code, b.member_code, b.join_time, b.is_owner, b.authorize from ms_project a, ms_project_member b where a.id=b.project_code and b.member_code= ? and b.project_code = ? limit 1")
	err := session.Raw(sql, id, code).Scan(&pam).Error
	return pam, err
}

func (p *ProjectDao) FindCollectByPidAndMemId(ctx context.Context, code int64, id int64) (bool, error) {
	var count int64
	session := p.conn.Session(ctx)
	sql := fmt.Sprintf("select count(*) from ms_project_collection where member_code= ? and project_code = ?")
	err := session.Raw(sql, id, code).Scan(&count).Error
	return count > 0, err
}

func (p *ProjectDao) UpdateDeletedProject(ctx context.Context, code int64, deleted bool) error {
	session := p.conn.Session(ctx)
	var err error
	if deleted {
		err = session.Model(&project.Project{}).Where("id=?", code).Update("deleted", 1).Error
	} else {
		err = session.Model(&project.Project{}).Where("id=?", code).Update("deleted", 0).Error
	}
	return err
}
func (p *ProjectDao) SaveProjectCollect(ctx context.Context, pc *project.ProjectCollection) error {
	return p.conn.Session(ctx).Save(&pc).Error
}

func (p *ProjectDao) DeleteProjectCollect(ctx context.Context, memId int64, projectCode int64) error {
	return p.conn.Session(ctx).Where("member_code=? and project_code=?", memId, projectCode).Delete(&project.ProjectCollection{}).Error
}
func (p *ProjectDao) UpdateProject(ctx context.Context, proj *project.Project) error {
	return p.conn.Session(ctx).Updates(&proj).Error
}
func (p *ProjectDao) FindMemberByProjectId(ctx context.Context, projectCode int64) (list []*project.ProjectMember, total int64, err error) {
	session := p.conn.Session(ctx)
	err = session.Model(&project.ProjectMember{}).Where("project_code=?", projectCode).
		Find(&list).Error
	if err != nil {
		return
	}
	err = session.Model(&project.ProjectMember{}).Where("project_code=?", projectCode).
		Count(&total).Error
	return
}
