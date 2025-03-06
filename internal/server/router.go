package server

import (
	"log"
	"reflect"
	
	"github.com/North-al/gin-template/internal/biz/service"
	"github.com/North-al/gin-template/internal/data/repository"
	"github.com/North-al/gin-template/internal/handler"
	"github.com/North-al/gin-template/internal/types/rest"
	"github.com/gin-gonic/gin"
)

/*
  此文件用于自动注册路由
*/

func InitRouter(publicRouter *gin.RouterGroup, privateRouter *gin.RouterGroup) {
	authHandler := handler.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(DB)))
	
	allHandler := []interface{}{
		authHandler,
	}
	
	// 遍历 allHandler，自动注册路由
	for _, h := range allHandler {
		AutoRegisterRoutes(publicRouter, privateRouter, h)
	}
}

func AutoRegisterRoutes(publicRouter *gin.RouterGroup, privateRouter *gin.RouterGroup, handler interface{}) {
	v := reflect.ValueOf(handler)
	// 查找 `RegisterRoutes` 方法
	method := v.MethodByName("RegisterRoutes")
	if !method.IsValid() {
		log.Printf("handler %T 没有 RegisterRoutes 方法，跳过\n", handler)
		return
	}
	
	// 检查方法返回值
	if method.Type().NumOut() != 1 || method.Type().Out(0) != reflect.TypeOf([]*rest.RouteConfig{}) {
		log.Printf("handler %T 的 RegisterRoutes 方法返回值不符合要求，跳过\n", handler)
	}
	
	// 调用方法获取路由配置
	results := method.Call(nil)
	routeConfigs, ok := results[0].Interface().([]*rest.RouteConfig)
	if !ok {
		log.Printf("handler %T 的 RegisterRoutes 方法返回值类型转换失败，跳过\n", handler)
	}
	
	// 遍历 routes 并注册
	for _, route := range routeConfigs {
		if route.IsAuth {
			privateRouter.Handle(route.Method, route.Path, route.Handler)
		} else {
			publicRouter.Handle(route.Method, route.Path, route.Handler)
		}
	}
}
