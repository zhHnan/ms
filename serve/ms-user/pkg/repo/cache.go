package repo

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Put(ctx context.Context, key, value string, expire time.Duration) error
	//Del(key string) error
}
