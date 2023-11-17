package model

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisClient interface {
	SMembers(key string) *redis.StringSliceCmd
	Get(key string) *redis.StringCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	SAdd(key string, members ...interface{}) *redis.IntCmd
	// Include other methods that you use from the Redis client
}
