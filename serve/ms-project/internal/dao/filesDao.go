package dao

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/files"
	"hnz.com/ms_serve/ms-project/internal/database/gorms"
)

type FileDao struct {
	conn *gorms.GormConn
}

func NewFileDao() *FileDao {
	return &FileDao{
		conn: gorms.New(),
	}
}

func (f *FileDao) FindByIds(ctx context.Context, ids []int64) (list []*files.File, err error) {
	session := f.conn.Session(ctx)
	err = session.Model(&files.File{}).Where("id in (?)", ids).Find(&list).Error
	return
}

func (f *FileDao) Save(ctx context.Context, file *files.File) error {
	err := f.conn.Session(ctx).Save(&file).Error
	return err
}
