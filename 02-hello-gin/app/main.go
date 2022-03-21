package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	if err := newServer().Run(); err != nil {
		log.Fatalf("Could not run HTTP Server with (%v)", err)
		return
	}
}

func newServer() *gin.Engine {
	r := gin.Default()
	r.GET("", helloHandler)
	r.GET("/:name", getOneSimpleHandler)
	r.POST("/", createSimpleHandler)

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
