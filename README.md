# CacherRedis

[![GoDoc](https://godoc.org/github.com/y-yagi/cacherredis?status.svg)](https://godoc.org/github.com/y-yagi/cacherredis)
[![Go Report Card](https://goreportcard.com/badge/github.com/y-yagi/cacherredis)](https://goreportcard.com/report/github.com/y-yagi/cacherredis)
[![Build Status](https://circleci.com/gh/y-yagi/cacherredis.svg?style=svg)](https://circleci.com/gh/y-yagi/cacherredis)

CacherRedis is the extension package for [cacher](https://github.com/y-yagi/cacher). This provides Redis store support.

Example:

```go
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
```
