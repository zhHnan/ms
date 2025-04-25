package dao

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/node"
	"hnz.com/ms_serve/ms-project/internal/database/gorms"
)

type ProjectNodeDao struct {
	conn *gorms.GormConn
}

func (m *ProjectNodeDao) FindAll(ctx context.Context) (pms []*node.ProjectNode, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&node.ProjectNode{}).Find(&pms).Error
	return
}

func NewProjectNodeDao() *ProjectNodeDao {
	return &ProjectNodeDao{
		conn: gorms.New(),
	}
}
