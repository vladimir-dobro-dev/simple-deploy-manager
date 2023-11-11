package main

import (
	backend "github.com/vladimir-dobro-dev/simple-deploy-manager"
	"github.com/vladimir-dobro-dev/simple-deploy-manager/desktop"
)

func main() {
	backend.Run()

	desktop.Run("http://127.0.0.1:2082/api/hello")
}
