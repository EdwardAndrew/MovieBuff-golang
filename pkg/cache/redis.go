package cache

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

func Connect(o *redis.Options) {
	rdb = redis.NewClient(o)
}

func Close() {
	if rdb != nil {
		rdb.Close()
	}
}

func Increment(key string) error {
	if err := check(); err != nil {
		log.Print(err)
	}

	return rdb.Incr(ctx, key).Err()
}

// Takes a key and returns a bool indicating if any data was found,
//  cached data as a string. And also an error.
func Get(key string) (bool, string, error) {
	err := check()
	if err != nil {
		log.Print(err)
	}
	resp, err := rdb.Get(ctx, key).Result()

	if err == nil {
		return true, resp, nil
	} else if err.Error() == redis.Nil.Error() {
		return false, resp, nil
	}

	return false, resp, err
}

func Set(key, val string, expiration time.Duration) error {
	err := check()
	if err != nil {
		log.Print(err)
	}

	return rdb.Set(ctx, key, val, expiration).Err()
}

func SetMultiple(keys []string, values []string) error {
	var pairs []interface{}
	for i := range keys {
		pairs = append(pairs, keys[i], values[i])
	}

	return rdb.MSet(ctx, pairs...).Err()
}

func check() error {
	var err error = nil
	if rdb == nil {
		err = errors.New("Redis is nil, have you called Connect?")
	}

	return err
}
