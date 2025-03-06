package pkg

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

// Response 统一响应结构体
type Response struct {
	Code    int    `json:"code" oneof:"200,400,401,404,500" example:"200"` // 业务状态码 200成功 | 500失败
	Message string `json:"message" example:"success"`                      // 消息
	Data    any    `json:"data"`                                           // 数据
}

func Success(c *gin.Context, data any, message ...string) {
	var resp = new(Response)
	resp.Code = http.StatusOK
	resp.Data = data
	if len(message) > 0 {
		resp.Message = message[0]
	} else {
		resp.Message = "success"
	}
	
	c.JSON(http.StatusOK, resp)
}

func Fail(c *gin.Context, message string) {
	var resp = new(Response)
	resp.Code = http.StatusBadRequest
	resp.Data = nil
	resp.Message = message
	c.JSON(http.StatusOK, resp)
}

// FailWithCode 失败返回(自定义业务状态码)
// code 业务状态码
// data 数据
// message 消息
func FailWithCode(c *gin.Context, code int, message string, data ...any) {
	var resp = new(Response)
	resp.Code = code
	resp.Message = message
	resp.Data = data
	
	c.JSON(http.StatusOK, resp)
}
