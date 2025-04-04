package main

import (
	"github.com/North-al/gin-template/internal/pkg/logger"
	"github.com/North-al/gin-template/internal/server"
)

//	@title			gin-template
//	@version		1.3.0
//	@description	这是一个GO Gin template系统的API文档。
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/api/v1/admin

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	logger.InitLogger()
	server.InitDB()
	server.InitRedis()
	server.InitHttp()
}
