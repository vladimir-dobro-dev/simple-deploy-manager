package backend

import (
	"github.com/vladimir-dobro-dev/simple-deploy-manager/webapi"
)

func Run() {
	router := webapi.SetRoutes()
	go router.Run("127.0.0.1:2082")
}
