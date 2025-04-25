package dao

import (
	"errors"

	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-project/internal/database"
	"hnz.com/ms_serve/ms-project/internal/database/gorms"
)

type Trans struct {
	conn database.DBConn
}

func NewTrans() *Trans {
	return &Trans{
		conn: gorms.NewTrans(),
	}
}

func (t Trans) Action(f func(conn database.DBConn) error) error {
	t.conn.Begin()
	err := f(t.conn)
	if err != nil {
		// 检查是否为BError类型
		var bErr *errs.BError
		if errors.As(err, &bErr) {
			t.conn.Rollback()
			return bErr
		}
		// 不是BError类型，直接返回一般错误
		t.conn.Rollback()
		return err
	}
	t.conn.Commit()
	return nil
}
