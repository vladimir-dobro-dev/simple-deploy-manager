package backend

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	staticfiles "github.com/vladimir-dobro-dev/simple-deploy-manager"
	"github.com/vladimir-dobro-dev/simple-deploy-manager/backend/routers"
)

func Run(url string) {
	router := SetRouter()
	router.StaticFS("www", http.FS(staticfiles.GetFS()))
	go router.Run(url)
}

func SetRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*"}
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	router.Use(cors.New(config))

	router.GET("/api/hello", routers.Hello)
	router.POST("/api/addserver", routers.ConnectServer)
	router.POST("/api/installapp", routers.InstallApp)

	return router
}
