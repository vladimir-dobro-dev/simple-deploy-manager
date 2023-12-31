package staticfiles

import (
	"embed"
	"io/fs"
)

//go:embed frontend/dist
var www embed.FS

func GetFS() fs.FS {
	staticFiles, _ := fs.Sub(www, "frontend/dist")
	return staticFiles
}
