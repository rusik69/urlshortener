package server

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rusik69/urlshortener/pkg/env"
)

var DB *sql.DB

// Serve registers starts the server.
func Serve(config env.Config, db *sql.DB) {
	DB = db
	rand.Seed(time.Now().UnixNano())
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")
	router.GET("/", rootHandler)
	router.POST("/shorten", shortenHandler)
	// router.GET("/:shortURL", redirectHandler)
	log.Fatal(router.Run(config.ListenHost + ":" + config.ListenPort))
}
