package dao

import "hnz.com/ms_serve/ms-project/internal/database/gorms"

type AccountDao struct {
	conn *gorms.GormConn
}

func NewAccountDao() *AccountDao {
	return &AccountDao{
		conn: gorms.New(),
	}
}
