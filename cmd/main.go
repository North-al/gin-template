package main

import (
	"github.com/North-al/go-gateway/internal/pkg/logger"
	"github.com/North-al/go-gateway/internal/server"
)

func main() {
	logger.InitLogger()
	defer logger.Sync()
	server.InitDB()
	server.InitRedis()
	server.InitHttp()
}
