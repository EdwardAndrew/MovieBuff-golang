package cache

import (
	"context"
	"errors"
	"log"

	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

var rdb *redis.Client

func Connect(options *redis.Options) {
	rdb = redis.NewClient(options)
}

func Increment(key string) error {
	err := checkRedis()
	if err != nil {
		log.Print(err)
	}

	return rdb.Incr(ctx, key).Err()
}

func Get(key string) (string, error) {
	err := checkRedis()
	if err != nil {
		log.Print(err)
	}

	return rdb.Get(ctx, key).Result()
}

func checkRedis() error {
	var err error = nil
	if rdb == nil {
		err = errors.New("Redis is nil, have you called Connect?")
	}

	return err
}
