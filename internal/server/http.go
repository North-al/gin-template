package server

import (
	"github.com/gin-gonic/gin"
)

func InitHttp() {
	// init gin
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}
