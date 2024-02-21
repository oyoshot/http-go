package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world")
	})
	return r
}

func main() {
	r := setupRouter()
	if err := r.Run(":80"); err != nil {
		log.Fatal(err)
	}
}
