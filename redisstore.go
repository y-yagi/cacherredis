package cacherredis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/y-yagi/cacher"
)

// WithRedisStore create a new Cache with Redis store.
func WithRedisStore(opt *redis.Options) *cacher.Cacher {
	cache := &cacher.Cacher{}
	client := redis.NewClient(opt)
	cache.Store = &RedisStore{client: client}
	return cache
}

// RedisStore is a type for RedisStore.
type RedisStore struct {
	client *redis.Client
}

// Read cache.
func (rs *RedisStore) Read(key string) ([]byte, error) {
	value, err := rs.client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	return []byte(value), nil
}

// Write create a new cache.
func (rs *RedisStore) Write(key string, value []byte, d time.Duration) error {
	err := rs.client.Set(key, string(value), d).Err()
	return err
}

// Delete delete cache.
func (rs *RedisStore) Delete(key string) error {
	_, err := rs.client.Del(key).Result()
	return err
}
