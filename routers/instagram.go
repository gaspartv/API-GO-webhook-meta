package routers

import (
	docs "github.com/gaspartv/API-GO-opportunities/docs"
	"github.com/gaspartv/API-GO-webhook-meta/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutesInstagram(router *gin.Engine) {
	// Initialize Handler
	handlers.InitializeHandler()
	basePath := "/api/v1/messenger/instagram/webhook"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/")
		v1.POST("/")

	}

}
