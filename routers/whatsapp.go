package routers

import (
	docs "github.com/gaspartv/API-GO-opportunities/docs"
	"github.com/gaspartv/API-GO-webhook-meta/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutesWhatsapp(router *gin.Engine) {
	handlers.InitializeHandler()

	basePath := "/api/v1/whatsapp/webhook"

	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/", handlers.WhatsappValidateHandler)
		v1.POST("/", handlers.WhatsappReceiveHandler)
	}
}
