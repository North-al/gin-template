package middleware

import (
	"net/http"
	"strconv"

	"github.com/North-al/gin-template/internal/pkg"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func AuthMiddleware(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			pkg.FailWithCode(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		_, claims, err := pkg.ParseToken(token)
		if err != nil {
			pkg.FailWithCode(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		userId := pkg.GetClaimsUserId(claims)
		username := pkg.GetClaimsUsername(claims)

		redisToken, err := redisClient.Get(c.Request.Context(), "token:"+strconv.FormatInt(userId, 10)).Result()

		if err == redis.Nil {
			pkg.FailWithCode(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		if err != nil || token != redisToken {
			pkg.FailWithCode(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Set("userId", userId)
		c.Set("username", username)

		c.Next()
	}
}
