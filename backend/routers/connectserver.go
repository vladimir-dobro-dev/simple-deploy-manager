package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vladimir-dobro-dev/simple-deploy-manager/backend/common/sshclient"
)

func ConnectServer(c *gin.Context) {
	var connectData sshclient.ConnectData
	if c.Bind(&connectData) == nil {
		fmt.Print(connectData)
		status := "ok"
		_, err := sshclient.New(connectData)
		if err != nil {
			status = err.Error()
		}
		c.JSON(http.StatusOK, gin.H{"status": status})
		// c.JSON(http.StatusOK, gin.H{"status": "ok"})
		// const path = "/storage/emulated/0/Documents/test.txt"
		// f, _ := os.Create(path)
		// defer f.Close()
	}
	// c.String(http.StatusOK, "Connect...")
	// c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
