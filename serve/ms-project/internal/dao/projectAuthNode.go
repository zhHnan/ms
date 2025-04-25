package dao

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/node"
	"hnz.com/ms_serve/ms-project/internal/database/gorms"
)

type ProjectAuthNodeDao struct {
	conn *gorms.GormConn
}

func NewProjectAuthNodeDao() *ProjectAuthNodeDao {
	return &ProjectAuthNodeDao{
		conn: gorms.New(),
	}
}

func (p *ProjectAuthNodeDao) FindNodeStringList(ctx context.Context, authId int64) (list []string, err error) {
	session := p.conn.Session(ctx)
	err = session.Model(&node.ProjectAuthNode{}).Where("auth=?", authId).Select("node").Find(&list).Error
	return
}
