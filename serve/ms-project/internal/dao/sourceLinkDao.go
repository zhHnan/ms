package dao

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/files"
	"hnz.com/ms_serve/ms-project/internal/database/gorms"
)

type SourceLinkDao struct {
	conn *gorms.GormConn
}

func NewSourceLinkDao() *SourceLinkDao {
	return &SourceLinkDao{
		conn: gorms.New(),
	}
}

func (s *SourceLinkDao) Save(ctx context.Context, link *files.SourceLink) error {
	return s.conn.Session(ctx).Save(&link).Error
}

func (s *SourceLinkDao) FindByTaskCode(ctx context.Context, taskCode int64) (list []*files.SourceLink, err error) {
	session := s.conn.Session(ctx)
	err = session.Model(&files.SourceLink{}).Where("link_code=?", taskCode).Find(&list).Error
	return
}
