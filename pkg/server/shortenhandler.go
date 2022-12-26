package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// shortenHandler is the handler for the shorten path.
func shortenHandler(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.HTML(400, "error.html", gin.H{"error": "no url provided"})
		return
	}
	key := generateKey()
	logrus.Println("add", key, url)
	err := saveURL(url, key)
	if err != nil {
		c.HTML(500, "error.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(200, "shorten.html", gin.H{"shortURL": key})
}
