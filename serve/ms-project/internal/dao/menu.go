package dao

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/menu"
	"hnz.com/ms_serve/ms-project/internal/database/gorms"
)

type MenuDao struct {
	conn *gorms.GormConn
}

func NewMenuDao() *MenuDao {
	return &MenuDao{
		conn: gorms.New(),
	}
}

func (m *MenuDao) FindAll(ctx context.Context) ([]*menu.ProjectMenu, error) {
	var pms []*menu.ProjectMenu
	err := m.conn.Session(ctx).Find(&pms).Error
	return pms, err
}
