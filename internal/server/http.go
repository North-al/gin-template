package server

import (
	"github.com/North-al/go-gateway/internal/biz/service"
	"github.com/North-al/go-gateway/internal/data/repository"
	"github.com/North-al/go-gateway/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitHttp() {
	// init gin
	r := gin.Default()

	userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepoImpl(&gorm.DB{})))

	r.GET("/user/create", userHandler.Create)
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}
