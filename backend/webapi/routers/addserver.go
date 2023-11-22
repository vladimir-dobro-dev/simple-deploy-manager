package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerConnect struct {
	Address  string
	UserName string
	Password string
	Port     string
}

func AddServer(c *gin.Context) {
	var connectData ServerConnect
	if c.Bind(&connectData) == nil {
		fmt.Print(connectData)
	}
	c.String(http.StatusOK, "Connect...")
	// c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
