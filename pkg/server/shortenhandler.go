package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rusik69/urlshortener/pkg/env"
	"github.com/sirupsen/logrus"
)

// shortenHandler is the handler for the shorten path.
func shortenHandler(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.Writer.WriteHeader(400)
		c.Writer.Write([]byte("no url provided"))
		return
	}
	key := generateKey()
	logrus.Println("add", key, url)
	err := saveURL(url, key)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
		return
	}
	c.Writer.WriteHeader(200)
	c.Writer.Write([]byte(env.ConfigInstance.Host + key))
}
