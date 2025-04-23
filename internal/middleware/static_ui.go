package middleware

import (
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type localFileSystem struct {
	http.FileSystem
	root      string
	urlPrefix string
}

func StaticUi(urlPrefix, spaDirectory string) gin.HandlerFunc {
	directory := &localFileSystem{
		FileSystem: gin.Dir(spaDirectory, true),
		root:       spaDirectory,
		urlPrefix:  urlPrefix,
	}
	fileServer := http.FileServer(directory)

	if urlPrefix != "" {
		fileServer = http.StripPrefix(urlPrefix, fileServer)
	}

	return func(c *gin.Context) {
		log.Info().Msgf("redirecting to UI, got path: %s", c.Request.URL.Path)
		if !directory.Exists(c.Request.URL.Path) || c.Request.URL.Path == urlPrefix {
			log.Info().Msgf("missing directory, serving index.html")
			c.Request.URL.Path = urlPrefix + "/index.html"
		}

		c.Request.URL.Path = strings.TrimSuffix(c.Request.URL.Path, "index.html")
		log.Info().Msgf("serving route: %s", c.Request.URL.Path)
		fileServer.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}

func (l *localFileSystem) Exists(filepath string) bool {
	if p := strings.TrimPrefix(filepath, l.urlPrefix); len(p) < len(filepath) {
		name := path.Join(l.root, p)
		_, err := os.Stat(name)
		return err == nil
	}
	return false
}
