package server

import "github.com/gin-gonic/gin"

// rootHandler is the handler for the root path.
func rootHandler(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
