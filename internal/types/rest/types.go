package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	Method  string          `json:"method" default:"GET"`
	Path    string          `json:"path" default:"/"`
	Handler gin.HandlerFunc `json:"handler"`
	IsAuth  bool            `json:"isAuth" default:"true"`
}

// NewRouteConfig 创建一个新的 RestRouteConfig。
//
// path: 路由的路径。
// handler: 路由的处理函数。
// options: 可选参数。如果传递一个 bool 值，则表示是否需要鉴权（true: 需要鉴权，false: 不需要鉴权）。
// 如果传递一个字符串，则表示请求方法(例如，GET、POST)。
//
// Example:
// NewRouteConfig("/auth/register", handler, true, "POST")
// NewRouteConfig("/auth/register", handler, http.MethodPost, true)
// NewRouteConfig("/auth/register", handler, true)
// NewRouteConfig("/auth/register", handler)
func NewRouteConfig(path string, handler gin.HandlerFunc, options ...any) *RouteConfig {
	method := http.MethodPost
	auth := true

	for _, option := range options {
		switch option := option.(type) {
		case string:
			if option != "GET" && option != "POST" && option != "PUT" && option != "DELETE" {
				panic("NewRestRouteConfig: (options) Invalid method, must be one of GET, POST, PUT, DELETE")
			}

			method = option
		case bool:
			auth = option
		}
	}

	return &RouteConfig{
		Method:  method,
		Path:    path,
		Handler: handler,
		IsAuth:  auth,
	}
}
