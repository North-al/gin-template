package server

import (
	"fmt"

	"github.com/North-al/gin-template/config"
	"github.com/North-al/gin-template/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitHttp() {
	// init gin
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	r.Static("/uploads", "./uploads")

	publicV1 := r.Group("/api/v1")
	// 管理后台的base
	publicAdminV1 := publicV1.Group("/admin")

	privateV1 := r.Group("/api/v1")
	privateV1.Use(middleware.AuthMiddleware(RedisClient))
	privateAdminV1 := privateV1.Group("/admin")

	// if gin.Mode() == gin.DebugMode {
	InitDocs(r)
	// }
	//
	InitRouter(publicAdminV1, privateAdminV1)

	r.Run(fmt.Sprintf(":%d", config.GetConfig().Application.Port))
}
