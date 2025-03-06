package pkg

import (
	"time"
	
	"github.com/North-al/gin-template/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string, userId int64) (string, error) {
	claims := &jwt.MapClaims{
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Duration(config.GetConfig().JWT.Expire) * time.Hour).Unix(), // 过期时间24小时
		"iss":      "North-restaurant",
		"sub":      userId,
		"username": username,
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString([]byte(config.GetConfig().JWT.Secret))
}

func ParseToken(tokenString string) (*jwt.Token, *jwt.MapClaims, error) {
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JWT.Secret), nil
	})
	
	if err != nil {
		return nil, nil, err
	}
	
	return token, claims, nil
}

func GetClaimsUsername(claims *jwt.MapClaims) string {
	return (*claims)["username"].(string)
}

func GetClaimsUserId(claims *jwt.MapClaims) int64 {
	return int64((*claims)["sub"].(float64))
}
