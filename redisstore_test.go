package cacherredis

import (
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/y-yagi/cacher"
)

func TestRedisStore(t *testing.T) {
	value := []byte("dummy")
	c := WithRedisStore(redisOpt())

	got, _ := c.Read("cacher-test")
	if got != nil {
		t.Fatalf("want nil, got %q", got)
	}

	c.Write("cacher-test", value, cacher.Forever)
	got, _ = c.Read("cacher-test")
	if string(got) != string(value) {
		t.Fatalf("want %q, got %q", value, got)
	}

	c.Delete("cacher-test")
	got, _ = c.Read("cacher-test")
	if got != nil {
		t.Fatalf("want nil, got %q", got)
	}
}

func TestReidsStoreWithExpired(t *testing.T) {
	value := []byte("dummy")
	c := WithRedisStore(redisOpt())

	c.Write("cacher-test", value, 1*time.Second)

	time.Sleep(2 * time.Second)

	got, _ := c.Read("cacher-test")
	if got != nil {
		t.Fatalf("want nil, got %q", got)
	}
}

func TestFileStore_Exist(t *testing.T) {
	c := WithRedisStore(redisOpt())

	if c.Exist("not-exist") {
		t.Fatalf("want false, got true")
	}

	c.Write("exist", []byte("foo"), cacher.Forever)
	if !c.Exist("exist") {
		t.Fatalf("want true, got false")
	}
}

func redisOpt() *redis.Options {
	return &redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}
}
