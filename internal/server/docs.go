package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	gs "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"

	"github.com/North-al/gin-template/config"
	_ "github.com/North-al/gin-template/docs"
)

func InitDocs(r *gin.Engine) {
	url := gs.URL(fmt.Sprintf("http://%s:%d/swagger/doc.json", config.GetConfig().Docs.Host, config.GetConfig().Application.Port))
	if gin.Mode() == gin.ReleaseMode {
		url = gs.URL(fmt.Sprintf("https://%s/swagger/doc.json", config.GetConfig().Docs.Host))
	}

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler,
		url, // Adjust URL as needed
		gs.DefaultModelsExpandDepth(-1),
		gs.DocExpansion("none")))
}
