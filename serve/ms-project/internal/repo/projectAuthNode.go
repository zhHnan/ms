package repo

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/database"
)

type ProjectAuthNodeRepo interface {
	FindNodeStringList(ctx context.Context, authId int64) (list []string, err error)
	DeleteByAuthId(background context.Context, conn database.DBConn, authId int64) error
	Save(background context.Context, conn database.DBConn, authId int64, nodes []string) error
}
