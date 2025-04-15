package dao

import (
	"context"
	"github.com/go-redis/redis"
	"time"
)

var Rc *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	Rc = &RedisCache{rdb: rdb}
}

func (rc *RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) error {

	err := rc.rdb.Set(key, value, expire).Err()
	return err
}
func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := rc.rdb.Get(key).Result()
	return result, err

}
