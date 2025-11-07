package cache

import (
	"context"
	"go-api/config"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisCache() *RedisCache {
	return &RedisCache{
		client: config.RedisClient,
		ctx:    context.Background(),
	}
}

func (r *RedisCache) SetOne(key string, value interface{}) error {
	return r.client.Set(r.ctx, key, value, 0).Err()
}

func (r *RedisCache) SetOneWithExpire(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(r.ctx, key, value, expiration).Err()
}

func (r *RedisCache) GetOne(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

func (r *RedisCache) DeleteOne(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

func (r *RedisCache) Incr(key string) (int64, error) {
	return r.client.Incr(r.ctx, key).Result()
}

func (r *RedisCache) Expire(key string, expiration time.Duration) error {
	return r.client.Expire(r.ctx, key, expiration).Err()
}
