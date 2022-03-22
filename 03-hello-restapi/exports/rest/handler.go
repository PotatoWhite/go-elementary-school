package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"potato/simple-rest/entities/dto"
)

func NewServer() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/", helloHandler)
	r.GET("/:name", getOneSimpleHandler)
	r.POST("/", createSimpleHandler)

	if err := r.SetTrustedProxies([]string{"127.0.0.0/8"}); err != nil {
		return nil
	}

	return r
}

func helloHandler(c *gin.Context) {
	message := c.Query("message")
	c.JSON(http.StatusOK, message)
}

func getOneSimpleHandler(c *gin.Context) {
	message := c.Param("name")
	c.JSON(http.StatusOK, message)
}

func createSimpleHandler(c *gin.Context) {
	var create dto.Simple
	if err := c.ShouldBind(&create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("err: %v", err),
		})
		return
	} else {
		c.JSON(http.StatusOK, create)
	}

	return
}
