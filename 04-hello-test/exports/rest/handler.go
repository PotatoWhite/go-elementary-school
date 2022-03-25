package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"potato/simple-rest/entities/dto"
	"strconv"
)

func NewServer() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/", simpleHandler)
	r.GET("/:name", pathParamHandler)
	r.GET("/:name/:quantity", pathParamsHandler)
	r.POST("/", requestBodyHandler)

	if err := r.SetTrustedProxies([]string{"127.0.0.0/8"}); err != nil {
		return nil
	}

	return r
}

func pathParamsHandler(c *gin.Context) {
	stringValue := c.Param("name")
	numericString := c.Param("quantity")
	if quantity, err := strconv.Atoi(numericString); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, &dto.BasicResponse{
			Code:    0,
			Message: fmt.Sprintf("%v, %v개 주세요.", stringValue, quantity),
		})
	}
}

func simpleHandler(c *gin.Context) {
	name := c.Query("name")
	c.JSON(http.StatusOK, &dto.BasicResponse{
		Code:    0,
		Message: fmt.Sprintf("%v 입니다.", name),
	})
}

func pathParamHandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, &dto.BasicResponse{
		Code:    0,
		Message: fmt.Sprintf("%v 좋아요.", name),
	})
}

func requestBodyHandler(c *gin.Context) {
	var requestBody dto.Simple
	if err := c.ShouldBind(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, &dto.BasicResponse{
			Code:    0,
			Message: fmt.Sprintf("%v", err.Error()),
		})
		return
	} else {
		c.JSON(http.StatusOK, requestBody)
	}

	return
}
