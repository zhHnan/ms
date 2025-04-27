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
//	Rc = &RedisCache{Rdb: rdb}
//}

func (rc *RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) error {
	err := rc.Rdb.Set(key, value, expire).Err()
	return err
}
func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := rc.Rdb.Get(key).Result()
	return result, err
}
func (rc *RedisCache) HSet(ctx context.Context, key string, field string, value string) {
	rc.Rdb.HSet(key, field, value)
}
func (rc *RedisCache) HKeys(ctx context.Context, key string) ([]string, error) {
	return rc.Rdb.HKeys(key).Result()
}
func (rc *RedisCache) Delete(ctx context.Context, keys []string) {
	rc.Rdb.Del(keys...)
}
