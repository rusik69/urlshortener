package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/urlshortener/pkg/env"
)

// rootHandler is the handler for the root path.
func rootHandler(c *gin.Context) {
	data := struct {
		MaxURLLength string
	}{
		MaxURLLength: fmt.Sprintf("%d", env.ConfigInstance.MaxURLLength),
	}
	c.HTML(200, "index.html", data)
}
