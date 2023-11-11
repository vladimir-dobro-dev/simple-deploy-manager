package webapi

import (
	"github.com/gin-gonic/gin"

	"github.com/vladimir-dobro-dev/simple-deploy-manager/webapi/controllers"
)

func SetRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/api/hello", controllers.Hello)

	return router
}
