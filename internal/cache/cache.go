package cache

import "time"

type CacheRepository interface {
	SetOne(key string, value interface{}) error
	SetOneWithExpire(key string, value interface{}, expiration time.Duration) error
	GetOne(key string) (string, error)
	DeleteOne(key string) error
	Incr(key string) (int64, error)
	Expire(key string, expiration time.Duration) error
}
