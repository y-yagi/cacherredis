package cacherredis_test

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/y-yagi/cacher"
	cacherredis "github.com/y-yagi/cacher-redis"
)

func ExampleRedisStore_Read() {
	opt := redis.Options{Addr: "localhost:6379"}
	c := cacherredis.WithRedisStore(&opt)
	c.Write("cache-key", []byte("value"), cacher.Forever)

	value, _ := c.Read("cache-key")
	fmt.Print(string(value))
	// Output: value
}

func ExampleRedisStore_Delete() {
	opt := redis.Options{Addr: "localhost:6379"}
	c := cacherredis.WithRedisStore(&opt)
	c.Write("cache-key", []byte("value"), cacher.Forever)

	c.Delete("cache-key")
	value, _ := c.Read("cache-key")
	fmt.Print(string(value))
	// Output:
}
