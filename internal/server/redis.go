package server

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6399",
		Password: "",
		DB:       0,
	})

	// 创建一个上下文对象
	ctx := context.Background()

	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(fmt.Errorf("redis ping error: %v", err))
	}

	RedisClient = rdb
}
