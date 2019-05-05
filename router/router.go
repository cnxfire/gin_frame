package router

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200,gin.H{"message":"pong"})
	})
}
