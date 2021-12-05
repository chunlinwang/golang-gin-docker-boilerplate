package controllers

import (
	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (this *TestController) JsonApi(c *gin.Context) {
	c.JSON(200, gin.H{
		"html": "<b>Hello, world!</b>",
	})
}

func (this *TestController) JsonPureApi(c *gin.Context) {
	c.PureJSON(200, gin.H{
		"html": "<b>Hello, world!</b>",
	})
}
