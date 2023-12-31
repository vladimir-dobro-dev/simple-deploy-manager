package routers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	//	"github.com/vladimir-dobro-dev/simple-deploy-manager/backend/common/sshclient"
)

type ConnectData struct {
	Address  string
	UserName string
	Password string
	Port     string
}

func ConnectServer(c *gin.Context) {
	var connectData ConnectData
	if c.Bind(&connectData) == nil {
		fmt.Print(connectData)
		//client := sshclient.Connect()
		//client.Close()
		const path = "/storage/emulated/0/Documents/test.txt"
		f, _ := os.Create(path)
		defer f.Close()
	}
	c.String(http.StatusOK, "Connect...")
	// c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
