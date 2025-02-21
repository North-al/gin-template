package main

import "github.com/North-al/go-gateway/internal/server"

func main() {
	server.InitDB()
	server.InitRedis()
	server.InitHttp()
}
