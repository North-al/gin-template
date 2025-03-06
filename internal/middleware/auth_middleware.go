package middleware

import (
	"net/http"

	"github.com/North-al/restaurant/internal/pkg"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
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

		// TODO: 缺少redis比对

		userId := pkg.GetClaimsUserId(claims)
		username := pkg.GetClaimsUsername(claims)

		c.Set("userId", userId)
		c.Set("username", username)

		c.Next()
	}
}
