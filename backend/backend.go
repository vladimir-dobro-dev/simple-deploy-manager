package backend

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/vladimir-dobro-dev/simple-deploy-manager/backend/webapi/routers"
)

func Run() {
	router := SetRoutes()
	router.Run("127.0.0.1:2082")
}

func SetRoutes() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*"}
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	router.Use(cors.New(config))

	router.GET("/api/hello", routers.Hello)
	router.GET("/api/installpath", routers.InstallPath)

	return router
}
