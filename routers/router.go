package routers

import (
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Initialize() {
	router := gin.Default()
	
	initializeRoutesFacebook(router)
	initializeRoutesInstagram(router)
	initializeRoutesWhatsapp(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")

	router.Run("0.0.0.0:" + port)
}
