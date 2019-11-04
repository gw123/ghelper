package cache

import (
	"github.com/go-redis/redis"
	"time"
)

//适配器模式 ,将redisclient ,包装成 cacheClient
type RedisCache struct {
	client *redis.Client
	prekey string
}

func NewRedisCache(client *redis.Client, prekey string) *RedisCache {
	return &RedisCache{client: client, prekey: prekey}
}

func (r RedisCache) Get(key string) (interface{}, error) {
	cmd := r.client.Get(key)
	return cmd.Val(), cmd.Err()
}

func (r RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(key, value, expiration).Err()
}

func (r RedisCache) Del(key string) error {
	return r.client.Del(key).Err()
}

func (r RedisCache) Incr(key string) error {
	return r.client.Incr(key).Err()
}
