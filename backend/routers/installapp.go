package routers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/vladimir-dobro-dev/simple-deploy-manager/backend/common/sshclient"
)

func InstallApp(c *gin.Context) {
	fileHeader, err := c.FormFile("appFile")

	if err != nil {
		return
	}

	src, err := fileHeader.Open()
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.CreateTemp("", "app")
	if err != nil {
		return
	}
	defer os.Remove(dst.Name())
	defer dst.Close()

	data := make([]byte, 1024*1024)
	for {
		n, err := src.Read(data)
		if err == io.EOF {
			break
		}
		dst.Write(data[:n])
	}

	client := sshclient.New()
	err = client.Upload(dst.Name(), "app.zip")
	if err != nil {
		return
	}
	c.String(http.StatusOK, fmt.Sprint(fileHeader.Size))
}
