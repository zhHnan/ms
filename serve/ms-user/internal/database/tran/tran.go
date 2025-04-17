package tran

import "hnz.com/ms_serve/ms-user/internal/database"

type Transaction interface {
	Action(func(conn database.DBConn) error) error
}
