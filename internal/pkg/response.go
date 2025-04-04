package pkg

import (
	"net/http"

	"github.com/North-al/gin-template/internal/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构体
type Response struct {
	Code    int    `json:"code" example:"200"`   // HTTP状态码
	Message string `json:"message" example:"成功"` // 消息
	Data    any    `json:"data"`                 // 数据
}

// successResponse 封装成功响应
func successResponse(c *gin.Context, code int, data any, message string) *gin.Context {
	resp := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	// 记录成功日志
	logger.WithContext(c.Request.Context()).Info("request success",
		"path", c.FullPath(),
		"status", code,
	)

	c.JSON(code, resp)
	return c
}

// errorResponse 封装错误响应
func errorResponse(c *gin.Context, code int, message string, data any) *gin.Context {
	resp := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	// 记录错误日志
	logger.WithContext(c.Request.Context()).Error("request failed",
		"path", c.FullPath(),
		"status", code,
		"message", message,
	)

	c.JSON(code, resp)
	return c
}

// Success 成功响应
func Success(c *gin.Context, data any, message ...string) {
	msg := "成功"
	if len(message) > 0 {
		msg = message[0]
	}
	successResponse(c, http.StatusOK, data, msg)
}

// BadRequest 400错误
func BadRequest(c *gin.Context, message string) {
	errorResponse(c, http.StatusBadRequest, message, nil)
}

// NotFound 404错误
func NotFound(c *gin.Context, message string) {
	errorResponse(c, http.StatusNotFound, message, nil)
}

// InternalError 500错误
func InternalError(c *gin.Context, message string) {
	errorResponse(c, http.StatusInternalServerError, message, nil)
}

// FailWithCode 自定义状态码的失败响应
func FailWithCode(c *gin.Context, code int, message string, data ...any) {
	var respData any
	if len(data) > 0 {
		respData = data[0]
	}
	errorResponse(c, code, message, respData)
}
