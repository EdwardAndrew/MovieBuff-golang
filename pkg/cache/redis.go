package cache

import (
	"context"

	"github.com/go-redis/redis/v9"
)

type Cache interface {
}

var ctx = context.
	Background()

var rdb *redis.Client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func Increment(key string) error {
	return rdb.Incr(ctx, key).Err()
}
