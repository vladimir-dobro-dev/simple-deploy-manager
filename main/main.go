package main

import (
	"github.com/vladimir-dobro-dev/simple-deploy-manager/backend"
	"github.com/vladimir-dobro-dev/simple-deploy-manager/desktop"
)

func main() {
	backend.Run("0.0.0.0:2082")
	desktop.Run("http://127.0.0.1:2082/www")
}
