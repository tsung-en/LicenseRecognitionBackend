package controller

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	c.String(200, "Hello World")
}
