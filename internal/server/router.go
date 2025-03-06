package server

import (
	"fmt"
	"reflect"

	"github.com/North-al/gin-template/internal/biz/service"
	"github.com/North-al/gin-template/internal/data/repository"
	"github.com/North-al/gin-template/internal/handler"
	"github.com/North-al/gin-template/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
  此文件用于自动注册路由
*/

func InitRouter(router *gin.Engine) {
	userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepoImpl(&gorm.DB{})))

	// r.GET("/user/create", userHandler.Create)
	allHandler := []interface{}{
		userHandler,
	}

	// 遍历 allHandler，自动注册路由
	for _, h := range allHandler {
		AutoRegisterRoutes(router, h)
	}
}

func AutoRegisterRoutes(router *gin.Engine, handler interface{}) {
	handlerTypeof := reflect.TypeOf(handler)
	handlerValue := reflect.ValueOf(handler)
	fmt.Println("run:", handlerValue)

	for i := 0; i < handlerTypeof.NumMethod(); i++ {
		// 获取方法
		method := handlerTypeof.Method(i)

		// 获取包的类型名称
		prefix := handlerTypeof.Elem().Name()
		// 获取方法名称
		methodName := handlerTypeof.Method(i).Name
		prefix = utils.CamelToSnake(prefix)
		methodName = utils.CamelToSnake(methodName)

		routePath := "/" + prefix + "/" + methodName
		fmt.Println("routePath:", routePath)

		handlerFunc := func(c *gin.Context) {
			method.Func.Call([]reflect.Value{handlerValue, reflect.ValueOf(c)})
		}

		router.Handle("POST", routePath, handlerFunc)
	}
}
