package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 Not Found")
	})

	r.Run(":8080") // Default listens and serves on 0.0.0.0:8080
}
