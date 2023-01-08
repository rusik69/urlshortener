package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/urlshortener/pkg/database"
)

// redirectHandler is the handler for the redirect path.
func redirectHandler(c *gin.Context) {
	key := strings.Split(c.Request.URL.Path, "/")[1]
	url, err := database.GetKey(key)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		return
	}
	c.Redirect(http.StatusMovedPermanently, url)
}
