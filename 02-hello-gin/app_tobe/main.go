package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	if err := newServer().Run(":8080"); err != nil {
		log.Fatalf("Could not run HTTP Server with (%v)", err)
		return
	}
}

func newServer() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("", helloHandler)
	r.GET("/:name", getOneSimpleHandler)
	r.POST("/", createSimpleHandler)

	if err := r.SetTrustedProxies([]string{"127.0.0.0/8"}); err != nil {
		return nil
	}

	return r
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}

func getOneSimpleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}

func createSimpleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}
