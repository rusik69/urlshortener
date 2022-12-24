package server

import "github.com/gin-gonic/gin"

// shortenHandler is the handler for the shorten path.
func shortenHandler(c *gin.Context) {
	url := c.PostForm("url")
	key := generateKey(url)
	err := saveURL(url, key)
	if err != nil {
		c.HTML(500, "error.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(200, "shorten.html", gin.H{"shortURL": key})
}
