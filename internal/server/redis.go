package server

import (
	"context"
	"fmt"

	"github.com/North-al/gin-template/config"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func InitRedis() {
	dbConfig := config.GetConfig().Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", dbConfig.Host, dbConfig.Port),
		Password: dbConfig.Password,
		DB:       dbConfig.Database,
	})

	// 创建一个上下文对象
	ctx := context.Background()

	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(fmt.Errorf("redis ping error: %v", err))
	}

	RedisClient = rdb
}
