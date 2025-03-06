package main

import (
	"github.com/North-al/gin-template/internal/pkg/logger"
	"github.com/North-al/gin-template/internal/server"
)

func main() {
	logger.InitLogger()
	defer logger.Sync()
	server.InitDB()
	server.InitRedis()
	server.InitHttp()
}
