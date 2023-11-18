package routers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func InstallPath(c *gin.Context) {
	path, error := os.UserConfigDir()
	if error != nil {
		path = ""
	}
	c.String(http.StatusOK, path)
}
