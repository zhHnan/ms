package dao

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/account"
	"hnz.com/ms_serve/ms-project/internal/database/gorms"
)

type ProjectAuthDao struct {
	conn *gorms.GormConn
}

func NewProjectAuthDao() *ProjectAuthDao {
	return &ProjectAuthDao{
		conn: gorms.New(),
	}
}

func (p *ProjectAuthDao) FindAuthList(ctx context.Context, orgCode int64) (list []*account.ProjectAuth, err error) {
	session := p.conn.Session(ctx)
	err = session.Model(&account.ProjectAuth{}).Where("organization_code=? and status=1", orgCode).Find(&list).Error
	return
}
