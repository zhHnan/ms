package repo

import "context"

type ProjectAuthNodeRepo interface {
	FindNodeStringList(ctx context.Context, authId int64) (list []string, err error)
}
