package dao

import (
	"hnz.com/ms_serve/ms-user/internal/database"
	"hnz.com/ms_serve/ms-user/internal/database/gorms"
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
		t.conn.Rollback()
		return err
	}
	t.conn.Commit()
	return nil
}
