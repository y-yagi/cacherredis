package cacherredis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/y-yagi/cacher"
)

// WithRedisStore create a new Cacher with a Reids store.
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

// Read reads cache from a Redis store.
func (rs *RedisStore) Read(key string) ([]byte, error) {
	value, err := rs.client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	return []byte(value), nil
}

// Write stores data to a Redis store.
func (rs *RedisStore) Write(key string, value []byte, d time.Duration) error {
	err := rs.client.Set(key, string(value), d).Err()
	return err
}

// Delete deletes data from a Redis store.
func (rs *RedisStore) Delete(key string) error {
	_, err := rs.client.Del(key).Result()
	return err
}

// Cleanup deletes the expired cache.
func (rs *RedisStore) Cleanup() error {
	// Redis clear expired cache.  So do not need to this in `RedisStore`.
	return nil
}

// Exist check the cache exists or not.
func (rs *RedisStore) Exist(key string) bool {
	result, _ := rs.client.Exists(key).Result()
	return result != 0
}
