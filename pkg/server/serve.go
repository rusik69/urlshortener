package server

import (
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/urlshortener/pkg/env"
)

// Serve registers starts the server.
func Serve() {
	rand.Seed(time.Now().UnixNano())
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")
	router.GET("/", rootHandler)
	router.GET("/shorten", shortenHandler)
	router.GET("/:shortURL", redirectHandler)
	router.GET("favicon.ico", func(c *gin.Context) {
		c.Status(404)
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	log.Fatal(router.Run(env.ConfigInstance.ListenHost + ":" + env.ConfigInstance.ListenPort))
}
