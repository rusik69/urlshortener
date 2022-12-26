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
	url, err := database.GetKey(key, DB)
	if err != nil {
		c.HTML(500, "error.html", gin.H{"error": err.Error()})
		return
	}
	http.RedirectHandler(url, http.StatusSeeOther)
}
