package dao

import (
	"context"
	"github.com/go-redis/redis"
	"time"
)

var Rc *RedisCache

type RedisCache struct {
	Rdb *redis.Client
}

//func init() {
//	rdb := redis.NewClient(config.Cfg.ReadRedisConfig())
//	Rc = &RedisCache{rdb: rdb}
//}

func (rc *RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) error {
	err := rc.Rdb.Set(key, value, expire).Err()
	return err
}
func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := rc.Rdb.Get(key).Result()
	return result, err
}
